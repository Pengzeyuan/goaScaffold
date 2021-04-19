package initDB

import (
	"basic/model"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm/schema"

	"gorm.io/gorm/clause"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	gorm2 "gorm.io/gorm"
)

var (
	DpDB *gorm.DB
	C    *DatabaseConfig
	DB2  *gorm2.DB
)

type DatabaseConfig struct {

	// 仅支持 mysql
	DSN          string `yaml:"dsn"`
	MaxIdleConns int    `yaml:"max_idle_conns" default:"10"`
	MaxOpenConns int    `yaml:"max_open_conns" default:"100"`
	// format: https://golang.org/pkg/time/#ParseDuration
	ConnMaxLifetime string `yaml:"conn_max_lifetime" default:"1h"`
}

func InitDB() {
	C = &DatabaseConfig{}
	if err := configor.New(&configor.Config{AutoReload: true}).Load(C, "cmd/config.yml"); err != nil {
		zap.L().Panic("init config fail", zap.Error(err))
	}

	DpDB = GetMyDB(true, C)
	// 自动迁移
	query := DpDB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci")
	if err := query.AutoMigrate(
		model.Employee{},
		model.User{},
		model.Company{},
		model.Article{},
		model.Category{},
		model.Tag{},
		model.ArticleTag{},
		model.Profile{},
		model.Demo{},
	).Error; err != nil {
		zap.L().Panic("migrate dp db fail", zap.Error(err))
	}
}

func GetMyDB(dbLogMode bool, dbConfig *DatabaseConfig) *gorm.DB {
	var err error
	zap.L().Debug("connect db ...", zap.String("dsn", dbConfig.DSN))
	myDB, err := gorm.Open("mysql", dbConfig.DSN)
	if err != nil {
		zap.L().Panic("connect db failed", zap.Error(err))
		return nil
	}
	if dbLogMode {
		myDB.LogMode(dbLogMode)
	}

	myDB.SingularTable(true)
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	if dbConfig.MaxIdleConns > 0 {
		myDB.DB().SetMaxIdleConns(dbConfig.MaxIdleConns)
	}

	// SetMaxOpenCons 设置数据库的最大连接数量。
	if dbConfig.MaxOpenConns > 0 {
		myDB.DB().SetMaxOpenConns(dbConfig.MaxOpenConns)
	}
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	if dbConfig.ConnMaxLifetime != "" {
		maxLifetime, err := time.ParseDuration(dbConfig.ConnMaxLifetime)
		if err != nil {
			zap.L().Panic("dp ConnMaxLifetime parse failed", zap.Error(err))
			return nil
		}

		myDB.DB().SetConnMaxLifetime(maxLifetime)
	}

	if err := myDB.DB().Ping(); err != nil {
		zap.L().Panic("ping db failed", zap.Error(err))
		return nil
	}
	zap.L().Debug("connected db", zap.String("dsn", dbConfig.DSN))
	return myDB
}

func InitGorm2() {
	C = &DatabaseConfig{}
	if err := configor.New(&configor.Config{AutoReload: true}).Load(C, "cmd/config.yml"); err != nil {
		zap.L().Panic("init config fail", zap.Error(err))
	}
	var err error
	DB2, err = gorm2.Open(mysql.Open("root:root@(localhost:3306)/easymall?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm2.Config{
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
		})
	if err != nil {
		zap.L().Panic("ping db failed", zap.Error(err))

	}
	fmt.Println(DB2)
	//gorm2Test1(err, db)

	//Gorm2Test2(db)

	//Gorm2Test3(db)
	//var users = []model.Employee{}
	//Gorm2Test4(db)

	//user1, user2 := test5(db)

	//Test5Onduplicate()

}

func Test5Onduplicate() {
	db := DB2
	var user1 = model.Employee{Id: 24, UserName: "探戈1"}
	var user2 = model.Employee{Id: 25, UserName: "探戈2"}
	var user = model.Employee{}
	tx := db.Begin()
	tx.Create(&user1)

	tx.SavePoint("sp1")
	tx.Create(&user2)
	tx.RollbackTo("sp1") // rollback user2

	tx.Commit() // commit user1

	find := db.Where("name = @name OR addr = @name", sql.Named("name", "jinzhu")).Find(&user)
	findstr := db.Session(&gorm2.Session{DryRun: true}).Where("name = @name OR addr = @name", sql.Named("name", "jinzhu")).Find(&user).Statement
	fmt.Printf("vars: %+v \n", find)
	fmt.Printf("findstr: %+v \n", findstr.SQL.String())

	var users = []model.Employee{}
	db.Where(
		db.Where("id = ?", "1").Where(db.Where("age = ?", 18).Or("id = ?", "2")),
	).Or(
		db.Where("id = ?", "3").Where("name = ?", "jinzhu2"),
	).Find(&users)

	db.Where("age > (?)", db.Table("employee").Select("AVG(age) as avgAge")).Find(&users)

	//db.Table("(?) as u", db.Model(&model.Employee{}).Select("name", "age")).Where("age = ?", 18).Find(&users)

	users = []model.Employee{{Id: 1, UserName: "金大珠1"}, {Id: 26, UserName: "金大珠2"}, {Id: 3, UserName: "金大珠3", Age: 12}}
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
	}).Create(&users)
}

func test5(db *gorm2.DB) (model.Employee, model.Employee) {
	var user1 = model.Employee{Id: 21, UserName: "飞机1"}
	var user2 = model.Employee{Id: 22, UserName: "飞机2"}
	var user3 = model.Employee{Id: 23, UserName: "飞机3"}

	db.Transaction(func(tx *gorm2.DB) error {
		tx.Create(&user1)

		tx.Transaction(func(tx2 *gorm2.DB) error {
			tx.Create(&user2)
			return errors.New("rollback user2") // rollback user2
		})

		tx.Transaction(func(tx2 *gorm2.DB) error {
			tx.Create(&user3)
			return nil
		})

		return nil // commit user1 and user3
	})
	return user1, user2
}

func Gorm2Test4(db *gorm2.DB) {
	var user = model.Employee{Id: 2}
	////var result map[string]interface{}
	result := make(map[string]interface{}, 10)
	db.Model(&user).First(&result, "id = ?", 1)
	//
	//db.Model(&model.Employee{}).Create(map[string]interface{}{"Name": "猪猪", "Age": 28})
	//
	//datas := []map[string]interface{}{
	//	{"Name": "jinzhu_1猪", "Age": 19},
	//	{"name": "jinzhu_2猪", "Age": 20},
	//}
	//
	//db.Model(&model.Employee{}).Create(datas)
	//
	//find := db.Session(&gorm2.Session{DryRun: true}).Joins("employee").Joins("do_process").Joins("categories").Find(&users, "users.id IN ?", []int{1, 2}).Statement
	//fmt.Printf("%s %+v", find.SQL.String(), find.SQL.String())
	//fmt.Printf("vars:%s %+v \n", find, find)

	resultDB := db.Where("age>?", 13).FindInBatches(&result, 100, func(tx *gorm2.DB, batch int) error {
		// 批量处理
		for i := 0; i < 10; i++ {
			fmt.Printf("vars: %+v \n", 1)
		}
		return nil
	})
	fmt.Printf("vars: %+v \n", &resultDB)
}

func Gorm2Test3(db *gorm2.DB) {
	var user = model.Employee{UserName: "jinzhu1", Id: 1}
	stmt := db.Session(&gorm2.Session{DryRun: true}).Find(&user, 1).Statement
	stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = $1 // PostgreSQL
	stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = ?  // MySQL
	fmt.Printf("%s %+v", stmt.SQL.String(), stmt.SQL.String())
	fmt.Printf("vars:%s %+v \n", stmt.Vars, stmt.Vars)
	//=> []interface{}{1}
}

func Gorm2Test2(db *gorm2.DB) {
	// 批量插入
	var users = []model.Employee{{UserName: "jinzhu1"}, {UserName: "jinzhu2"}, {UserName: "jinzhu3"}}
	db.Create(&users)

	for _, user := range users {
		fmt.Printf("id:%d", user.Id) // 1,2,3
	}
	//使用 CreateInBatches 创建时，你还可以指定创建的数量，例如：

	var users2 = []model.Employee{{UserName: "jinzhu_1"}, {UserName: "jinzhu_10000"}}

	// 数量为 100
	db.CreateInBatches(users2, 100)
}

func gorm2Test1(err error, db *gorm2.DB) {
	//大部分 CRUD API 都是兼容的

	err = db.AutoMigrate(&model.Employee{})
	if err != nil {
		zap.L().Panic("AutoMigrate db failed", zap.Error(err))

	}

	oneEmp := model.Employee{Id: 1, UserName: "小白", Addr: "青玉路"}
	db.Create(&oneEmp)
	first := db.First(&model.Employee{}, 1)
	fmt.Println(first)
	db.Model(&oneEmp).Update("Age", 18)
	db.Model(&oneEmp).Omit("Age").Updates(map[string]interface{}{"Name": "jinzhu", "Addr": "adminaddr"})
}
