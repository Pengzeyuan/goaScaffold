package dao

import (
	"boot/model"
	libsgorm "git.chinaopen.ai/yottacloud/go-libs/gorm"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type FileUtilDao interface {

	// 删除"马上办"数量统计
	DeleteImmediateInfo(scopes FileUtilScopeConstructor) error
	// 删除"网上办"数量统计
	DeleteOnlineInfo(scopes FileUtilScopeConstructor) error
	// 删除"就近办"数量统计
	DeleteNearbyInfo(scopes FileUtilScopeConstructor) error
	// 删除"一次办"数量统计
	DeleteOnceInfo(scopes FileUtilScopeConstructor) error

	// 创建"马上办"信息
	CreateImmediateInfo(information *model.ImmediateInfo) (model.ImmediateInfo, error)
	// 创建"网上办"信息
	CreateOnlineInfo(information *model.OnlineInfo) (model.OnlineInfo, error)
	// 创建"就近办"信息
	CreateNearbyInfo(information *model.NearbyInfo) (model.NearbyInfo, error)
	// 创建"一次办"信息
	CreateOnceInfo(information *model.OnceInfo) (model.OnceInfo, error)

	// "马上办"数量统计
	GetImmediateInfoNum(startYear, endYear int, scopes FileUtilScopeConstructor) ([]*model.FileCount, error)
	// "网上办"数量统计
	GetOnlineInfoNum(startYear, endYear int, scopes FileUtilScopeConstructor) ([]*model.FileCount, error)
	// "就近办"数量统计
	GetNearbyInfoNum(startYear, endYear int, scopes FileUtilScopeConstructor) ([]*model.FileCount, error)
	// "一次办"数量统计
	GetOnceInfoNum(startYear, endYear int, scopes FileUtilScopeConstructor) ([]*model.FileCount, error)
}

type FileUtilDaoImpl struct {
	db     *gorm.DB
	logger *zap.Logger
}

// 一个方法类型  可让service层在 controller层 插槽式调用db数据源
type NewFourDaoFunc = func(*gorm.DB, *zap.Logger) FileUtilDao

func NewFileUtilDaoImpl(db *gorm.DB, logger *zap.Logger) FileUtilDao {
	return FileUtilDaoImpl{
		db:     db,
		logger: logger,
	}
}

func (f FileUtilDaoImpl) DeleteImmediateInfo(scopes FileUtilScopeConstructor) error {
	var res model.ImmediateInfo
	db := f.db.Model(&res).Scopes(scopes.Scopes()...)
	if err := db.Unscoped().Delete(&res).Error; err != nil {
		f.logger.Error("delete immediateInfo failed", zap.Error(err))
		return err
	}
	return nil
}

func (f FileUtilDaoImpl) DeleteOnlineInfo(scopes FileUtilScopeConstructor) error {
	var res model.OnlineInfo
	db := f.db.Model(&res).Scopes(scopes.Scopes()...)
	if err := db.Unscoped().Delete(&res).Error; err != nil {
		f.logger.Error("delete onlineInfo failed", zap.Error(err))
		return err
	}
	return nil
}

func (f FileUtilDaoImpl) DeleteNearbyInfo(scopes FileUtilScopeConstructor) error {
	var res model.NearbyInfo
	db := f.db.Model(&res).Scopes(scopes.Scopes()...)
	if err := db.Unscoped().Delete(&res).Error; err != nil {
		f.logger.Error("delete nearbyInfo failed", zap.Error(err))
		return err
	}
	return nil
}

func (f FileUtilDaoImpl) DeleteOnceInfo(scopes FileUtilScopeConstructor) error {
	var res model.OnceInfo
	db := f.db.Model(&res).Scopes(scopes.Scopes()...)
	if err := db.Unscoped().Delete(&res).Error; err != nil {
		f.logger.Error("delete onceInfo failed", zap.Error(err))
		return err
	}
	return nil
}

func (f FileUtilDaoImpl) CreateImmediateInfo(information *model.ImmediateInfo) (model.ImmediateInfo, error) {
	if err := f.db.Create(information).Error; err != nil {
		f.logger.Error("create immediate info failed", zap.Error(err))
		return model.ImmediateInfo{}, err
	}
	return *information, nil
}

func (f FileUtilDaoImpl) CreateOnlineInfo(information *model.OnlineInfo) (model.OnlineInfo, error) {
	if err := f.db.Create(information).Error; err != nil {
		f.logger.Error("create nearby_info failed", zap.Error(err))
		return model.OnlineInfo{}, err
	}
	return *information, nil
}

func (f FileUtilDaoImpl) CreateNearbyInfo(information *model.NearbyInfo) (model.NearbyInfo, error) {
	if err := f.db.Create(information).Error; err != nil {
		f.logger.Error("create once_info failed", zap.Error(err))
		return model.NearbyInfo{}, err
	}
	return *information, nil
}

func (f FileUtilDaoImpl) CreateOnceInfo(information *model.OnceInfo) (model.OnceInfo, error) {
	if err := f.db.Create(information).Error; err != nil {
		f.logger.Error("create online_info failed", zap.Error(err))
		return model.OnceInfo{}, err
	}
	return *information, nil
}

func (f FileUtilDaoImpl) GetImmediateInfoNum(startYear, endYear int, scopes FileUtilScopeConstructor) ([]*model.FileCount, error) {
	var res []*model.FileCount
	table := model.ImmediateInfo{}.TableName()
	querySQL := `immediate_info.year,count(*) as count`
	query := f.db.Table(table).Scopes(scopes.Scopes()...)
	err := query.Select(querySQL).Where("immediate_info.year BETWEEN ? AND ? ", startYear, endYear).
		Group("immediate_info.year").
		Scan(&res).Error
	if err != nil {
		f.logger.Error("get immediate_info failed", zap.Error(err))
		return res, err
	}
	return res, nil
}

func (f FileUtilDaoImpl) GetOnlineInfoNum(startYear, endYear int, scopes FileUtilScopeConstructor) ([]*model.FileCount, error) {
	var res []*model.FileCount
	table := model.OnlineInfo{}.TableName()
	querySQL := `online_info.year,count(*) as count`
	query := f.db.Table(table).Scopes(scopes.Scopes()...)
	err := query.Select(querySQL).Where("online_info.year BETWEEN ? AND ? ", startYear, endYear).
		Group("online_info.year").
		Scan(&res).Error
	if err != nil {
		f.logger.Error("get online_info failed", zap.Error(err))
		return res, err
	}
	return res, nil
}

func (f FileUtilDaoImpl) GetNearbyInfoNum(startYear, endYear int, scopes FileUtilScopeConstructor) ([]*model.FileCount, error) {
	var res []*model.FileCount
	table := model.NearbyInfo{}.TableName()
	querySQL := `nearby_info.year,count(*) as count`
	query := f.db.Table(table).Scopes(scopes.Scopes()...)
	err := query.Select(querySQL).Where("nearby_info.year BETWEEN ? AND ? ", startYear, endYear).
		Group("nearby_info.year").
		Scan(&res).Error
	if err != nil {
		f.logger.Error("get nearby_info failed", zap.Error(err))
		return res, err
	}
	return res, nil
}

func (f FileUtilDaoImpl) GetOnceInfoNum(startYear, endYear int, scopes FileUtilScopeConstructor) ([]*model.FileCount, error) {
	var res []*model.FileCount
	querySQL := `once_info.year,count(*) as count`
	table := model.OnceInfo{}.TableName()
	query := f.db.Table(table).Scopes(scopes.Scopes()...)
	err := query.Select(querySQL).Where("once_info.year BETWEEN ? AND ? ", startYear, endYear).
		Group("once_info.year").
		Scan(&res).Error
	if err != nil {
		f.logger.Error("get once_info failed", zap.Error(err))
		return res, err
	}
	return res, nil
}

func NewFileUtilScopeConstructor() FileUtilScopeConstructor {
	return FileUtilScopeConstructor{
		scopes: []libsgorm.Scope{},
	}
}

type FileUtilScopeConstructor struct {
	scopes []libsgorm.Scope
}

func (f FileUtilScopeConstructor) Scopes() []libsgorm.Scope {
	return f.scopes
}
func (f FileUtilScopeConstructor) AddYear(year int) FileUtilScopeConstructor {
	query := func(db *gorm.DB) *gorm.DB {
		return db.Where("year = ?", year)
	}
	f.scopes = append(f.scopes, query)
	return f
}

func (f FileUtilScopeConstructor) AddArea(area string) FileUtilScopeConstructor {
	query := func(db *gorm.DB) *gorm.DB {
		return db.Where("area = ?", area)
	}
	f.scopes = append(f.scopes, query)
	return f
}

func (f FileUtilScopeConstructor) AddPk(pk uuid.UUID) FileUtilScopeConstructor {
	query := func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", pk)
	}
	f.scopes = append(f.scopes, query)
	return f
}
