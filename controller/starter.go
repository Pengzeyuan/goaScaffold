package controller

import (
	"context"
	log "starter/gen/log"
	starter "starter/gen/starter"
)

// starter service example implementation.
// The example methods log the requests and return zero values.
type startersrvc struct {
	logger *log.Logger
}

// NewStarter returns the starter service implementation.
func NewStarter(logger *log.Logger) starter.Service {
	return &startersrvc{logger}
}

// 使用账号密码登录
func (s *startersrvc) LoginByUsername(ctx context.Context, p *starter.LoginByUsernamePayload) (res *starter.LoginByUsernameResult, err error) {
	res = &starter.LoginByUsernameResult{}
	logger := L(ctx, s.logger)
	logger.Info("starter.LoginByUsername")

	return
}
