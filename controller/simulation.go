package boot

import (
	"boot/dao"
	log "boot/gen/log"
	simulation "boot/gen/simulation"
	"boot/serializer"
	"context"

	goalibs "git.chinaopen.ai/yottacloud/go-libs/goa-libs"

	"git.chinaopen.ai/yottacloud/go-libs/jwt"

	"go.uber.org/zap"

	"goa.design/goa/v3/security"
)

// simulation service example implementation.
// The example methods log the requests and return zero values.
type simulationsrvc struct {
	logger *log.Logger
}

// NewSimulation returns the simulation service implementation.
func NewSimulation(logger *log.Logger) simulation.Service {
	return &simulationsrvc{logger}
}

const (
	UserIDKey = "userId"
)

// JWTAuth implements the authorization logic for service "simulation" for the
// "jwt" security scheme.
func (s *simulationsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//

	validator := jwt.NewValidator()
	jwtClaims := &goalibs.ExtendedClaims{}
	// parse && verify JWT token,

	if _, err := validator.Verify(token, scheme); err != nil {
		s.logger.Error("验证JWT失败", zap.Error(err))
		//s.logger.Info(fmt.Sprintf("%s", userClaims.Name))
		return ctx, err
	}

	// nolint
	ctx = context.WithValue(ctx, UserIDKey, jwtClaims.UserID)
	return ctx, nil
	//return ctx, fmt.Errorf("not implemented")
}

// 设置数据
func (s *simulationsrvc) SetData(ctx context.Context, p *simulation.SetDataPayload) (res *simulation.SetDataResult, err error) {
	res = &simulation.SetDataResult{}
	logger := L(ctx, s.logger)
	logger.Info("simulation.SetData")

	tx := dao.DpDB.Begin()
	if tx.Error != nil {
		_ = tx.Rollback()
		logger.Error("begin tx failed", zap.Error(tx.Error))
		return nil, MakeInternalServerError(ctx, "服务器内部错误")
	}
	svc := simulationSVC(ctx, tx, logger)

	_, err = svc.Set(p.Key, p.Val, p.IsShowMock, p.OrderBy, p.OrderTimeScope)

	if err != nil {
		_ = tx.Rollback()
		logger.Error("set simulation failed", zap.Error(err))
		return nil, MakeInternalServerError(ctx, "设置模拟数据失败")
	}

	if err := tx.Commit().Error; err != nil {
		_ = tx.Rollback()
		logger.Error("commit tx failed", zap.Error(err))
		return nil, MakeInternalServerError(ctx, "服务器内部错误")
	}

	success := simulation.SuccessResult{}
	success.OK = true
	res.Result = &success
	return res, nil
}

// 获取模拟数据
func (s *simulationsrvc) GetData(ctx context.Context, p *simulation.GetDataPayload) (res *simulation.GetDataResult, err error) {
	res = &simulation.GetDataResult{}

	logger := L(ctx, s.logger)
	logger.Info("simulation.GetData")

	svc := simulationSVC(ctx, dao.DpDB, logger)
	data, err := svc.Get(p.Key)
	if err != nil {
		logger.Error("get  simulation failed", zap.Error(err))
		return nil, MakeInternalServerError(ctx, "正在对接开发中")
	}
	res.Data = serializer.SimulationModel2GetDataResp(data)
	return res, nil
}
