package testing

import (
	"boot/dao"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func DbCnnForTest(t *testing.T) *gorm.DB {
	viper.AutomaticEnv()
	dsn := "root:root@(127.0.0.1:3306)/gyzw_dp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		t.Error(err)
		zap.L().Panic("connect db failed", zap.Error(err))
	}
	logModel := true
	db.LogMode(logModel)
	dao.DpDB = db

	return dao.DpDB
}
