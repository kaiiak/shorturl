package cache

import (
	"github.com/go-redis/redis"
)

// Cache redis cache
type Cache struct {
	client *redis.Client
}

// New 创建缓存
func New(addr, pwd string) (*Cache, error) {
	return &Cache{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pwd,
			DB:       0,
		}),
	}, nil
}

// Close 关闭连接
func (c *Cache) Close() error {
	return c.client.Close()
}

// Get 从缓存中读取
func (c *Cache) Get(key string) (string, error) {
	value := c.client.Get(key)
	if value.Err() == redis.Nil {
		return "", ErrNotExist
	}
	return value.String(), nil
}

// Set 写入
func (c *Cache) Set(key, vaule string) error {
	return c.client.Set(key, vaule, 0).Err()
}
