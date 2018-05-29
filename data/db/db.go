package db

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kaiiak/shorturl/models"

	// 现在只支持 mysql, postgresql, sqlite
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// ShortURLDB 短链接存储数据库
type ShortURLDB struct {
	DB *gorm.DB
}

// New 创建新的数据库对象
func New(typeStr, pathStr string) (*ShortURLDB, error) {
	d, err := gorm.Open(typeStr, pathStr)
	if err != nil {
		return nil, err
	}
	d = d.AutoMigrate(&models.URLMap{})
	return &ShortURLDB{
		DB: d,
	}, nil
}

// Get 获取
func (d *ShortURLDB) Get(key string) (string, error) {
	um := &models.URLMap{ShortURL: key}
	if err := d.DB.First(um).Error; err != nil {
		return "", err
	}
	return um.RawURL, nil
}

// Set 设置
func (d *ShortURLDB) Set(key string) (*models.URLMap, error) {
	um := &models.URLMap{RawURL: key, CreateAt: time.Now()}
	if err := d.DB.FirstOrCreate(um).Error; err != nil {
		return nil, err
	}
	return um, nil
}

// CLose close database
func (d *ShortURLDB) CLose() error {
	return d.DB.Close()
}
