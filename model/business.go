package model

import (
	"gorm.io/gorm"
)

type Dog struct {
	gorm.Model
	Id          int    `gorm:"type:int;not null;DEFAULT:0;primary_key; comment:'id号'"json:"id"`
	DogName     string `gorm:"TYPE:VARCHAR(255);DEFAULT:'';comment:'狗名称'"json:"yiemianDogName"`
	DogPrice    string `gorm:"TYPE:VARCHAR(255);DEFAULT:'0';comment:'狗价格'"json:"yiemianDogPrice"`
	DogType     int    `gorm:"TYPE:INT(11);DEFAULT:0 ;comment:'狗价格'"`
	Description string `gorm:"TYPE:VARCHAR(255);DEFAULT:''; comment:'描述'"`
}
type Cat struct {
	gorm.Model
	Id          int    `gorm:"type:int;not null;default:0;primary_key; comment:'id号'" json:"id"`
	CatName     string `gorm:"TYPE:VARCHAR(255);DEFAULT:'';comment:'猫名称'" json:"cat_name"`
	CatPrice    string `gorm:"TYPE:VARCHAR(255);DEFAULT:'0';comment:'猫价格'" json:"cat_price"`
	CatType     int    `gorm:"TYPE:INT(11);DEFAULT:0;comment:'猫价格' "json:"cat_type"`
	Description string `gorm:"TYPE:VARCHAR(255);DEFAULT:''; comment:'描述' "json:"description"`
}
type NatsCat struct {
	Id          string ` json:"id"`
	CatName     string ` json:"cat_name"`
	CatPrice    string ` json:"cat_price"`
	CatType     string `json:"cat_type"`
	Description string `json:"description"`
}
type Animals struct {
	Id         int64   ` json:"id"`
	Name       string  ` json:"name"`
	Age        int32   ` json:"age"`
	Price      float64 ` json:"price"`
	UpdateTime string  `json:"update_time"`
}
