package data

import (
	"github.com/kaiiak/shorturl/data/cache"
	"github.com/kaiiak/shorturl/data/db"
	"github.com/kaiiak/shorturl/config"
)

// Data 数据管理
type Data struct {
	db *db.ShortURLDB
	cache *cache.Cache
}

// New 控制数据库和缓存
func New(cnf *config.Config) (*Data, error) {
	return nil, nil
}

// Get 获取数据
func (d *Data) Get(shorturlStr string) (string, error) {
	return "", nil
}

// Set 设置,如果已存在则返回原始的
func (d *Data) Set(raw string) (string, error) {
	return "", nil
}