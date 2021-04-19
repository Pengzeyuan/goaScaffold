package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type CanalData struct {
	DataType    int         `json:"dataType"`
	InfoDetails interface{} `json:"infoDetails"` // 监听消息详情
}

// 通用查询条件
type CommonQueryModel struct {
	RegionCode     string // 区域代码 如 520100
	StartDate      string // 开始日期 格式 yyyy-MM-dd
	EndDate        string // 结速日期 格式 yyyy-MM-dd
	LimitStartDate string // 限制查询日期 格式 yyyy-MM-dd
	ArgsData       string // 必要的查询参数
	Method         int    // 访问的service方法
	Count          int    // int类型参数
	PullData       []byte // 拉取第三方

}

// 返回参数
type ServiceActualTimeResp struct {
	Count int    // 接收数据量
	Msg   string // 接收成功信息
}
type WindowInfo struct {
	gorm.Model
	WindowStatus   int    `gorm:"type:int;not null;default:0;comment:'窗口状态'"json:"windowStatus"`             // 窗口状态 正常服务:0 暂停服务:1 停用:2
	Worker         string `gorm:"type:varchar(255);not null;default:'';comment:'窗口办事人'"json:"worker"`        // 窗口办事人
	WindowName     string `gorm:"type:varchar(255);not null;default:'';comment:'窗口名称'"json:"windowName"`     // 窗口名称
	WindowNum      string `gorm:"type:varchar(255);not null;default:'';comment:'窗口编号'"json:"windowNum"`      // 窗口编号
	Department     string `gorm:"type:varchar(255);not null;default:'';comment:'所属部门'"json:"department"`     // 所属部门
	WindowNature   string `gorm:"type:varchar(255);not null;default:'';comment:'窗口性质'"json:"windowNature"`   // 窗口性质
	CenterName     string `gorm:"type:varchar(255);not null;default:'';comment:'中心名称'"json:"centerName"`     // 中心名称
	PhoneNum       string `gorm:"type:varchar(255);not null;default:'';comment:'联系电话'"json:"phoneNum"`       // 联系电话
	BelongHall     string `gorm:"type:varchar(255);not null;default:'';comment:'所属大厅'"json:"belongHall"`     // 所属大厅
	ChangeMark     string `gorm:"type:varchar(255);not null;default:'';comment:'变更标记'"json:"changeMark"`     // 变更标记
	WindowAppraise string `gorm:"type:varchar(255);not null;default:'';comment:'窗口评价'"json:"windowAppraise"` // 窗口评价
}

func (WindowInfo) TableName() string {
	return "window_info"
}

// 办事过程表
type DoProcess struct {
	gorm.Model
	// 取号
	Code            string    `gorm:"type:varchar(120);not null;comment:'顾客办事唯一编号code'"json:"code"`                   // 顾客办事唯一编号code
	WaitLineNum     string    `gorm:"type:varchar(255);not null;default:''; comment:'排队号'"json:"waitLineNum"`         // 排队号
	WindowNum       string    `gorm:"type:varchar(255);not null;default:'';comment:'窗口编号'"json:"windowNum"`           // 窗口编号
	WaitLineType    string    `gorm:"type:varchar(255);not null;default:'';comment:'排队类型'"json:"waitLineType"`        // 排队类型
	WaitManType     string    `gorm:"type:varchar(255);not null;default:'';comment:'排队人类型'"json:"waitManType"`        // 排队人类型
	TakeManName     string    `gorm:"type:varchar(255);not null;default:'';comment:'取号人姓名'"json:"takeManName"`        // 取号人姓名
	TakeManNum      string    `gorm:"type:varchar(255);not null;default:'';comment:'取号人身份证'"json:"takeManNum"`        // 取号人身份证
	TakeManSociety  string    `gorm:"type:varchar(255);not null;default:'';comment:'取号人社会信用代码'"json:"takeManSociety"` // 取号人社会信用代码
	TakeManPhoneNum string    `gorm:"type:varchar(255);not null;default:'';comment:'取号人手机号'"json:"takeManPhoneNum"`   // 取号人手机号
	Area            string    `gorm:"type:varchar(255);not null;default:'';comment:'行政区划'"json:"area"`                // 行政区划
	DoPlace         string    `gorm:"type:varchar(255);not null;default:'';comment:'办事地点'"json:"doPlace"`             // 办事地点
	Department      string    `gorm:"type:varchar(255);not null;default:'';comment:'部门信息'"json:"department"`          // 办事项所属部门信息
	SubjectType     string    `gorm:"type:varchar(255);not null;default:'';comment:'主题类型'"json:"subjectType"`         // 主题类型
	ItemCode        string    `gorm:"type:varchar(255);not null;default:'';comment:'事项编码'"json:"itemCode"`            // 事项编码
	ItemName        string    `gorm:"type:varchar(255);not null;default:'';comment:'事项名称'"json:"itemName"`            // 事项名称
	TakeTime        time.Time `json:"takeTime"`                                                                       // 取号时间:叫号时间:办结时间

	// 办结
	HandleResult string `gorm:"type:varchar(255);not null;default:'';comment:'事项名称'"json:"handleResult"` // 办理结果                                                                // 呼叫时间
	// 窗口评价
	Worker           string `gorm:"type:varchar(255);not null;default:'';comment:'窗口办事人'"json:"worker"`          // 窗口办事人
	WindowName       string `gorm:"type:varchar(255);not null;default:'';comment:'窗口名称'"json:"windowName"`       // 窗口名称
	WindowDepartment string `gorm:"type:varchar(255);not null;default:'';comment:'所属部门'"json:"windowDepartment"` // 窗口所属部门
	WindowNature     string `gorm:"type:varchar(255);not null;default:'';comment:'窗口性质'"json:"windowNature"`     // 窗口性质
	CenterName       string `gorm:"type:varchar(255);not null;default:'';comment:'中心名称'"json:"centerName"`       // 中心名称
	PhoneNum         string `gorm:"type:varchar(255);not null;default:'';comment:'联系电话'"json:"phoneNum"`         // 联系电话
	BelongHall       string `gorm:"type:varchar(255);not null;default:'';comment:'所属大厅'"json:"belongHall"`       // 所属大厅
	ChangeMark       string `gorm:"type:varchar(255);not null;default:'';comment:'变更标记'"json:"changeMark"`       // 变更标记
	WindowAppraise   string `gorm:"type:varchar(255);not null;default:'';comment:'窗口评价'"json:"windowAppraise"`   // 窗口评价
	// 状态
	DoProcessState string `gorm:"type:varchar(255);not null;default:'';comment:'办事状态'"json:"doProcessState"` // 办事状态
	WindowId       int    `gorm:"type:int;not null;default:0;comment:'窗口id'"json:"windowId"`                 // 窗口id
}

func (DoProcess) TableName() string {
	return "do_process"
}

func (c *DoProcess) UnmarshalJSON(data []byte) (err error) {
	type Temp DoProcess
	t := struct {
		TakeTime DateTime `json:"takeTime"`
		Temp
	}{
		Temp: (Temp)(*c),
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	// 取号
	c.Code = t.Code
	c.WaitLineNum = t.WaitLineNum
	c.WindowNum = t.WindowNum
	c.WaitLineType = t.WaitLineType
	c.WaitManType = t.WaitManType
	c.TakeManName = t.TakeManName
	c.TakeManNum = t.TakeManNum
	c.TakeManSociety = t.TakeManSociety
	c.TakeManPhoneNum = t.TakeManPhoneNum
	c.Area = t.Area
	c.DoPlace = t.DoPlace
	c.Department = t.Department
	c.SubjectType = t.SubjectType
	c.ItemCode = t.ItemCode
	c.ItemName = t.ItemName
	unix := t.TakeTime.Unix()
	c.TakeTime = time.Unix(unix, 0)

	// 办结
	c.HandleResult = t.HandleResult

	// 评价
	c.Worker = t.Worker
	c.WindowName = t.WindowName
	c.WindowDepartment = t.WindowDepartment
	c.WindowNature = t.WindowNature
	c.CenterName = t.CenterName
	c.PhoneNum = t.PhoneNum
	c.BelongHall = t.BelongHall
	c.ChangeMark = t.ChangeMark
	c.WindowAppraise = t.WindowAppraise
	// 新增
	c.DoProcessState = t.DoProcessState
	c.WindowId = t.WindowId

	return nil
}

type DateTime struct {
	time.Time
}

func (t *DateTime) UnmarshalJSON(data []byte) (err error) {
	data = bytes.Trim(data, "\"") //此除需要去掉传入的数据的两端的 ""
	ext, err := time.ParseInLocation("2006-01-02 15:04:05", string(data), time.Local)
	if err != nil {
		zap.L().Error("string DateTime ParseInLocation fail", zap.Error(err))
		return nil
	}
	*t = DateTime{ext}
	return nil
}

//大厅管理系统设立情况  本地数据库对象
type HallManagementInfo struct {
	gorm.Model
	CardNum string `gorm:"unique_index:S_R;type:varchar(50); comment:'id卡号'"`
	Name    string `gorm:"type:varchar(50); comment:'名字'"`
	OuName  string `gorm:"type:varchar(50); comment:'部门名字'"`
}

//type HallManagementSystem struct {
//	"cardnum": "522125199301100062",
//	"img": "",
//	"loginname": "18096101310",
//	"sex": "女",
//	"name": "李燕",
//	"oucode": "520100SJ",
//	"ouname": "贵阳市市场监督管理局",
//	"gonghao": "0425"
//
//
//	WindowStatus   int    `gorm:"type:int;not null;default:0;comment:'窗口状态'"json:"windowStatus"`               // 窗口状态 正常服务:0 暂停服务:1 停用:2
//	cardnum         string `gorm:"type:varchar(255);not null;default:'';comment:'省份证号'"json:"cardnum"`          // 窗口办事人
//	WindowName     string `gorm:"type:varchar(255);not null;default:'';comment:'窗口名称'"json:"windowName"`       // 窗口名称
//	WindowNum      int    `gorm:"unique_index:S_R;type:int;not null;default:0;comment:'窗口编号'"json:"windowNum"` // 窗口编号
//	Department     string `gorm:"type:varchar(255);not null;default:'';comment:'所属部门'"json:"department"`       // 所属部门
//	WindowNature   string `gorm:"type:varchar(255);not null;default:'';comment:'窗口性质'"json:"windowNature"`     // 窗口性质
//	CenterName     string `gorm:"type:varchar(255);not null;default:'';comment:'中心名称'"json:"centerName"`       // 中心名称
//	PhoneNum       string `gorm:"type:varchar(255);not null;default:'';comment:'联系电话'"json:"phoneNum"`         // 联系电话
//	BelongHall     string `gorm:"type:varchar(255);not null;default:'';comment:'所属大厅'"json:"belongHall"`       // 所属大厅
//	ChangeMark     string `gorm:"type:varchar(255);not null;default:'';comment:'变更标记'"json:"changeMark"`       // 变更标记
//	WindowAppraise string `gorm:"type:varchar(255);not null;default:'';comment:'窗口评价'"json:"windowAppraise"`   // 窗口评价
//}

type DateTime2 struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05"
const ctLayout_nosec = "2006-01-02 15:04"
const ctLayout_date = "2006-01-02"

func (this *DateTime2) UnmarshalJSON(b []byte) (err error) {

	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	sv := string(b)
	if len(sv) == 10 {
		sv += " 00:00:00"
	} else if len(sv) == 16 {
		sv += ":00"
	}
	this.Time, err = time.ParseInLocation(ctLayout, string(b), loc)
	if err != nil {
		if this.Time, err = time.ParseInLocation(ctLayout_nosec, string(b), loc); err != nil {
			this.Time, err = time.ParseInLocation(ctLayout_date, string(b), loc)
		}
	}

	return
}

func (this *DateTime2) MarshalJSON() ([]byte, error) {

	rs := []byte(fmt.Sprintf(`"%s"`, this.Time.Format(ctLayout)))

	return rs, nil
}

var nilTime = (time.Time{}).UnixNano()

func (this *DateTime2) IsSet() bool {
	return this.UnixNano() != nilTime
}
