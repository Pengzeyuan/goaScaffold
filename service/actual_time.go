package service

import (
	"boot/constant"
	"boot/dao"
	"boot/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"time"
)

// 大厅排队办事实时图service
type HallActualTimeSVC interface {

	// 接收第三方推送数据--大厅排队办事实时图基础数据
	ReceiveThirdPartyPushData(queryModel model.CommonQueryModel) (model.ServiceActualTimeResp, error)

	// 存储WindowInfo第三方推送数据列表
	SaveWindowInfoDataList(valueList []model.WindowInfo) error
	// 存储TakeNumber第三方推送数据列表
	SaveTakeNumberDataList(valueList []model.TakeNumber) error
	// 存储CallNumber第三方推送数据列表
	SaveCallNumberDataList(valueList []model.CallNumber) error
}

type HallActualTimeSVCImpl struct {
	logger            *zap.Logger
	ctx               context.Context
	hallActualTimeDao dao.HallActualTimeDao
}

func NewHallActualTimeSVCImpl(ctx context.Context, db *gorm.DB, logger *zap.Logger) HallActualTimeSVC {
	hallActualTimeDao := dao.NewHallActualTimeDaoImpl(db, logger)
	return HallActualTimeSVCImpl{
		ctx:               ctx,
		logger:            logger,
		hallActualTimeDao: hallActualTimeDao,
	}
}

// 存储WindowInfo第三方推送数据列表
func (h HallActualTimeSVCImpl) SaveWindowInfoDataList(valueList []model.WindowInfo) error {
	actualTimeDao := h.hallActualTimeDao
	for _, v := range valueList {
		v.CreatedAt = time.Now()
		v.UpdatedAt = time.Now()

		err := actualTimeDao.SaveWindowInfoData(&v)
		if err != nil {
			h.logger.Error("SaveWindowInfoDataList is failed:", zap.Error(err))
			return err
		}
	}
	return nil
}

// 存储TakeNumber第三方推送数据列表
func (h HallActualTimeSVCImpl) SaveTakeNumberDataList(valueList []model.TakeNumber) error {
	actualTimeDao := h.hallActualTimeDao
	for _, v := range valueList {
		v.CreatedAt = time.Now()
		v.UpdatedAt = time.Now()

		err := actualTimeDao.SaveTakeNumberData(&v)
		if err != nil {
			h.logger.Error("SaveTakeNumberDataList is failed:", zap.Error(err))
			return err
		}
	}
	return nil
}

// 存储CallNumber第三方推送数据列表
func (h HallActualTimeSVCImpl) SaveCallNumberDataList(valueList []model.CallNumber) error {
	actualTimeDao := h.hallActualTimeDao
	for _, v := range valueList {
		v.CreatedAt = time.Now()
		v.UpdatedAt = time.Now()

		err := actualTimeDao.SaveCallNumberData(&v)
		if err != nil {
			h.logger.Error("SaveCallNumberDataList is failed:", zap.Error(err))
			return err
		}
	}
	return nil
}

// 接收第三方推送数据--大厅排队办事实时图基础数据
func (h HallActualTimeSVCImpl) ReceiveThirdPartyPushData(queryModel model.CommonQueryModel) (model.ServiceActualTimeResp, error) {
	res := model.ServiceActualTimeResp{}
	switch queryModel.Method {

	case constant.HallWindowInfo:
		var result []model.WindowInfo

		err := json.Unmarshal(queryModel.PullData, &result)
		if err != nil {
			h.logger.Error("ReceiveThirdPartyPushData unmarshall err:", zap.Error(err))
			return res, err
		}
		err = h.SaveWindowInfoDataList(result)
		if err != nil {
			h.logger.Error("SaveWindowInfoDataList failed", zap.Error(err))
			return res, err
		}

		res.Msg = fmt.Sprintf("%d 条WindowInfo数据接收成功", len(result))
	case constant.HallTakeNumber:
		var result []model.TakeNumber

		err := json.Unmarshal(queryModel.PullData, &result)
		if err != nil {
			h.logger.Error("ReceiveThirdPartyPushData unmarshall err:", zap.Error(err))
			return res, err
		}
		err = h.SaveTakeNumberDataList(result)
		if err != nil {

			h.logger.Error("SaveTakeNumberDataList failed", zap.Error(err))
			return res, err
		}

		res.Msg = fmt.Sprintf("%d 条TakeNumber数据接收成功", len(result))
	case constant.HallCallNumber:
		var result []model.CallNumber

		err := json.Unmarshal(queryModel.PullData, &result)
		if err != nil {
			h.logger.Error("ReceiveThirdPartyPushData unmarshall err:", zap.Error(err))
			return res, err
		}

		err = h.SaveCallNumberDataList(result)
		if err != nil {

			h.logger.Error("SaveCallNumberDataList failed", zap.Error(err))
			return res, err
		}

		res.Msg = fmt.Sprintf("%d 条CallNumber数据接收成功", len(result))
	default:
		h.logger.Error("review no func name")
		res.Msg = "review no func name"
		return res, nil
	}
	return res, nil
}
