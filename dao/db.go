package dao

import (
	"fmt"
	"github.com/withlin/canal-go/client"
	"time"

	"boot/config"
	"boot/model"
	// init mysql driver
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	DpDB      *gorm.DB                     // 政务大屏二期需要使用的独立于一期的数据库
	Connector *client.SimpleCanalConnector // canal数据库增量订阅消费

)

func InitDB(cfg *config.Config) {
	DpDB = GetMyDB(cfg.Debug, cfg.DpDatabase)

}
func InitCanal(cfg *config.Config) {
	Connector = GetMyCanal(cfg.Canal)
	ab := Connector
	fmt.Println(ab)
}
func GetMyCanal(c config.CanalConf) *client.SimpleCanalConnector {
	var err error
	zap.L().Debug("connect Canal ...", zap.String("dsn", c.Destination))
	switch c.Host {
	case "0.0.0.0", "":
		c.Host = "127.0.0.1"
	default:
	}
	if c.Port == 0 {
		c.Port = 8086
	}
	connector := client.NewSimpleCanalConnector(c.Host, c.Port, "", "", c.Destination, c.SoTimeOut, c.IdleTimeOut)
	err = connector.Connect()
	if err != nil {
		zap.L().Panic("connect Canal failed", zap.Error(err))
		return nil

	}
	zap.L().Debug("canal 连接成功......")

	err = connector.Subscribe("ginessential\\.animals")
	if err != nil {
		zap.L().Panic("observe specific table failed", zap.Error(err))
		return nil

	}
	return connector
}
func GetMyDB(dbLogMode bool, dbConfig config.DatabaseConfig) *gorm.DB {
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

//func InitDB(cfg *config.Config) {
//	zap.L().Debug("connect db", zap.String("dsn", cfg.Database.DSN))
//
//	var dbLogger logger.Interface
//
//	if cfg.Debug {
//		dbLogger = logger.Default.LogMode(logger.Info)
//	} else {
//		dbLogger = logger.Default.LogMode(logger.Warn)
//	}
//
//	var err error
//
//	DB, err = gorm.Open(mysql.Open(cfg.Database.DSN), &gorm.Config{
//		Logger: dbLogger,
//		NamingStrategy: schema.NamingStrategy{
//			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
//		},
//	})
//	if err != nil {
//		zap.L().Panic("connect db failed", zap.Error(err))
//	}
//
//	sqlDB, err := DB.DB()
//	if err != nil {
//		zap.L().Panic("connect db failed", zap.Error(err))
//	}
//
//	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
//	if cfg.Database.MaxIdleConns > 0 {
//		sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
//	}
//
//	// SetMaxOpenCons 设置数据库的最大连接数量。
//	if cfg.Database.MaxOpenConns > 0 {
//		sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
//	}
//
//	// SetConnMaxLifetiment 设置连接的最大可复用时间。
//	if cfg.Database.ConnMaxLifetime != "" {
//		maxLifetime, err := time.ParseDuration(cfg.Database.ConnMaxLifetime)
//		if err != nil {
//			zap.L().Panic("db ConnMaxLifetime parse failed", zap.Error(err))
//		}
//
//		sqlDB.SetConnMaxLifetime(maxLifetime)
//	}
//
//	if err := sqlDB.Ping(); err != nil {
//		zap.L().Panic("ping db failed", zap.Error(err))
//	}
//}

func CloseDB() {
	if DpDB != nil {

	}
}

func AutoMigrateDB() {
	query := DpDB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci")
	if err := query.AutoMigrate(
		&model.AopUser{},
		&model.WindowInfo{},
		&model.TakeNumber{},
		&model.CallNumber{},
		&model.HallManagementInfo{},
		&model.LegalPersonUser{},
		&model.CompanyProfile{},
	).Error; err != nil {
		zap.L().Panic("migrate db fail", zap.Error(err))
	}
}
