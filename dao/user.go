package dao

import (
	"starter/service"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func NewUserDao(db *gorm.DB, cache *service.CacheService, logger *zap.Logger) *userDao {
	return &userDao{
		db:     db,
		logger: logger,
		cache:  cache,
	}
}

type userDao struct {
	db     *gorm.DB
	logger *zap.Logger
	cache  *service.CacheService
}
