package utils

import (
	"boot/model"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
)

// Rates key 存活时间
const ratesKeyExpireTime = time.Hour * 24 * 30

type RedisStorage struct {
	RedisCli *redis.Client
}

func NewRedisStorage(redisCli *redis.Client) RedisStorage {
	return RedisStorage{
		RedisCli: redisCli,
	}
}

func (storage RedisStorage) SaveValue(modelName, methodName, value string, queryModel *model.CommonQueryModel) error {
	key := "%s:%s:%s_%s_%s"
	key = fmt.Sprintf(key, modelName, methodName, queryModel.RegionCode, queryModel.StartDate, queryModel.EndDate)
	if err := storage.RedisCli.Set(key, value, ratesKeyExpireTime).Err(); err != nil {
		zap.L().Error("save value failed", zap.Error(err))
		return err
	}
	return nil
}

func (storage RedisStorage) GetValue(modelName, methodName string, queryModel *model.CommonQueryModel) (string, error) {
	key := "%s:%s:%s_%s_%s"
	key = fmt.Sprintf(key, modelName, methodName, queryModel.RegionCode, queryModel.StartDate, queryModel.EndDate)
	value, err := storage.RedisCli.Get(key).Result()
	if err != nil {
		zap.L().Error("get value failed", zap.Error(err))
		return value, err
	}

	return value, nil
}

func (storage RedisStorage) SaveValueByKey(modelName, methodName, key, value string) error {
	redisKey := "%s:%s:%s"
	redisKey = fmt.Sprintf(redisKey, modelName, methodName, key)
	if err := storage.RedisCli.Set(redisKey, value, time.Second*3600).Err(); err != nil {
		zap.L().Error("save value failed", zap.Error(err))
		return err
	}
	return nil
}
