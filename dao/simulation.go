package dao

import (
	"boot/model"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"log"
)

type SimulationDao interface {
	// 获取模拟数据
	Get(key string) (*model.Simulation, error)
	// 获取模拟数据并且序列化
	GetByUnmarshal(key string, data interface{}) (bool, error)
	// 更新模拟数据
	Update(key string, fields map[string]interface{}) (*model.Simulation, error)
	// 创建模拟数据
	Create(key string, val string, isMock bool, orderBy, orderTimeScope int) (*model.Simulation, error)
}
type SimulationDaoImpl struct {
	db     *gorm.DB
	logger *zap.Logger
}

type NewSimulationDaoFunc = func(*gorm.DB, *log.Logger) SimulationDao

func NewSimulationDaoImpl(db *gorm.DB, logger *zap.Logger) SimulationDao {
	return &SimulationDaoImpl{
		db:     db,
		logger: logger,
	}
}

func (d SimulationDaoImpl) Create(key string, val string, isMock bool, orderBy, orderTimeScope int) (*model.Simulation, error) {

	res := model.Simulation{
		BaseModel:  model.BaseModel{},
		Key:        key,
		Val:        val,
		IsShowMock: isMock,
	}
	if orderBy != 0 {
		res.OrderBy = orderBy
	}
	if orderTimeScope != 0 {
		res.OrderTimeScope = orderTimeScope
	}
	err := d.db.Create(&res).Scan(&res).Error
	return &res, err
}

func (d SimulationDaoImpl) Update(key string, fields map[string]interface{}) (*model.Simulation, error) {
	res := model.Simulation{}
	err := d.db.Model(&res).Where("`key` = ?", key).UpdateColumns(fields).Scan(&res).Error
	return &res, err
}

func (d SimulationDaoImpl) GetByUnmarshal(key string, data interface{}) (bool, error) {
	res := model.Simulation{}
	if err := d.db.Where("`key` = ?", key).Find(&res).Error; err != nil {
		d.logger.Error("get simulation failed", zap.Error(err))
		return false, err
	}
	if err := json.Unmarshal([]byte(res.Val), data); err != nil {
		d.logger.Error("unmarshal failed", zap.Error(err))
		return false, err
	}
	return res.IsShowMock, nil
}

func (d SimulationDaoImpl) Get(key string) (*model.Simulation, error) {
	res := model.Simulation{}
	err := d.db.Where("`key` = ?", key).Find(&res).Error
	return &res, err
}
