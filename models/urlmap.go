package models

import (
	"time"
)

// URLMap 长短链接映射表
type URLMap struct {
	CreateTime time.Time `json:"createtime"`
	RawURL     string    `json:"raw_url"`
	ShortURL   string    `json:"short_url"`
	ID         int64     `json:"id"`
}
