package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/piovani/kafka_aux/config"
)

type Cache struct {
	redis *redis.Client
}

func NewCache() *Cache {
	return &Cache{
		redis: redis.NewClient(&redis.Options{
			Addr:     config.Env.RedisHost,
			Username: "",
			Password: config.Env.RedisPassword,
		}),
	}
}

func (c *Cache) Set(key string, value any, ttl time.Duration) error {
	return c.redis.Set(context.TODO(), key, value, ttl).Err()
}

func (c *Cache) Get(key string) (string, error) {
	return c.redis.Get(context.TODO(), key).Result()
}

func (c *Cache) GetAllKeys() (keys []string, err error) {
	cursor := uint64(0)		
	keys, cursor, err = c.redis.Scan(context.TODO(), cursor, "", 0).Result()
	return keys, err
}