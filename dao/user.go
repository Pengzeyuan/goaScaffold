package dao

import (
	"boot/pkg/cache"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewUserDao(db *gorm.DB, cache *cache.CacheService, logger *zap.Logger) *userDao {
	return &userDao{
		db:     db,
		logger: logger,
		cache:  cache,
	}
}

type userDao struct {
	db     *gorm.DB
	logger *zap.Logger
	cache  *cache.CacheService
}
