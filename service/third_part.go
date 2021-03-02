package service

//
//import (
//	"boot/dao"
//	"boot/model"
//	"context"
//	"github.com/jinzhu/gorm"
//	"go.uber.org/zap"
//)
//
//// 大厅排队办事实时图service
//type ThirdPartSVC interface {
//
//	// gorm 关联查询
//	GormRelationSearch(queryModel model.CommonQueryModel) ([]*model.LegalPersonUser, error)
//}
//
//type ThirdPartSVCSVCImpl struct {
//	logger       *zap.Logger
//	ctx          context.Context
//	thirdPartDao dao.ThirdPartDao
//}
//
//func NewThirdPartSVCImpl(ctx context.Context, db *gorm.DB, logger *zap.Logger) ThirdPartSVC {
//	thirdPartDao := dao.NewThirdPartDaoImpl(db, logger)
//	return ThirdPartSVCSVCImpl{
//		ctx:          ctx,
//		logger:       logger,
//		thirdPartDao: thirdPartDao,
//	}
//}
//
//func (t ThirdPartSVCSVCImpl) GormRelationSearch(queryModel model.CommonQueryModel) ([]*model.LegalPersonUser, error) {
//
//	res, err := t.thirdPartDao.GormRelationSearch(queryModel)
//	if err != nil {
//		t.logger.Error("GormRelationSearch is failed:", zap.Error(err))
//		return nil, err
//	}
//	return res, nil
//}
