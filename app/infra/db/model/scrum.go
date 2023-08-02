package model

import (
	"time"
)

type Sprint struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Name      *string   `gorm:"size:255"`
	StartDate time.Time `gorm:"type:date"`
	EndDate   time.Time `gorm:"type:date"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	UpdatedBy *string   `gorm:"size:255"`
}