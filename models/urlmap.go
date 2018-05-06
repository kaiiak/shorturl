package models

import (
	"time"
)

// URLMap 长短链接映射表
type URLMap struct {
	ID       int64     `json:"id" gorm:"primary_key"`
	RawURL   string    `json:"raw_url"`
	ShortURL string    `json:"short_url"`
	CreateAt time.Time `json:"createtime"`
}
