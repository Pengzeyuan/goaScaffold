package service

import (
	"context"

	"boot/dao"
	"boot/model"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type SimulationSVC interface {
	// 获取模拟数据
	Get(key string) (*model.Simulation, error)
	// 获取序列化后的模拟数据
	GetByUnmarshal(key string, data interface{}) (bool, error)
	// 设置模拟数据
	Set(key string, val string, isMock bool, orderBy, orderTimeScope int) (*model.Simulation, error)
}

type SimulationSVCImpl struct {
	simulationDao dao.SimulationDao
	logger        *zap.Logger
	ctx           context.Context
}

func NewSimulationSVCImpl(ctx context.Context, db *gorm.DB, logger *zap.Logger) SimulationSVC {
	simulationDao := dao.NewSimulationDaoImpl(db, logger)
	return SimulationSVCImpl{
		simulationDao: simulationDao,
		logger:        logger,
		ctx:           ctx,
	}
}

func (svc SimulationSVCImpl) GetByUnmarshal(key string, data interface{}) (bool, error) {
	isMock, err := svc.simulationDao.GetByUnmarshal(key, data)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		svc.logger.Error("get simulation failed", zap.Error(err))
		return false, err
	}
	return isMock, nil
}

func (svc SimulationSVCImpl) Get(key string) (*model.Simulation, error) {
	res, err := svc.simulationDao.Get(key)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		svc.logger.Error("get simulation failed", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (svc SimulationSVCImpl) Set(key string, val string, isMock bool, orderBy, orderTimeScope int) (*model.Simulation, error) {
	_, err := svc.simulationDao.Get(key)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		svc.logger.Error("get simulation failed", zap.Error(err))
		return nil, err
	}
	if err != nil && gorm.IsRecordNotFoundError(err) {
		// 创建模拟数据
		res, err := svc.simulationDao.Create(key, val, isMock, orderBy, orderTimeScope)
		if err != nil {
			svc.logger.Error("create simulation failed", zap.Error(err))
			return nil, err
		}
		return res, err
	}
	// 更新模拟数据
	fields := make(map[string]interface{})
	fields["val"] = val
	fields["is_show_mock"] = isMock
	if orderBy != 0 {
		fields["order_by"] = orderBy
	}
	if orderTimeScope != 0 {
		fields["order_time_scope"] = orderTimeScope
	}
	res, err := svc.simulationDao.Update(key, fields)
	if err != nil {
		svc.logger.Error("update simulation failed", zap.Error(err))
		return nil, err
	}
	return res, nil
}
