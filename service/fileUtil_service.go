package service

import (
	"boot/constant"
	"boot/dao"
	"boot/model"
	"bytes"
	"context"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type FileUtilService interface {
	// Import 导入数据
	Import(year, typ int, area string, fileByte []byte) error

	// 获取"四个办"统计数据(数量)
	GetFourCount(startYear, endYear int, Area string) (model.FourDoCount, error)
}

type FileUtilSVCImpl struct {
	db                 *gorm.DB
	logger             *zap.Logger
	ctx                context.Context
	newFileUtilDaoFunc dao.NewFourDaoFunc
}

func NewFileUtilSVCImpl(db *gorm.DB, logger *zap.Logger, ctx context.Context, newFileUtilDaoFunc dao.NewFourDaoFunc) FileUtilService {
	return FileUtilSVCImpl{
		db:                 db,
		logger:             logger,
		ctx:                ctx,
		newFileUtilDaoFunc: newFileUtilDaoFunc,
	}
}

func (svc FileUtilSVCImpl) Import(year, typ int, area string, fileByte []byte) error {
	fourDao := svc.newFileUtilDaoFunc(svc.db, svc.logger)
	// 转换导入文件
	openExcl, err := excelize.OpenReader(bytes.NewReader(fileByte))
	if err != nil {
		svc.logger.Error("open file failed", zap.Error(err))
		return err
	}

	// 删除相同条件数据
	fileUtilScopes := dao.NewFileUtilScopeConstructor()
	fileUtilScopes = fileUtilScopes.AddYear(year)
	if area != "" {
		fileUtilScopes = fileUtilScopes.AddArea(area)
	}

	switch typ {
	case constant.ImmediateInfo:
		err := fourDao.DeleteImmediateInfo(fileUtilScopes)
		if err != nil {
			svc.logger.Error("delete immediateInfo failed", zap.Error(err))
			return err
		}
	case constant.OnlineInfo:
		err := fourDao.DeleteOnlineInfo(fileUtilScopes)
		if err != nil {
			svc.logger.Error("delete onlineInfo failed", zap.Error(err))
			return err
		}
	case constant.NearbyInfo:
		err := fourDao.DeleteNearbyInfo(fileUtilScopes)
		if err != nil {
			svc.logger.Error("delete nearbyInfo failed", zap.Error(err))
			return err
		}
	case constant.OnceInfo:
		err := fourDao.DeleteOnceInfo(fileUtilScopes)
		if err != nil {
			svc.logger.Error("delete onfo failed", zap.Error(err))
			return err
		}
	}

	//  加入
	rows := openExcl.GetRows(openExcl.GetSheetName(openExcl.GetActiveSheetIndex()))

	if typ == constant.ImmediateInfo { // 导入"马上办"
		for _, row := range rows[2:] {
			immediateInfo := model.ImmediateInfo{
				Model:          gorm.Model{},
				DeptName:       row[1],
				ParItemName:    row[2],
				ChiItemName:    row[3],
				ManagedObj:     row[5],
				ItemType:       row[6],
				Implementation: row[7],
				Year:           year,
				Area:           area,
			}
			pre, err := strconv.Atoi(strings.TrimSpace(row[4]))
			if err != nil {
				svc.logger.Error("转换出错", zap.Error(err))
				return err
			}
			immediateInfo.PresentTimes = pre
			_, err = fourDao.CreateImmediateInfo(&immediateInfo)
			if err != nil {
				svc.logger.Error("保存出错", zap.Error(err))
				return err
			}
		}
	}
	if typ == constant.OnlineInfo { // 导入"网上办"信息
		for _, row := range rows[2:] {
			onlineInfo := model.OnlineInfo{
				Model:          gorm.Model{},
				DeptName:       row[1],
				ParItemName:    row[2],
				ChiItemName:    row[3],
				ManagedObj:     row[5],
				ItemType:       row[6],
				Implementation: row[7],
				Year:           year,
				Area:           area,
			}
			pre, err := strconv.Atoi(strings.TrimSpace(row[4]))
			if err != nil {
				svc.logger.Error("转换出错", zap.Error(err))
				return err
			}
			onlineInfo.PresentTimes = pre
			_, err = fourDao.CreateOnlineInfo(&onlineInfo)
			if err != nil {
				svc.logger.Error("保存出错", zap.Error(err))
				return err
			}
		}
	}
	if typ == constant.NearbyInfo { // 导入"就近办"信息
		for _, row := range rows[2:] {
			nearbyInfo := model.NearbyInfo{
				Model:          gorm.Model{},
				OrgName:        row[1],
				ParItemName:    row[2],
				ChiItemName:    row[3],
				AdmissibleArea: row[4],
				Implementation: row[5],
				Year:           year,
				Area:           area,
			}
			_, err = fourDao.CreateNearbyInfo(&nearbyInfo)
			if err != nil {
				svc.logger.Error("保存出错", zap.Error(err))
				return err
			}
		}
	}
	if typ == constant.OnceInfo { // 导入"一次办"信息
		for _, row := range rows[2:] {
			onceInfo := model.OnceInfo{
				Model:          gorm.Model{},
				DeptName:       row[1],
				ParItemName:    row[2],
				ChiItemName:    row[3],
				ManagedObj:     row[5],
				ItemType:       row[6],
				Implementation: row[7],
				Year:           year,
				Area:           area,
			}
			pre, err := strconv.Atoi(strings.TrimSpace(row[4]))
			if err != nil {
				svc.logger.Error("转换出错", zap.Error(err))
				return err
			}
			onceInfo.PresentTimes = pre
			_, err = fourDao.CreateOnceInfo(&onceInfo)
			if err != nil {
				svc.logger.Error("保存出错", zap.Error(err))
				return err
			}
		}
	}
	return nil
}

func (svc FileUtilSVCImpl) GetFourCount(startYear, endYear int, Area string) (model.FourDoCount, error) {
	var res model.FourDoCount
	fourDao := svc.newFileUtilDaoFunc(svc.db, svc.logger)
	// 设置范围 where条件
	fourScopes := dao.NewFileUtilScopeConstructor()
	if Area != "" {
		fourScopes = fourScopes.AddArea(Area)
	}
	//1、及时办
	immediateInfoCount, err := fourDao.GetImmediateInfoNum(startYear, endYear, fourScopes)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		svc.logger.Error("get immediateInfoCount failed", zap.Error(err))
		return res, err
	}
	//2、网上办
	onlineInfoCount, err := fourDao.GetOnlineInfoNum(startYear, endYear, fourScopes)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		svc.logger.Error("get onlineInfoCount failed", zap.Error(err))
		return res, err
	}
	//3、就近办
	nearbyInfoCount, err := fourDao.GetNearbyInfoNum(startYear, endYear, fourScopes)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		svc.logger.Error("get nearbyInfoCount failed", zap.Error(err))
		return res, err
	}
	//4、一次办
	onceInfoCount, err := fourDao.GetOnceInfoNum(startYear, endYear, fourScopes)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		svc.logger.Error("get onceInfoCount failed", zap.Error(err))
		return res, err
	}
	res.ImmediateInfoCount = immediateInfoCount
	res.OnlineInfoCount = onlineInfoCount
	res.NearbyInfoCount = nearbyInfoCount
	res.OnceInfoCount = onceInfoCount

	return res, nil

}
