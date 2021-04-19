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
)

// 大厅排队办事实时图service
type HallActualTimeSVC interface {

	// 接收第三方推送数据--大厅排队办事实时图基础数据
	ReceiveThirdPartyPushData(queryModel model.CommonQueryModel) (model.ServiceActualTimeResp, error)

	// 存储WindowInfo第三方推送数据列表
	SaveWindowInfoDataList(valueList []model.WindowInfo) (int64, error)
	// 存储TakeNumber第三方推送数据列表
	SaveTakeNumberDataList(valueList []model.DoProcess) (int64, error)
	// 存储CallNumber第三方推送数据列表
	SaveCallNumberDataList(valueList []model.DoProcess) (int64, error)

	//存储TransactionCompleted第三方推送数据列表
	SaveTransactionCompletedDataList(valueList []model.DoProcess) (int64, error)

	//存储EvaluateTable第三方推送数据列表
	SaveEvaluateDataList(valueList []model.DoProcess) (int64, error)
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
func (h HallActualTimeSVCImpl) SaveWindowInfoDataList(valueList []model.WindowInfo) (int64, error) {
	actualTimeDao := h.hallActualTimeDao

	affectRow, err := actualTimeDao.SaveWindowInfoData(valueList)
	if err != nil {
		h.logger.Error("SaveWindowInfoDataList is failed:", zap.Error(err))
		return affectRow, err
	}

	return affectRow, nil
}

// 存储TakeNumber第三方推送数据列表
func (h HallActualTimeSVCImpl) SaveTakeNumberDataList(valueList []model.DoProcess) (int64, error) {
	actualTimeDao := h.hallActualTimeDao

	affectRow, err := actualTimeDao.SaveTakeNumberData(valueList)
	if err != nil {
		h.logger.Error("SaveTakeNumberDataList is failed:", zap.Error(err))
		return affectRow, err
	}

	return affectRow, err
}

// 存储CallNumber第三方推送数据列表
func (h HallActualTimeSVCImpl) SaveCallNumberDataList(valueList []model.DoProcess) (int64, error) {
	actualTimeDao := h.hallActualTimeDao

	affectRow, err := actualTimeDao.SaveCallNumberData(valueList)
	if err != nil {
		h.logger.Error("SaveCallNumberDataList is failed:", zap.Error(err))
		return affectRow, err
	}

	return affectRow, err
}

// 存储TransactionCompleted第三方推送数据列表
func (h HallActualTimeSVCImpl) SaveTransactionCompletedDataList(valueList []model.DoProcess) (int64, error) {
	actualTimeDao := h.hallActualTimeDao

	affectRow, err := actualTimeDao.SaveTransactionCompletedData(valueList)
	if err != nil {
		h.logger.Error("SaveTransactionCompletedDataList is failed:", zap.Error(err))
		return affectRow, err
	}

	return affectRow, err
}

// 存储Evaluate第三方推送数据列表
func (h HallActualTimeSVCImpl) SaveEvaluateDataList(valueList []model.DoProcess) (int64, error) {
	actualTimeDao := h.hallActualTimeDao

	affectRow, err := actualTimeDao.SaveEvaluateData(valueList)
	if err != nil {
		h.logger.Error("SaveEvaluateDataList is failed:", zap.Error(err))
		return affectRow, err
	}
	return affectRow, err
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

		affectRow, err := h.SaveWindowInfoDataList(result)
		if err != nil {
			h.logger.Error("SaveWindowInfoDataList failed", zap.Error(err))
			return res, err
		}

		res.Msg = fmt.Sprintf("接收%d 条WindowInfo数据,成功插入%d 条", len(result), affectRow)
	case constant.HallTakeNumber:

		var result []model.DoProcess

		err := json.Unmarshal(queryModel.PullData, &result)
		if err != nil {
			h.logger.Error("ReceiveThirdPartyPushData unmarshall err:", zap.Error(err))
			return res, err
		}
		affectRows, err := h.SaveTakeNumberDataList(result)
		if err != nil {

			h.logger.Error("SaveTakeNumberDataList failed", zap.Error(err))
			return res, err
		}

		res.Msg = fmt.Sprintf("接收%d 条TakeNumber数据,成功插入%d 条", len(result), affectRows)
	case constant.HallCallNumber:

		var result []model.DoProcess

		err := json.Unmarshal(queryModel.PullData, &result)
		if err != nil {
			h.logger.Error("ReceiveThirdPartyPushData unmarshall err:", zap.Error(err))
			return res, err
		}

		affectRows, err := h.SaveCallNumberDataList(result)
		if err != nil {

			h.logger.Error("SaveCallNumberDataList failed", zap.Error(err))
			return res, err
		}

		res.Msg = fmt.Sprintf("接收%d 条CallNumber数据,成功插入%d 条", len(result), affectRows)
	case constant.HallTransactionCompleted:
		var result []model.DoProcess

		err := json.Unmarshal(queryModel.PullData, &result)
		if err != nil {
			h.logger.Error("ReceiveThirdPartyPushData unmarshall err:", zap.Error(err))
			return res, err
		}

		affectRows, err := h.SaveTransactionCompletedDataList(result)
		if err != nil {

			h.logger.Error("SaveTransactionCompletedDataList failed", zap.Error(err))
			return res, err
		}

		res.Msg = fmt.Sprintf("接收%d 条TransactionCompleted数据,成功插入%d 条", len(result), affectRows)
	case constant.HallEvaluate:
		var result []model.DoProcess

		err := json.Unmarshal(queryModel.PullData, &result)
		if err != nil {
			h.logger.Error("ReceiveThirdPartyPushData unmarshall err:", zap.Error(err))
			return res, err
		}

		affectRows, err := h.SaveEvaluateDataList(result)
		if err != nil {

			h.logger.Error("SaveEvaluateDataList failed", zap.Error(err))
			return res, err
		}

		res.Msg = fmt.Sprintf("接收%d 条Evaluate数据,成功插入%d 条", len(result), affectRows)
	default:
		h.logger.Error("review no func name")
		res.Msg = "review no func name"
		return res, nil
	}
	return res, nil
}
