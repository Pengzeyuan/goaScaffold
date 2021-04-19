package model

type Simulation struct {
	BaseModel
	// key 使用接口名称
	Key string `gorm:"type:varchar(255);not null;unique"`
	// val 使用接口返回值 json
	Val            string `gorm:"type:text;not null"`
	IsShowMock     bool   // 是否显示配置项
	OrderBy        int    // 以哪个字段排序
	OrderTimeScope int    // 排序的时间范围
}
