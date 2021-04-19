package service

import (
	"boot/dao"
	"boot/model"
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type ItemsDoworkSVC interface {
	// 事项大项拆分数量（今年、去年）、事项拆分比例
	StatReformItemSplit(queryModel model.CommonQueryModel) (model.ItemSplitRate, error)

	// 群众少跑腿
	StatRunOneRate(queryModel model.CommonQueryModel) (model.LimitSceneNumELOneByAllStat, error)
}

type ItemsDoWorkSVCImpl struct {
	ctx            context.Context
	itemsDoworkDao dao.ItemsDoworkDao
	logger         *zap.Logger
}

func NewItemsDoWorkSVCImpl(ctx context.Context, dbItems *gorm.DB, dbDoWork *gorm.DB, logger zap.Logger) ItemsDoworkSVC {
	itemsDoworkDao := dao.NewItemsDoworkDaoImpl(dbItems, dbDoWork, &logger)
	return ItemsDoWorkSVCImpl{
		ctx:            ctx,
		itemsDoworkDao: itemsDoworkDao,
		logger:         &logger,
	}
}

//事项大项拆分数量（今年、去年）、事项拆分比例
func (svc ItemsDoWorkSVCImpl) StatReformItemSplit(queryModel model.CommonQueryModel) (model.ItemSplitRate, error) {
	itemsDoworkDao := svc.itemsDoworkDao

	res := model.ItemSplitRate{}

	timeLayoutStr := "2006-01"
	queryModel.EndDate = time.Now().AddDate(0, 1, 0).Format(timeLayoutStr)
	queryModel.StartDate = "2015-01"

	// 统计今年事项拆分数
	itemSplitCount, errCount := itemsDoworkDao.ItemSplitCount(queryModel)
	if errCount != nil {
		svc.logger.Error("stat item splitcount failed", zap.Error(errCount))
		return res, errCount
	}

	//构建去年的时间范围
	year, _, _ := time.Now().AddDate(-1, 0, 0).Date()
	queryModel.EndDate = fmt.Sprintf("%d-12", year)

	// 统计去年事项拆分数
	pastYearItemSplitCount, errPastYrCont := itemsDoworkDao.ItemSplitCount(queryModel)
	if errPastYrCont != nil {
		svc.logger.Error("stat pastyear item splitcount failed", zap.Error(errCount))
		return res, errPastYrCont
	}
	res.SplitCount = itemSplitCount.SplitCount
	res.PastYearSplit = pastYearItemSplitCount.SplitCount
	// 统计两年间事项拆分比率
	//根据筛选页面所选的时间区间，结束时间的年份作为一个时间点，起始时间的年份作为一个时间点，分别进行筛选（筛选“同步数据Cd_time”时间<=时间点年份的，默认值为当前年份）
	if float32(res.PastYearSplit) > 0 {
		res.SplitRate = (float32(res.SplitCount) - float32(res.PastYearSplit)) / float32(res.PastYearSplit)
	}
	return res, nil
}

//群众少跑腿
func (svc ItemsDoWorkSVCImpl) StatRunOneRate(queryModel model.CommonQueryModel) (model.LimitSceneNumELOneByAllStat, error) {
	itemsDoworkDao := svc.itemsDoworkDao
	res := model.LimitSceneNumELOneByAllStat{}

	if queryModel.EndDate == "" {
		// 获取当前日期
		queryModel.EndDate = time.Now().Format("2006-01-02")
	}

	// 统计一次办的比率 到办事现场次数小于等于1的事项，除以全部事项数目 *100%
	limitSceneOneByAllStat, err := itemsDoworkDao.StatLimitSceneNumELOneByAll(queryModel)
	if err != nil {
		svc.logger.Error("stat  limitscen run one by all failed", zap.Error(err))
		return res, err
	}

	// 根据queryModel.EndDate 获取前一年的日期
	timeEnd, _ := time.Parse("2006-01-02", queryModel.EndDate)
	timeStart := timeEnd.AddDate(-1, 0, 0)
	// 固定结束时间为 yyyy-12-31
	queryModel.EndDate = timeStart.Format("2006") + "-12-31"

	// 统计去年一次办的比率 到办事现场次数小于等于1的事项，除以全部事项数目 *100%
	pastLimitSceneOneByAllStat, err := itemsDoworkDao.StatLimitSceneNumELOneByAll(queryModel)
	if err != nil {
		svc.logger.Error("stat  past year limitscen run one by all failed", zap.Error(err))
		return res, err
	}
	//今年的一次办事项数
	res.RunOneCount = limitSceneOneByAllStat.RunOneCount
	res.TotalItemCount = limitSceneOneByAllStat.TotalItemCount
	//去年的一次办事项数
	res.PastRunOneCount = pastLimitSceneOneByAllStat.RunOneCount
	res.PastTotalItemCount = pastLimitSceneOneByAllStat.TotalItemCount

	//比率提升 = 今年的一次办比率 - 去年的一次办比率
	res.RunProportion = limitSceneOneByAllStat.RunProportion - pastLimitSceneOneByAllStat.RunProportion
	return res, nil
}
