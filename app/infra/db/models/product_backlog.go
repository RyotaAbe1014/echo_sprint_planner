package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductBacklog struct {
	ID        uuid.UUID `gorm:"primaryKey;autoIncrement"`
	Title     *string   `gorm:"size:255"`
	Progress  int64     `gorm:"type:smallint"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	UpdatedBy *string   `gorm:"size:255"`
	SprintID  uuid.UUID `gorm:"index"` // note the capitalization
	Sprint    *Sprint   `gorm:"foreignKey:SprintID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
