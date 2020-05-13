package service

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"git.chinaopen.ai/yottacloud/go-libs/redis"
	libredis "github.com/go-redis/redis/v7"
	"github.com/go-redis/redis_rate/v8"
	"go.uber.org/zap"

	"starter/config"
	"starter/gen/log"
	"starter/util"
	"starter/util/qcloudsms"
)

const (
	ErrSaveSmsCodeFailed = "保存短信验证码失败"
	ErrSendSmsCodeFailed = "发送短信验证码失败"
)

func NewSmsService(logger *log.Logger) *SmsService {
	opts, _ := libredis.ParseURL(config.C.Redis.URI)
	rdb := libredis.NewClient(opts)

	limiter := redis_rate.NewLimiter(rdb)

	return &SmsService{
		logger:  logger,
		limiter: limiter,
	}
}

type SmsService struct {
	logger  *log.Logger
	limiter *redis_rate.Limiter
}

func PerDay(rate int) *redis_rate.Limit {
	return &redis_rate.Limit{
		Rate:   rate,
		Period: time.Hour,
		Burst:  rate,
	}
}

// 生成验证码字符串
func (s SmsService) SendSmsCode(userID, mobile string) error {
	// 当前用户每分钟限制
	if res, _ := s.limiter.Allow(userID, redis_rate.PerMinute(config.C.VerifyCode.PerMinute)); !res.Allowed {
		return errors.New("超过频率限制")
	}
	// 当前用户每小时限制
	if res, _ := s.limiter.Allow(userID, redis_rate.PerHour(config.C.VerifyCode.PerHour)); !res.Allowed {
		return errors.New("超过频率限制")
	}
	// 当前用户同一手机号限制
	if res, _ := s.limiter.Allow(fmt.Sprintf("%s%s", userID, mobile), redis_rate.PerHour(config.C.VerifyCode.PerMinute)); !res.Allowed {
		return errors.New("超过频率限制")
	}
	// 当前用户每天限制
	if res, _ := s.limiter.AllowN(userID, PerDay(config.C.VerifyCode.PerDay), 1); !res.Allowed {
		return errors.New("超过频率限制")
	}
	// 同一手机号每天限制
	if res, _ := s.limiter.AllowN(mobile, PerDay(config.C.VerifyCode.PerDay), 1); !res.Allowed {
		return errors.New("超过频率限制")
	}

	vc := config.C.VerifyCode
	verifyCode := s.GenVerifyCode(vc.CodeLength)

	s.logger.Debug("生成验证码", zap.String("userID", userID),
		zap.String("mobile", util.GetMaskedMobile(mobile)),
		zap.String("verify_code", verifyCode))

	// 保存短信验证码
	if err := s.SaveVerifyCode(userID, mobile, verifyCode, time.Duration(vc.ExpireMinute)*time.Minute); err != nil {
		return err
	}

	params := []string{verifyCode, strconv.Itoa(vc.ExpireMinute)}
	if err := s.SendTemplateSms(mobile, config.C.VerifyCode.TplId, params); err != nil {
		return err
	}

	return nil
}

// SaveVerifyCode 保存短信验证码到redis
// key 验证码保存时的key
// code 验证码
func (s SmsService) SaveVerifyCode(userID, mobile, code string, expiration time.Duration) error {
	key := s.BuildVerifyCodeKey(userID, mobile, code)
	s.logger.Debug("save verify code", zap.String("key", key), zap.String("code", code))

	if err := redis.Client.Set(key, code, expiration).Err(); err != nil {
		s.logger.Error("save_verify_code failed", zap.Error(err))
		return err
	}
	return nil
}

func (s SmsService) CheckVerifyCode(userID, mobile, code string) (bool, error) {
	key := s.BuildVerifyCodeKey(userID, mobile, code)
	s.logger.Debug("check verify code", zap.String("key", key))

	// 使用 mock
	if config.C.VerifyCode.MockCode != "" && config.C.VerifyCode.MockCode == code {
		return true, nil
	}

	expectCode, err := redis.Client.Get(key).Result()
	if err != nil {
		s.logger.Error("get_verify_code failed", zap.Error(err), zap.String("key", key))
		if err == libredis.Nil {
			return false, nil
		}
		return false, err
	}

	// 通过验证清除记录
	if expectCode == code {
		redis.Client.Del(key)
	}

	return expectCode == code, nil
}

// 生成验证码字符串
func (s SmsService) GenVerifyCode(codeLength int) (code string) {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < codeLength; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func (s SmsService) BuildVerifyCodeKey(userID, mobile, verifyCode string) string {
	return fmt.Sprintf("verify_code:%s:%s:%s", userID, mobile, verifyCode)
}

// 发送模板短信
func (s SmsService) SendTemplateSms(mobile string, tplId int, params []string) error {
	sms := config.C.Sms
	sender := qcloudsms.NewSmsSingleSender(sms.AppId, sms.AppSecret, sms.URL)

	var callback = func(err error, resp *http.Response, resData string) {
		if err != nil {
			s.logger.Error("err: ", zap.Error(err))
		} else {
			s.logger.Debug("response data: ", zap.String("resData", resData))
		}
	}

	if err := sender.SendWithParam(86, mobile, tplId, params, sms.Sign, "", "", callback); err != nil {
		return err
	}

	return nil
}
