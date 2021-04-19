package model

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	//这里grom是映射到表
	ID        int        `gorm:"TYPE:int(11);NOT NULL;PRIMARY_KEY;INDEX"`
	UserName  string     `gorm:"TYPE: VARCHAR(255); DEFAULT:'';INDEX"`
	NickName  string     `gorm:"TYPE: VARCHAR(255); DEFAULT:'无' "`
	Companies []Company  `gorm:"FOREIGNKEY:UserId;ASSOCIATION_FOREIGNKEY:ID"`
	CreatedAt time.Time  `gorm:"TYPE:DATETIME"`
	UpdatedAt time.Time  `gorm:"TYPE:DATETIME"`
	DeletedAt *time.Time `gorm:"TYPE:DATETIME;DEFAULT:NULL"`
	IsFemale  int        `gorm:"TYPE:int(11);DEFAULT:1"` // 数据库中默认值为 1
	IsActived int        `gorm:"TYPE:int(11);DEFAULT:1"` // 数据库中默认值为 1
	Password  string     `gorm:"TYPE: VARCHAR(255); column:password; DEFAULT:'无' "`
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.NickName == "tom1" {
		return errors.New("admin user not allowed to update")
	}
	return
}

// 个性化信息表
type Profile struct {
	gorm.Model
	UserID uint // 外键
	// 定义user属性关联users表，默认情况使用 类型名 + ID 组成外键名，在这里UserID属性就是外键
	User User //`gorm:"foreignkey:UserRefer"` //使用 UserRefer 作为外键
	Name string
}

type Company struct {
	gorm.Model
	Industry int    `gorm:"TYPE:INT(11);DEFAULT:0"`
	Name     string `gorm:"TYPE:VARCHAR(255);DEFAULT:'';INDEX"`
	Job      string `gorm:"TYPE:VARCHAR(255);DEFAULT:''"`
	UserId   int    `gorm:"TYPE:int(11);NOT NULL;INDEX"`
}

//这里的Articles和Labels分别代表表名

//文章表
type Article struct {
	Id         int      `json:"id"   gorm:"TYPE:int(11);NOT NULL;PRIMARY_KEY;INDEX"`
	Title      string   `json:"title" gorm:"TYPE: VARCHAR(255); DEFAULT:'';INDEX"`
	CategoryId int      `json:"category_id" gorm:"TYPE:INT(11);DEFAULT:0"`
	Category   Category `json:"category";gorm:"foreignkey:CategoryID"` //指定关联外键
	Tag        []Tag    `gorm:"many2many:article_tag" json:"tag"`      //多对多关系.
	//article_tag表默认article_id字段对应article表id.tag_id字段对应tag表id
	//可以把对应sql日志打印出来,便于调试
}

//远程一对多.一对一

//文章_标签中间表
type ArticleTag struct {
	Id        int       `json:"id"  gorm:"TYPE:int(11);NOT NULL;PRIMARY_KEY;INDEX" `
	ArticleId string    `json:"article_id" gorm:"TYPE: VARCHAR(255); DEFAULT:'';INDEX"`
	TagId     string    `json:"tag_id" gorm:"TYPE: VARCHAR(255); DEFAULT:'';INDEX"`
	CreatedAt time.Time `json:"created_at"   gorm:"TYPE:DATETIME" `
	UpdatedAt time.Time `json:"updated_at"   gorm:"TYPE:DATETIME" `
}

//标签表
type Tag struct {
	Id      int    `json:"id" gorm:"TYPE:int(11);NOT NULL;PRIMARY_KEY;INDEX"`
	TagName string `json:"tag_name" gorm:"TYPE: VARCHAR(255); DEFAULT:'';INDEX"`
}

//分类表
type Category struct {
	ID           int       `json:"id" gorm:"TYPE:int(11);NOT NULL;PRIMARY_KEY;INDEX"`
	CategoryName string    `json:"category_name" gorm:"TYPE: VARCHAR(255); DEFAULT:'';INDEX"`
	Status       int       `json:"status"`
	CreatedAt    time.Time `json:"created_at"  gorm:"TYPE:DATETIME" `
	UpdatedAt    time.Time `json:"updated_at"  gorm:"TYPE:DATETIME" `
}
