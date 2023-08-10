package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name     string    `gorm:"type:varchar(100);"`
	Email    string    `gorm:"type:varchar(100); unique"`
	Password string    `gorm:"type:varchar(255);"`
	CreateAt time.Time `gorm:"autoCreateTime"`
	UpdateAt time.Time `gorm:"autoUpdateTime"`
	IsActive bool      `gorm:"not null"`
}
