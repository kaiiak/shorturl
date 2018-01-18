package db

import (
	"github.com/jinzhu/gorm"
	"github.com/kaiiak/shorturl/models"

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
