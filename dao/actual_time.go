package dao

import (
	"boot/model"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// 大厅排队办事实时图
type HallActualTimeDao interface {

	// 存储WindowInfo第三方推送数据
	SaveWindowInfoData(values []model.WindowInfo) (int64, error)
	// 存储TakeNumber第三方推送数据
	SaveTakeNumberData(valueList []model.DoProcess) (int64, error)
	// 存储CallNumber第三方推送数据
	SaveCallNumberData(valueList []model.DoProcess) (int64, error)

	// 存储TransactionCompleted第三方推送数据
	SaveTransactionCompletedData(valueList []model.DoProcess) (int64, error)
	// 存储EvaluateData第三方推送数据
	SaveEvaluateData(valueList []model.DoProcess) (int64, error)
}

type HallActualTimeDaoImpl struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewHallActualTimeDaoImpl(db *gorm.DB, logger *zap.Logger) HallActualTimeDao {
	return HallActualTimeDaoImpl{db: db, logger: logger}
}
func (h HallActualTimeDaoImpl) SaveWindowInfoData(values []model.WindowInfo) (int64, error) {

	nowTime := time.Now()
	sqlStr := "INSERT IGNORE INTO window_info( created_at,updated_at,  window_status, worker, window_name, " +
		"window_num, department, window_nature, center_name, phone_num, belong_hall, change_mark, window_appraise) VALUES "
	vals := []interface{}{}
	const rowSQL = "(?,?,?,?,?,  ?,?,?,?,?  ,?,?,?)"
	var inserts []string
	for _, elem := range values {
		elem.CreatedAt, elem.UpdatedAt = nowTime, nowTime
		inserts = append(inserts, rowSQL)
		vals = append(vals, elem.CreatedAt, elem.UpdatedAt, elem.WindowStatus, elem.Worker, elem.WindowName,
			elem.WindowNum, elem.Department, elem.WindowNature, elem.CenterName, elem.PhoneNum, elem.BelongHall, elem.ChangeMark,
			elem.WindowAppraise)
	}
	sqlStr = sqlStr + strings.Join(inserts, ",")

	db := h.db.Exec(sqlStr, vals...)
	if err := db.Error; err != nil {
		return db.RowsAffected, err
	}
	return db.RowsAffected, nil
}

func (h HallActualTimeDaoImpl) SaveTakeNumberData(valueList []model.DoProcess) (int64, error) {
	nowTime := time.Now()
	sqlStr := "INSERT IGNORE INTO do_process( created_at,updated_at,  code, wait_line_num, window_num, " +
		"wait_line_type, wait_man_type, take_man_name, take_man_num, take_man_society, take_man_phone_num, area," +
		" do_place,department,subject_type,item_code,item_name,take_time,do_process_state) VALUES "
	vals := []interface{}{}
	const rowSQL = "(?,?,?,?,?,   ?,?,?,?,?,  ?,?,?,?,?  ,?,?,?,?)"
	var inserts []string
	for _, elem := range valueList {
		elem.CreatedAt, elem.UpdatedAt = nowTime, nowTime
		inserts = append(inserts, rowSQL)
		vals = append(vals, elem.CreatedAt, elem.UpdatedAt, elem.Code, elem.WaitLineNum, elem.WindowNum,
			elem.WaitLineType, elem.WaitManType, elem.TakeManName, elem.TakeManNum, elem.TakeManSociety, elem.TakeManPhoneNum,
			elem.Area, elem.DoPlace, elem.Department, elem.SubjectType, elem.ItemCode, elem.ItemName, elem.TakeTime, "取号")
	}
	sqlStr = sqlStr + strings.Join(inserts, ",")

	db := h.db.Exec(sqlStr, vals...)
	if err := db.Error; err != nil {
		return db.RowsAffected, err
	}
	return db.RowsAffected, nil
}

// 临时表的方案
func (h HallActualTimeDaoImpl) SaveCallNumberData(valueList []model.DoProcess) (int64, error) {
	nowTime := time.Now()
	sqlStr := "INSERT IGNORE INTO do_process( created_at,updated_at,  code, take_time,do_process_state) VALUES "
	vals := []interface{}{}
	const rowSQL = "(?,?,?,?,?)"
	var inserts []string
	for _, elem := range valueList {
		elem.CreatedAt, elem.UpdatedAt = nowTime, nowTime
		inserts = append(inserts, rowSQL)
		vals = append(vals, elem.CreatedAt, elem.UpdatedAt, elem.Code, elem.TakeTime, "叫号")
	}
	sqlStr = sqlStr + strings.Join(inserts, ",")

	db := h.db.Exec(sqlStr, vals...)
	if err := db.Error; err != nil {
		return db.RowsAffected, err
	}
	return db.RowsAffected, nil
}

// 存储TransactionCompleted第三方推送数据
func (h HallActualTimeDaoImpl) SaveTransactionCompletedData(valueList []model.DoProcess) (int64, error) {
	nowTime := time.Now()
	sqlStr := "INSERT IGNORE INTO do_process( created_at,updated_at,  code, handle_result, take_time, do_process_state) VALUES "
	vals := []interface{}{}
	const rowSQL = "(?,?,?,?,?,   ?)"
	var inserts []string
	for _, elem := range valueList {
		elem.CreatedAt, elem.UpdatedAt = nowTime, nowTime
		inserts = append(inserts, rowSQL)
		vals = append(vals, elem.CreatedAt, elem.UpdatedAt, elem.Code, elem.HandleResult, elem.TakeTime, "办结")
	}
	sqlStr = sqlStr + strings.Join(inserts, ",")

	db := h.db.Exec(sqlStr, vals...)
	if err := db.Error; err != nil {
		return db.RowsAffected, err
	}
	return db.RowsAffected, nil
}

// 存储EvaluateData第三方推送数据
func (h HallActualTimeDaoImpl) SaveEvaluateData(valueList []model.DoProcess) (int64, error) {
	nowTime := time.Now()
	sqlStr := "INSERT  IGNORE INTO do_process( created_at,updated_at,  code, worker, window_name, " +
		"window_department, window_nature, center_name, phone_num, belong_hall, change_mark, window_appraise,do_process_state) VALUES "
	vals := []interface{}{}
	const rowSQL = "(?,?,?,?,?,   ?,?,?,?,?,  ?,?,?)"
	var inserts []string
	for _, elem := range valueList {
		elem.CreatedAt, elem.UpdatedAt = nowTime, nowTime
		inserts = append(inserts, rowSQL)
		vals = append(vals, elem.CreatedAt, elem.UpdatedAt, elem.Code, elem.Worker, elem.WindowName,
			elem.WindowDepartment, elem.WindowNature, elem.CenterName, elem.PhoneNum, elem.BelongHall, elem.ChangeMark,
			elem.WindowAppraise, "评价")
	}
	sqlStr = sqlStr + strings.Join(inserts, ",")

	db := h.db.Exec(sqlStr, vals...)
	if err := db.Error; err != nil {
		return db.RowsAffected, err
	}
	return db.RowsAffected, nil
}
