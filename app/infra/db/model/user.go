package model

import (
	"time"
)

type User struct {
	ID       int64     `gorm:"primary_key;AUTO_INCREMENT"`
	Name     string    `gorm:"type:varchar(100);"`
	Email    string    `gorm:"type:varchar(100);"`
	CreateAt time.Time `gorm:"autoCreateTime"`
	UpdateAt time.Time `gorm:"autoUpdateTime"`
	IsActive bool      `gorm:"not null"`
}
