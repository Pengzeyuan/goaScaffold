package dao

//
//import (
//	"boot/model"
//	"github.com/jinzhu/gorm"
//	"go.uber.org/zap"
//	"gopkg.in/resty.v1"
//	"strings"
//	"time"
//)
//
//type ThirdPartDao interface {
//
//	// 获取大厅管理系统情况
//	SaveHallManagement(value *model.HallManagementInfo) error
//
//	// gorm 关联查询
//	GormRelationSearch(queryModel model.CommonQueryModel) ([]*model.LegalPersonUser, error)
//}
//
//type ThirdPartDaoImpl struct {
//	db         *gorm.DB
//	logger     *zap.Logger
//	httpClient *resty.Client
//}
//
//func (d ThirdPartDaoImpl) GormRelationSearch(queryModel model.CommonQueryModel) ([]*model.LegalPersonUser, error) {
//
//	legalUsers := []*model.LegalPersonUser{}
//
//	sqlStr := "INSERT INTO animals (id,name,age) VALUES "
//	vals := []interface{}{}
//	const rowSQL = "(?,?, ?)"
//	var inserts []string
//
//	reportForms := []model.LegalPersonUser{{7, "牛魔", 400}, {8, "猴子", 500}}
//	for _, elem := range reportForms {
//		inserts = append(inserts, rowSQL)
//		vals = append(vals, elem.ID, elem.Name, elem.Age)
//	}
//	sqlStr = sqlStr + strings.Join(inserts, ",")
//
//	if err := d.db.Exec(sqlStr, vals...).Error; err != nil {
//		d.logger.Error(" insert is failed", zap.Error(err))
//		return nil, err
//	}
//
//	//在查询 User 时希望把 Company 的信息也一并查询, 有以下三种方法:
//	if err := d.db.Find(&legalUsers).Error; err != nil {
//		d.logger.Error(" Find legalUsers is failed", zap.Error(err))
//		return nil, err
//	}
//
//	for i := 0; i < len(legalUsers); i++ {
//		d.db.Model(&legalUsers[i]).Related(&legalUsers[i].Companies).Find(&legalUsers[i].Companies)
//	}
//
//	return legalUsers, nil
//}
//
//func NewThirdPartDaoImpl(db *gorm.DB, logger *zap.Logger) ThirdPartDao {
//	httpClient := resty.New().SetTimeout(1000 * time.Millisecond)
//	return ThirdPartDaoImpl{db: db, logger: logger, httpClient: httpClient}
//}
//
////  存储大厅窗口设立情况信息接口数据
//func (d ThirdPartDaoImpl) SaveHallManagement(value *model.HallManagementInfo) error {
//
//	if err := d.db.Exec(
//		"INSERT INTO `hall_management_info`(`created_at`,`updated_at`,`deleted_at`,`card_num`,`name`,`ou_name`) "+
//			"VALUES (?, ?, ?, ?, ?,   ?) "+
//			"ON DUPLICATE KEY UPDATE updated_at = ?, deleted_at = ?, card_num = ?, name = ?, ou_name = ?; ",
//
//		value.Model.CreatedAt, value.Model.UpdatedAt, value.Model.DeletedAt, value.CardNum, value.Name, value.OuName,
//
//		value.Model.UpdatedAt, value.Model.DeletedAt, value.CardNum, value.Name, value.OuName,
//	).Error; err != nil {
//		d.logger.Error(" Save HallManagement is failed", zap.Error(err))
//		return err
//	}
//	return nil
//}
