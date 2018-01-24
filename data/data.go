package data

import (
	"github.com/kaiiak/shorturl/config"
	"github.com/kaiiak/shorturl/data/cache"
	"github.com/kaiiak/shorturl/data/db"
)

// Data 数据管理
type Data struct {
	db    *db.ShortURLDB
	cache *cache.Cache
}

// New 控制数据库和缓存
func New(cnf *config.Config) (*Data, error) {
	tempDb, err := db.New(cnf.DBType, cnf.DBPath)
	if err != nil {
		return nil, err
	}
	tempCache, err := cache.New(cnf.CachePath, cnf.CachePwd)
	if err != nil {
		return nil, err
	}
	return &Data{tempDb, tempCache}, nil
}

// Get 获取数据
func (d *Data) Get(shorturlStr string) (string, error) {
	v, err := d.cache.Get(shorturlStr)
	if err == nil {
		return v, nil
	}
	v, err = d.db.Get(shorturlStr)
	if err == nil {
		return v, nil
	}
	if err == db.ErrNotExist {
		return "", ErrNotFound
	}
	return "", err
}

// Set 设置,如果已存在则返回原始的
func (d *Data) Set(raw string) (string, error) {
	return "", nil
}

// Close 安全的关闭
func (d *Data) Close() error {
	d.cache.Close()
	d.db.CLose()
	return nil
}
