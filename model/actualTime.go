package model

import (
	"bytes"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"time"
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
	WindowStatus   int    `gorm:"type:int;not null;default:0;comment:'窗口状态'"json:"windowStatus"`               // 窗口状态 正常服务:0 暂停服务:1 停用:2
	Worker         string `gorm:"type:varchar(255);not null;default:'';comment:'窗口办事人'"json:"worker"`          // 窗口办事人
	WindowName     string `gorm:"type:varchar(255);not null;default:'';comment:'窗口名称'"json:"windowName"`       // 窗口名称
	WindowNum      int    `gorm:"unique_index:S_R;type:int;not null;default:0;comment:'窗口编号'"json:"windowNum"` // 窗口编号
	Department     string `gorm:"type:varchar(255);not null;default:'';comment:'所属部门'"json:"department"`       // 所属部门
	WindowNature   string `gorm:"type:varchar(255);not null;default:'';comment:'窗口性质'"json:"windowNature"`     // 窗口性质
	CenterName     string `gorm:"type:varchar(255);not null;default:'';comment:'中心名称'"json:"centerName"`       // 中心名称
	PhoneNum       string `gorm:"type:varchar(255);not null;default:'';comment:'联系电话'"json:"phoneNum"`         // 联系电话
	BelongHall     string `gorm:"type:varchar(255);not null;default:'';comment:'所属大厅'"json:"belongHall"`       // 所属大厅
	ChangeMark     string `gorm:"type:varchar(255);not null;default:'';comment:'变更标记'"json:"changeMark"`       // 变更标记
	WindowAppraise string `gorm:"type:varchar(255);not null;default:'';comment:'窗口评价'"json:"windowAppraise"`   // 窗口评价
}

func (WindowInfo) TableName() string {
	return "window_info"
}

type TakeNumber struct {
	gorm.Model
	Code            string    `gorm:"unique_index:S_R;type:varchar(120);not null;unique;comment:'唯一编号code'"  json:"code"` // 唯一编号code
	WaitLineNum     int       `gorm:"type:int;not null;default:0;comment:'排队号'"json:"waitLineNum"`                        // 排队号
	WindowNum       int       `gorm:"type:int;not null;default:0;comment:'窗口号'"json:"windowNum"`                          // 窗口号
	WaitLineType    string    `gorm:"type:varchar(255);not null;default:'';comment:'排队类型'"json:"waitLineType"`            // 排队类型
	WaitManType     string    `gorm:"type:varchar(255);not null;default:'';comment:'排队人类型'"json:"waitManType"`            // 排队人类型
	TakeManName     string    `gorm:"type:varchar(255);not null;default:'';comment:'取号人姓名'"json:"takeManName"`            // 取号人姓名
	TakeManNum      string    `gorm:"type:varchar(255);not null;default:'';comment:'取号人身份证'"json:"takeManNum"`            // 取号人身份证
	TakeManSociety  string    `gorm:"type:varchar(255);not null;default:'';comment:'取号人社会信用代码'"json:"takeManSociety"`     // 取号人社会信用代码
	TakeManPhoneNum string    `gorm:"type:varchar(255);not null;default:'';comment:'取号人手机号'"json:"takeManPhoneNum"`       // 取号人手机号
	Area            string    `gorm:"type:varchar(255);not null;default:'';comment:'行政区划'"json:"area"`                    // 行政区划
	DoPlace         string    `gorm:"type:varchar(255);not null;default:'';comment:'办事地点'"json:"doPlace"`                 // 办事地点
	Department      string    `gorm:"type:varchar(255);not null;default:'';comment:'部门信息'"json:"department"`              // 部门信息
	SubjectType     string    `gorm:"type:varchar(255);not null;default:'';comment:'主题类型'"json:"subjectType"`             // 主题类型
	ItemCode        string    `gorm:"type:varchar(255);not null;default:'';comment:'事项编码'"json:"itemCode"`                // 事项编码
	ItemName        string    `gorm:"type:varchar(255);not null;default:'';comment:'事项名称'"json:"itemName"`                // 事项名称
	TakeTime        time.Time `json:"takeTime"`                                                                           // 取号时间
}

func (TakeNumber) TableName() string {
	return "take_number"
}

type CallNumber struct {
	gorm.Model
	Code        string    `gorm:"unique_index:S_R;type:varchar(120);not null;unique;comment:'唯一编号code'"json:"code"` // 唯一编号code
	WaitLineNum int       `gorm:"type:int;not null;default:0;comment:'排队号'"json:"waitLineNum"`                      // 排队号
	CallTime    time.Time `json:"callTime"`                                                                         // 呼叫时间
	WindowNum   int       `gorm:"type:int;not null;default:0;comment:'窗口号'"json:"windowNum"`                        // 窗口号
	TakeManName string    `gorm:"type:varchar(255);not null;default:'';comment:'取号人姓名'"json:"takeManName"`          // 取号人姓名
	ItemName    string    `gorm:"type:varchar(255);not null;default:'';comment:'事项名称'"json:"itemName"`              // 事项名称
}

func (CallNumber) TableName() string {
	return "call_number"
}
func (c *CallNumber) UnmarshalJSON(data []byte) (err error) {

	type Temp CallNumber
	t := struct {
		CallTime DateTime `json:"callTime"`
		Temp
	}{
		Temp: (Temp)(*c),
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	c.Code = t.Code
	c.WaitLineNum = t.WaitLineNum

	c.WindowNum = t.WindowNum
	c.TakeManName = t.TakeManName
	c.ItemName = t.ItemName

	unix := t.CallTime.Unix()
	c.CallTime = time.Unix(unix, 0)
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
