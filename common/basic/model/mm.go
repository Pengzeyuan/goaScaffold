package model

// 测试批量更新
type Demo struct {
	ID    int64   `gorm:"column:id;primary_key"  json:"id"`
	Name  string  `gorm:"column:name"  json:"name"`
	Width float64 `gorm:"column:width"  json:"width"`
}
