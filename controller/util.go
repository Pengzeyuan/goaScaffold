package boot

import (
	"boot/model"
	"boot/utils"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"boot/config"
	"boot/gen/log"

	"go.uber.org/zap"
	"goa.design/goa/v3/middleware"
	goa "goa.design/goa/v3/pkg"
)

func getErrorID(ctx context.Context) string {
	reqID, ok := ctx.Value(middleware.RequestIDKey).(string)
	if !ok {
		return goa.NewErrorID()
	}

	return reqID
}

func L(ctx context.Context, logger *log.Logger) *zap.Logger {
	reqID, ok := ctx.Value(middleware.RequestIDKey).(string)
	if ok {
		return logger.Desugar().With(zap.String("reqID", reqID))
	}

	return logger.Desugar()
}

// 创建内部错误
func MakeInternalServerError(ctx context.Context, errmsg string) *goa.ServiceError {
	if errmsg == "" || !config.C.Debug {
		errmsg = "服务器开小差了，稍后再试吧"
	}

	return &goa.ServiceError{
		Name:    "internal_server_error",
		ID:      getErrorID(ctx),
		Message: errmsg,
		Fault:   true,
	}
}

// 创建登录错误
func MakeLoginError(ctx context.Context, errmsg string) *goa.ServiceError {
	if errmsg == "" {
		errmsg = "服务器开小差了，稍后再试吧"
	}

	return &goa.ServiceError{
		Name:    "internal_server_error",
		ID:      getErrorID(ctx),
		Message: errmsg,
		Fault:   true,
	}
}

// 创建参数错误
func MakeBadRequestError(ctx context.Context, errmsg string) *goa.ServiceError {
	if errmsg == "" {
		errmsg = "参数错误"
	}

	return &goa.ServiceError{
		Name:    "bad_request",
		ID:      getErrorID(ctx),
		Message: errmsg,
	}
}

func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "bad_request",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorizedError(ctx context.Context, errmsg string) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "unauthorized",
		ID:      getErrorID(ctx),
		Message: errmsg,
	}
}

// MakeForbidden builds a goa.ServiceError from an error.
func MakeForbiddenError(ctx context.Context, errmsg string) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "forbidden",
		ID:      getErrorID(ctx),
		Message: errmsg,
	}
}

func GetRedisData(queryModel *model.CommonQueryModel, redisQueryModel model.RedisQueryModel,
	storage *utils.RedisStorage, resData interface{}) error {
	// 根据结束时间获取查询的年份
	// 获取最近一年前日期
	currentYear, currentMonth, _ := time.Now().Date()
	endTime, err := time.Parse("2006-01", queryModel.EndDate)

	if err == nil {
		currentYear, currentMonth, _ = endTime.Date()
	} else {
		zap.L().Error("parse date fail", zap.Error(err))
	}

	// 将时间处理为年-月-日 从redis里面查找
	queryModel.StartDate = fmt.Sprintf("%d-01-01", currentYear)
	queryModel.EndDate = fmt.Sprintf("%d-%02d-01", currentYear, currentMonth)
	resDataJson, err := storage.GetValue(redisQueryModel.ModelName, redisQueryModel.MethodName, queryModel)

	// 处理从dao查询时 日期格式为 yyyy-mm-dd
	queryModel.StartDate = fmt.Sprintf("%d-01-01", currentYear)

	if currentMonth < 12 {
		//查询4月 = 1月1日到5月1日的数据
		currentMonth++
		queryModel.EndDate = fmt.Sprintf("%d-%02d-01", currentYear, currentMonth)
	} else {
		queryModel.EndDate = fmt.Sprintf("%d-12-31", currentYear)
	}

	// 缓存存在
	if err == nil {
		err := json.Unmarshal([]byte(resDataJson), resData)
		if err != nil {
			zap.L().Error("json to res data struct fail", zap.Error(err))
			return err
		}
	} else {
		zap.L().Error("get redis data fail", zap.Error(err))
		return err
	}

	return nil
}

func SaveRedisData(queryModel *model.CommonQueryModel, redisQueryModel model.RedisQueryModel,
	storage *utils.RedisStorage, resDataPtr interface{}) {
	// 根据结束时间获取查询的年份
	// 获取最近一年前日期
	currentYear, currentMonth, _ := time.Now().Date()
	endTime, err := time.Parse("2006-01-02", queryModel.EndDate)

	if err == nil {
		currentYear, currentMonth, _ = endTime.Date()
	}

	// 处理从redis 存储时 日期格式为 yyyy-mm-dd
	queryModel.StartDate = fmt.Sprintf("%d-01-01", currentYear)
	//查询如果之前月份没有加 1  月数不变
	if strings.Split(queryModel.EndDate, "-")[2] != "31" {
		currentMonth--
	}

	queryModel.EndDate = fmt.Sprintf("%d-%02d-01", currentYear, currentMonth)

	resDataJsonBytes, err := json.Marshal(resDataPtr)
	if err != nil {
		zap.L().Error("res data to bytes failed", zap.Error(err))
	}

	err = storage.SaveValue(redisQueryModel.ModelName, redisQueryModel.MethodName,
		utils.Bytes2String(resDataJsonBytes), queryModel)
	if err != nil {
		zap.L().Error("save redis data failed", zap.Error(err))
	}
}
