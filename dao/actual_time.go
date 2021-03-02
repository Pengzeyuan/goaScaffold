package dao

import (
	"boot/model"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// 大厅排队办事实时图
type HallActualTimeDao interface {

	// 存储WindowInfo第三方推送数据
	SaveWindowInfoData(value *model.WindowInfo) error
	// 存储TakeNumber第三方推送数据
	SaveTakeNumberData(value *model.TakeNumber) error
	// 存储CallNumber第三方推送数据
	SaveCallNumberData(value *model.CallNumber) error
}

type HallActualTimeDaoImpl struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewHallActualTimeDaoImpl(db *gorm.DB, logger *zap.Logger) HallActualTimeDao {
	return HallActualTimeDaoImpl{db: db, logger: logger}
}
func (h HallActualTimeDaoImpl) SaveWindowInfoData(value *model.WindowInfo) error {

	if err := h.db.Exec(
		"INSERT INTO `window_info`(`created_at`,`updated_at`,`deleted_at`,`window_status`, `worker`, `window_name`, "+
			"`window_num`, `department`, `window_nature`, `center_name`, `phone_num`, `belong_hall`, `change_mark`, `window_appraise`) "+
			"VALUES (?,?,?,?,?,   ?,?,?,?,?,  ?,?,?,?) "+
			"ON DUPLICATE KEY UPDATE "+
			"updated_at = ?, deleted_at = ?, window_status = ?, worker = ?, window_name = ?, department = ?, window_nature = ?, center_name = ?, "+
			"phone_num = ?, belong_hall = ?, change_mark = ?, window_appraise = ?;",
		value.Model.CreatedAt, value.Model.UpdatedAt, value.Model.DeletedAt, value.WindowStatus, value.Worker, value.WindowName,
		value.WindowNum, value.Department, value.WindowNature, value.CenterName, value.PhoneNum, value.BelongHall,
		value.ChangeMark, value.WindowAppraise,

		value.Model.UpdatedAt, value.Model.DeletedAt, value.WindowStatus, value.Worker, value.WindowName, value.Department, value.WindowNature,
		value.CenterName, value.PhoneNum, value.BelongHall, value.ChangeMark, value.WindowAppraise,
	).Error; err != nil {
		return err
	}
	return nil
}

func (h HallActualTimeDaoImpl) SaveTakeNumberData(value *model.TakeNumber) error {

	if err := h.db.Exec(
		"INSERT INTO `take_number`(`created_at`,`updated_at`,`deleted_at`,`code`, `wait_line_num`, `window_num`, "+
			"`wait_line_type`, `wait_man_type`, `take_man_name`, `take_man_num`, `take_man_society`, `take_man_phone_num`, `area`, "+
			"`do_place`,`department`,`subject_type`,`item_code`,`item_name`,`take_time`) "+
			"VALUES (?,?,?,?,?,   ?,?,?,?,?,  ?,?,?,?,?  ,?,?,?,?) "+
			"ON DUPLICATE KEY UPDATE "+
			"updated_at = ?, deleted_at = ?, wait_line_num = ?, window_num = ?, wait_line_type = ?, wait_man_type = ?, "+
			"take_man_name = ?, take_man_num = ?, take_man_society=?,take_man_phone_num=?, "+
			"area = ?, do_place = ?, department = ?, subject_type = ?,item_code=?,item_name=?,take_time=?;",
		value.Model.CreatedAt, value.Model.UpdatedAt, value.Model.DeletedAt, value.Code, value.WaitLineNum, value.WindowNum, value.WaitLineType,
		value.WaitManType, value.TakeManName, value.TakeManNum, value.TakeManSociety, value.TakeManPhoneNum, value.Area,
		value.DoPlace, value.Department, value.SubjectType, value.ItemCode, value.ItemName, value.TakeTime,

		value.Model.UpdatedAt, value.Model.DeletedAt, value.WaitLineNum, value.WindowNum, value.WaitLineType,
		value.WaitManType, value.TakeManName, value.TakeManNum, value.TakeManSociety, value.TakeManPhoneNum, value.Area,
		value.DoPlace, value.Department, value.SubjectType, value.ItemCode, value.ItemName, value.TakeTime,
	).Error; err != nil {
		return err
	}
	return nil
}

func (h HallActualTimeDaoImpl) SaveCallNumberData(value *model.CallNumber) error {

	if err := h.db.Exec(
		"INSERT INTO `call_number`(`created_at`,`updated_at`,`deleted_at`,`code`, `wait_line_num`, `call_time`, "+
			"`window_num`, `take_man_name`, `item_name`) "+
			"VALUES (?,?,?,?,?,   ?,?,?,?) "+
			"ON DUPLICATE KEY UPDATE "+
			"updated_at = ?, deleted_at = ?, wait_line_num = ?, call_time = ?, window_num = ?, take_man_name = ?, item_name = ?;",
		value.Model.CreatedAt, value.Model.UpdatedAt, value.Model.DeletedAt, value.Code, value.WaitLineNum, value.CallTime,
		value.WindowNum, value.TakeManName, value.ItemName,

		value.Model.UpdatedAt, value.Model.DeletedAt, value.WaitLineNum, value.CallTime, value.WindowNum, value.TakeManName, value.ItemName,
	).Error; err != nil {
		return err
	}
	return nil
}
