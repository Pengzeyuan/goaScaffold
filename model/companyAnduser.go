package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// gorm 关联查询
type User struct {
	//这里grom是映射到表
	ID        int        `gorm:"TYPE:int(11);NOT NULL;PRIMARY_KEY;INDEX"`     // 主键
	Name      string     `gorm:"TYPE: VARCHAR(255); DEFAULT:'';INDEX"`        // 一般建
	Companies []Company  `gorm:"FOREIGNKEY:UserId;ASSOCIATION_FOREIGNKEY:ID"` //  外键 UserId  关联外键—的本表字段  Id
	CreatedAt time.Time  `gorm:"TYPE:DATETIME"`                               // 时间撮
	UpdatedAt time.Time  `gorm:"TYPE:DATETIME"`
	DeletedAt *time.Time `gorm:"TYPE:DATETIME;DEFAULT:NULL"`
}

type Company struct {
	gorm.Model
	Industry int    `gorm:"TYPE:INT(11);DEFAULT:0"`             // 行业
	Name     string `gorm:"TYPE:VARCHAR(255);DEFAULT:'';INDEX"` //名字 一般建
	Job      string `gorm:"TYPE:VARCHAR(255);DEFAULT:''"`
	UserId   int    `gorm:"TYPE:int(11);NOT NULL;INDEX"` // 用户 id 一般建
}
