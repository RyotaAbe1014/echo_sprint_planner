package models

import (
	"time"

	"github.com/google/uuid"
)

type Sprint struct {
	ID        uuid.UUID `gorm:"primaryKey;autoIncrement"`
	Name      *string   `gorm:"size:255"`
	StartDate time.Time `gorm:"type:date"`
	EndDate   time.Time `gorm:"type:date"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	UpdatedBy *string   `gorm:"size:255"`
}
