package model

import "github.com/jinzhu/gorm"

// "马上办"

type ImmediateInfo struct {
	gorm.Model
	DeptName       string `gorm:"type:text"` // 部门名称
	ParItemName    string `gorm:"type:text"` // 事项大项名称
	ChiItemName    string `gorm:"type:text"` // 事项小项名称
	PresentTimes   int    `gorm:"type:int"`  // 到现场次数
	ManagedObj     string `gorm:"type:text"` // 办理对象
	ItemType       string `gorm:"type:text"` // 事项类型
	Implementation string `gorm:"type:text"` // 实现方式
	Year           int    `gorm:"type:int"`  // 年份
	Area           string `gorm:"type:text"` // 区域
}

func (ImmediateInfo) TableName() string {
	return "immediate_info"
}

// "网上办"
type OnlineInfo struct {
	gorm.Model
	DeptName       string `gorm:"type:text"` // 部门名称
	ParItemName    string `gorm:"type:text"` // 事项大项名称
	ChiItemName    string `gorm:"type:text"` // 事项小项名称
	PresentTimes   int    `gorm:"type:int"`  // 到现场次数
	ManagedObj     string `gorm:"type:text"` // 办理对象
	ItemType       string `gorm:"type:text"` // 事项类型
	Implementation string `gorm:"type:text"` // 实现方式
	Year           int    `gorm:"type:int"`  // 年份
	Area           string `gorm:"type:text"` // 区域
}

func (OnlineInfo) TableName() string {
	return "online_info"
}

// "就近办"
type NearbyInfo struct {
	gorm.Model
	OrgName        string `gorm:"type:text"` // 单位名称
	ParItemName    string `gorm:"type:text"` // 事项大项名称
	ChiItemName    string `gorm:"type:text"` // 事项小项名称
	AdmissibleArea string `gorm:"type:text"` // 可受理行政区域
	Implementation string `gorm:"type:text"` // 实现方式
	Year           int    `gorm:"type:int"`  // 年份
	Area           string `gorm:"type:text"` // 区域
}

func (NearbyInfo) TableName() string {
	return "nearby_info"
}

// "一次办"
type OnceInfo struct {
	gorm.Model
	DeptName       string `gorm:"type:text"` // 部门名称
	ParItemName    string `gorm:"type:text"` // 事项大项名称
	ChiItemName    string `gorm:"type:text"` // 事项小项名称
	PresentTimes   int    `gorm:"type:int"`  // 到现场次数
	ManagedObj     string `gorm:"type:text"` // 办理对象
	ItemType       string `gorm:"type:text"` // 事项类型
	Implementation string `gorm:"type:text"` // 实现方式
	Year           int    `gorm:"type:int"`  // 年份
	Area           string `gorm:"type:text"` // 区域
}

func (OnceInfo) TableName() string {
	return "once_info"
}

type FourDoCount struct { // 用于查询返回
	ImmediateInfoCount []*FileCount // "马上办统计数量"
	OnlineInfoCount    []*FileCount // "网上办统计数量"
	NearbyInfoCount    []*FileCount // "就近办统计数量"
	OnceInfoCount      []*FileCount // "一次办统计数量"
}

type FileCount struct { // 用于查询返回
	Year  int // 年份
	Count int // 数量
}

// 通用查询条件
type RedisQueryModel struct {
	ModelName  string // 区域代码 如 520100
	MethodName string // 开始日期 格式 yyyy-MM-dd
}

//事项拆分比例
type ItemSplitRate struct {
	SplitCount    int32   // 今年事项拆分数
	PastYearSplit int32   // 去年事项拆分数
	SplitRate     float32 //事项拆分比率
}

type ReformResult struct { // 改革成效
	SplitCount          int32   // 事项拆分数
	HandleTime          float32 // 办理时长
	HandleTimeCompress  float32 // 时间压缩
	PraiseProportion    float32 // 好评率
	OneWindowProportion float32 // 一窗业务占比
	RunProportion       float32 // 最多跑一次占比
}
