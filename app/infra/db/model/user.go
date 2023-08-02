package model

import (
	"time"
)

type User struct {
	Id       int64     `gorm:"primary_key;AUTO_INCREMENT"`
	Name     string    `gorm:"type:varchar(100);not null"`
	Email    string    `gorm:"type:varchar(100);not null"`
	CreateAt time.Time `gorm:"not null"`
	UpdateAt time.Time `gorm:"not null"`
	IsActive bool      `gorm:"not null"`
}
