package service

import (
	"starter/config"

	"time"

	cache "github.com/go-redis/cache/v7"
	"github.com/go-redis/redis/v7"
	msgpack "github.com/vmihailenco/msgpack/v4"
	"go.uber.org/zap"
)

type CacheService struct {
	logger *zap.Logger
	codec  *cache.Codec
}

func NewCacheService(logger *zap.Logger) *CacheService {
	opts, _ := redis.ParseURL(config.C.Redis.URI)
	rdb := redis.NewClient(opts)

	codec := &cache.Codec{
		Redis: rdb,
		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	return &CacheService{
		logger: logger,
		codec:  codec,
	}
}

func (cs *CacheService) Get(key string, wanted interface{}) error {
	return cs.codec.Get(key, &wanted)
}

func (cs *CacheService) Set(key string, obj interface{}, exp time.Duration) error {
	return cs.codec.Set(&cache.Item{
		Key:        key,
		Object:     obj,
		Expiration: exp,
	})
}
