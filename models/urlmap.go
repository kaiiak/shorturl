package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kaiiak/shorturl/util"
)

// URLMap 长短链接映射表
type URLMap struct {
	ID       int64     `json:"id" gorm:"primary_key"`
	RawURL   string    `json:"raw_url" gorm:"unique_index"`
	ShortURL string    `json:"short_url" gorm:"unique_index"`
	CreateAt time.Time `json:"createtime"`
}

// AfterCreate gorm hook
func (u *URLMap) AfterCreate(scope *gorm.Scope) (err error) {
	shortURL := util.ConvertTo62(u.ID+10000)
	u.ShortURL = shortURL
	return scope.DB().Model(u).Update("short_url", shortURL).Error
}
