package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 文章
type Topics struct {
	Id         int         `gorm:"primary_key"`
	Title      string      `gorm:"not null"`
	UserId     int         `gorm:"not null"`
	CategoryId int         `gorm:"not null"`
	Category   Categories  `gorm:"foreignkey:CategoryId"` //文章所属分类外键
	User       PreloadUser `gorm:"foreignkey:UserId"`     //文章所属用户外键
}

// 用户
type PreloadUser struct {
	Id   int    `gorm:"primary_key"`
	Name string `gorm:"not null"`
}

// 分类
type Categories struct {
	Id   int    `gorm:"primary_key"`
	Name string `gorm:"not null"`
}

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/gyzw_dp?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("db error:", err)
	} else {
		fmt.Println("database connection success")
	}
	//defer db.Close()
	return db
}

func main() {
	db := GetDB()
	models := []interface{}{
		&Topics{},
		&PreloadUser{},
		&Categories{},
	}
	//1.执行建表语句
	err := db.Debug().AutoMigrate(models...).Error
	if err != nil {
		fmt.Println("db error:", err)
	}
	//2.执行sql
	//INSERT INTO topics("id", "title", "user_id", "category_id") VALUES (1, '测试', 1, 1);
	//INSERT INTO categories("id", "name") VALUES (1, '测试分类');
	//INSERT INTO users("id", "name") VALUES (1, '测试用户');

	//3.执行预加载
	topics, err := GetTopicsById(db, 2)
	if err != nil {
		fmt.Println("get topics error:", err)
	}
	for i := 0; i < len(topics); i++ {
		fmt.Println(topics[i])
	}

}

func GetTopicsById(db *gorm.DB, id int) ([]*Topics, error) {
	var topic []*Topics
	//查询方法1
	//err := db.Model(&topic).Where("id=?", id).First(&topic).
	//	Related(&topic.Category, "CategoryId").
	//	Related(&topic.User, "UserId").Error

	//查询方法2
	err := db.Where("id=?", id).
		Preload("Category").
		Preload("User").
		Find(&topic).Error
	if err != nil {
		return nil, err
	}
	return topic, nil
}
