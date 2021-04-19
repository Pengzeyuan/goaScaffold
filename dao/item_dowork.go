package dao

import (
	"boot/model"
	"fmt"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// 办件信息库
type ItemsDoworkDao interface {
	// 统计事项拆分数 事项大项拆分数量（今年、去年）、事项拆分比例
	ItemSplitCount(queryModel model.CommonQueryModel) (model.ReformResult, error)

	// 到办事现场次数小于等于1的事项，除以全部事项数目 *100%
	StatLimitSceneNumELOneByAll(queryModel model.CommonQueryModel) (model.LimitSceneNumELOneByAllStat, error)
}

type ItemsDoworkDaoImpl struct {
	itemsDb  *gorm.DB
	doWorkDb *gorm.DB
	logger   *zap.Logger
}

func NewItemsDoworkDaoImpl(db1 *gorm.DB, db2 *gorm.DB, logger *zap.Logger) ItemsDoworkDao {
	return ItemsDoworkDaoImpl{itemsDb: db1, doWorkDb: db2, logger: logger}
}

// 事项拆分数
func (t ItemsDoworkDaoImpl) ItemSplitCount(queryModel model.CommonQueryModel) (model.ReformResult, error) {
	var res model.ReformResult
	regionCodeWhere := ""
	if queryModel.RegionCode != "" {
		regionCodeWhere = "and REGION_CODE = ? "
	}
	querySQL := `
		SELECT ItemCount-TotalCount AS split_count
		FROM
		(
			SELECT count(ID) ItemCount
			FROM(
					SELECT TreeCode,ItemCode,max(ID) ID 
					FROM up_task_general_basic 
					WHERE IsItem=1
						AND taskstate=1
						AND DATE_FORMAT(Cd_time,'%%Y-%%m') BETWEEN ? AND ?
						%s
					GROUP BY TreeCode,ItemCode
			) inner_tb
		) child_item_tb,
		(
	    	SELECT count(ID) TotalCount
			FROM(
					SELECT TreeCode,max(ID) ID
					FROM up_task_general_basic
					WHERE IsItem=1
						AND taskstate=1
						AND DATE_FORMAT(Cd_time,'%%Y-%%m') BETWEEN ? AND ?
						%s
					GROUP BY TreeCode
			) inner_tb
		) item_tb
	`
	//加入区域编码条件
	querySQL = fmt.Sprintf(querySQL, regionCodeWhere, regionCodeWhere)

	var err error
	if queryModel.RegionCode != "" {
		err = t.itemsDb.Raw(querySQL, queryModel.StartDate, queryModel.EndDate, queryModel.RegionCode,
			queryModel.StartDate, queryModel.EndDate, queryModel.RegionCode).Scan(&res).Error
	} else {
		err = t.itemsDb.Raw(querySQL, queryModel.StartDate, queryModel.EndDate,
			queryModel.StartDate, queryModel.EndDate).Scan(&res).Error
	}

	return res, err
}

// 到办事现场次数小于等于1的事项，除以全部事项数目 *100%
func (t ItemsDoworkDaoImpl) StatLimitSceneNumELOneByAll(queryModel model.CommonQueryModel) (model.LimitSceneNumELOneByAllStat, error) {
	var res model.LimitSceneNumELOneByAllStat
	sql := `
		SELECT COUNT(up_task_general_basic.ID)/COUNT(join_tb.ID) AS run_proportion,COUNT(up_task_general_basic.ID) AS run_one_count,COUNT(join_tb.ID) AS total_item_count
		FROM up_task_general_basic
		RIGHT JOIN
		(
				SELECT ItemCode,max(ID) ID 
				FROM up_task_general_basic 
				WHERE IsItem=1
					AND DATE_FORMAT(Cd_time,'%%Y-%%m-%%d') BETWEEN '2000-01-01' AND ?
						%s
				GROUP BY ItemCode
		) join_tb
		ON up_task_general_basic.ID = join_tb.ID
			AND LimitSceneNum<=1
	`
	regionCodeWhere := "AND REGION_CODE like  '5201%%'"
	if queryModel.RegionCode != "" {
		regionCodeWhere = "AND REGION_CODE = ?"
	}
	var err error
	querySQL := fmt.Sprintf(sql, regionCodeWhere)
	if queryModel.RegionCode != "" {
		err = t.itemsDb.Raw(querySQL, queryModel.EndDate, queryModel.RegionCode).Scan(&res).Error
	} else {
		err = t.itemsDb.Raw(querySQL, queryModel.EndDate).Scan(&res).Error
	}
	return res, err

}

// 让某个字段值加一
func (t ItemsDoworkDaoImpl) AdUpdateSignUpNum(id int64) bool {
	//var res model.LimitSceneNumELOneByAllStat

	//sql := `
	//	SELECT COUNT(up_task_general_basic.ID)/COUNT(join_tb.ID) AS run_proportion,COUNT(up_task_general_basic.ID) AS run_one_count,COUNT(join_tb.ID) AS total_item_count
	//	FROM up_task_general_basic
	//	RIGHT JOIN
	//	(
	//			SELECT ItemCode,max(ID) ID
	//			FROM up_task_general_basic
	//			WHERE IsItem=1
	//				AND DATE_FORMAT(Cd_time,'%%Y-%%m-%%d') BETWEEN '2000-01-01' AND ?
	//					%s
	//			GROUP BY ItemCode
	//	) join_tb
	//	ON up_task_general_basic.ID = join_tb.ID
	//		AND LimitSceneNum<=1
	//`
	//regionCodeWhere := "AND REGION_CODE like  '5201%%'"
	//if queryModel.RegionCode != "" {
	//	regionCodeWhere = "AND REGION_CODE = ?"
	//}
	//var err error
	//querySQL := fmt.Sprintf(sql, regionCodeWhere)
	//if queryModel.RegionCode != "" {
	//	err = t.itemsDb.Raw(querySQL, queryModel.EndDate, queryModel.RegionCode).Scan(&res).Error
	//} else {
	//	err = t.itemsDb.Raw(querySQL, queryModel.EndDate).Scan(&res).Error
	//}
	//return res, err
	//
	//if xy.ID <= 0 {
	//	return false
	//}
	//if err := Db.Model(xy).Where("id = ? ", id).Update("sign_up_num", gorm.Expr("sign_up_num+ ?", 1)).Error; err != nil {
	//	return false
	//}
	return true
}
