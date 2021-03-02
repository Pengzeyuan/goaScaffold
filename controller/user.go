package boot

import (
	"boot/dao"
	"boot/serializer"
	"boot/service"
	"boot/utils"
	"context"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"net/http"
	"time"

	log "boot/gen/log"
	user "boot/gen/user"
)

// User service example implementation.
// The example methods log the requests and return zero values.
type usersrvc struct {
	Auther
	logger *log.Logger
}

// NewUser returns the User service implementation.
func NewUser(logger *log.Logger) user.Service {
	return &usersrvc{
		logger: logger,
		Auther: Auther{
			logger: logger,
		},
	}
}

// 使用账号密码登录
func (s *usersrvc) LoginByUsername(ctx context.Context, p *user.LoginByUsernamePayload) (res *user.LoginByUsernameResult, err error) {
	res = &user.LoginByUsernameResult{}
	logger := L(ctx, s.logger)
	logger.Info("user.LoginByUsername")

	tx := dao.DpDB.Begin()
	if tx.Error != nil {
		_ = tx.Rollback()
		logger.Error("begin tx failed", zap.Error(tx.Error))
		return nil, MakeInternalServerError(ctx, "内部服务器错误")
	}
	userServer := service.NewUserSVCImpl(tx, logger, dao.NewUserDaoImpl)
	loginUser, err := userServer.GetAccountByUserName(p.Username)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			logger.Error("get user failed", zap.Error(err))
			return nil, MakeLoginError(ctx, "账号不存在")
		}
		logger.Error("get user failed", zap.Error(err))
		return nil, MakeLoginError(ctx, "获取用户失败")
	}
	// 校验密码
	if !loginUser.CheckPassword(p.Password, loginUser.Password) {
		logger.Error("check password failed", zap.Error(err))
		return nil, MakeLoginError(ctx, "密码错误")
	}

	// 更新登陆时间
	if err = userServer.UpdateLoginTime(loginUser.UserName); err != nil {
		_ = tx.Rollback()
		s.logger.Error("update login time failed", zap.Error(err))
		return nil, MakeLoginError(ctx, "更新登陆时间失败")
	}
	// 事务提交
	if err := tx.Commit().Error; err != nil {
		_ = tx.Rollback()
		logger.Error("commit tx failed", zap.Error(err))
		return nil, MakeInternalServerError(ctx, "内部服务器错误")
	}

	// 生成返回的数据
	expireAt := time.Now().Add(utils.JwtMaxLifetime)
	tokenStr, err := utils.GenToken(string(loginUser.ID), expireAt)
	if err != nil {
		s.logger.Error("gen login token failed", zap.Error(err))
		return nil, MakeLoginError(ctx, "获取token失败")
	}
	expireIn := int(utils.JwtMaxLifetime.Seconds())
	// 同时写入 cookie
	cookie := http.Cookie{
		Name:    utils.JwtTokenCookieName,
		Value:   tokenStr,
		Path:    "/",
		Expires: expireAt,
		MaxAge:  expireIn,
	}

	users := serializer.ModelAccount2AuthAopUser(loginUser)
	credentials := &user.Credentials{
		Token:     tokenStr,
		ExpiresIn: expireIn,
	}
	newUser := &user.Session{
		User:        users,
		Credentials: credentials,
		Cookie:      cookie.String(),
	}
	res.Data = newUser
	return res, nil
}

// 使用短信验证码登录
func (s *usersrvc) LoginBySmsCode(ctx context.Context, p *user.LoginBySmsCodePayload) (res *user.LoginBySmsCodeResult, err error) {
	res = &user.LoginBySmsCodeResult{}
	logger := L(ctx, s.logger)
	logger.Info("user.LoginBySmsCode")

	return
}

// 修改登录密码
func (s *usersrvc) UpdatePassword(ctx context.Context, p *user.UpdatePasswordPayload) (res *user.UpdatePasswordResult, err error) {
	res = &user.UpdatePasswordResult{}
	logger := L(ctx, s.logger)
	logger.Info("user.UpdatePassword")

	return
}

// 获取图形验证码
func (s *usersrvc) GetCaptchaImage(ctx context.Context) (res *user.GetCaptchaImageResult, err error) {
	res = &user.GetCaptchaImageResult{}
	logger := L(ctx, s.logger)
	logger.Info("user.GetCaptchaImage")

	return
}

// 发送短信验证码
func (s *usersrvc) SendSmsCode(ctx context.Context, p *user.SendSmsCodePayload) (res *user.SendSmsCodeResult, err error) {
	res = &user.SendSmsCodeResult{}
	logger := L(ctx, s.logger)
	logger.Info("user.SendSmsCode")

	return
}
