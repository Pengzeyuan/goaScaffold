package boot

import (
	entityhall "boot/gen/entity_hall"
	log "boot/gen/log"
	"context"
)

// entity_hall service example implementation.
// The example methods log the requests and return zero values.
type entityHallsrvc struct {
	logger *log.Logger
}

// NewEntityHall returns the entity_hall service implementation.
func NewEntityHall(logger *log.Logger) entityhall.Service {
	return &entityHallsrvc{logger}
}

// 排号总览
func (s *entityHallsrvc) WaitLineOverview(ctx context.Context, p *entityhall.WaitLineOverviewPayload) (res *entityhall.WaitLineOverviewResult, err error) {
	res = &entityhall.WaitLineOverviewResult{}

	resp := entityhall.WaitLineOverviewResp{TodayDQ: 100, CumulativeDQ: 200}
	res.Data = &resp
	s.logger.Info("entityHall.WaitLineOverview")
	return res, nil
}
