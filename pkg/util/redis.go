package util

import (
	"time"

	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
)

type RedisStorage struct {
	RedisCli *redis.Client
}

func NewRedisStorage(redisCli *redis.Client) RedisStorage {
	return RedisStorage{
		RedisCli: redisCli,
	}
}

// GetCode 通过key,从redis中获取验证码
// key 验证码保存时的key
func (storage RedisStorage) GetCode(key string) (string, error) {
	value, err := storage.RedisCli.Get(key).Result()
	if err != nil {
		zap.L().Error("get code failed", zap.Error(err))
		return value, err
	}

	return value, nil
}

// SaveCode 保存验证码到redis
// key 验证码保存时的key
// code 验证码
func (storage RedisStorage) SaveCode(key, value string) error {

	if err := storage.RedisCli.Set(key, value, time.Second*3600).Err(); err != nil {
		zap.L().Error("save code failed", zap.Error(err))
		return err
	}
	return nil
}
