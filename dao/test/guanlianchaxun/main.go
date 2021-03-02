package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type MaintenanceOrder struct {
	OrderId     string `gorm:"type:varchar(50);null"` //订单编号
	OrderStatus int    //
	Images      []ImageTable
	ImageId     string `gorm:"ForeignKey:Id"`
	// 报修人信息
}
type ImageTable struct {
	order    *MaintenanceOrder `gorm:"ForeignKey:OrderId"`
	Id       uint              `gorm:"primary_key"`
	ImageUrl string
	RemarkId string
}

type Posts struct {
	gorm.Model
	//Key string `gorm:"unique:not null"`
	Userid int
	Title  string `gorm:"type:varchar(200)"`
}

type PostUsers struct {
	gorm.Model
	Username string `gorm:"unique_index"`
	// 一对多关系映射，一个用户有多篇文章
	PostsArticle []Posts `gorm:"FOREIGNKEY:Userid;ASSOCIATION_FOREIGNKEY:ID;ON DELETE CASCADE" `
}

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:root@(localhost:3306)/gyzw_dp?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='图片表'").AutoMigrate(&ImageTable{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='图片表'").AutoMigrate(&MaintenanceOrder{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='图片表'").AutoMigrate(&Posts{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='图片表'").AutoMigrate(&PostUsers{})

}

func main() {
	var user = PostUsers{
		Username: "h7",
		PostsArticle: []Posts{
			{
				Title: "这是一首简单的小情歌",
			},
			{
				Title: "东风破",
			},
		},
	}

	DB.Preload("PostsArticle").Where("id=1").Find(&user)
	//err := DB.Create(&user).Error
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	fmt.Println(user, user.Username, user.PostsArticle)
	fmt.Println(user.PostsArticle)
	user2 := PostUsers{}
	user2.ID = 1
	DB.Where("username= ?", "h7").Delete(&user2)

}
