package models

import (
	"time"

	"github.com/google/uuid"
)

type Sprint struct {
	ID        uuid.UUID  `json:"id"`
	Name      *string    `json:"name"`
	StartDate time.Time  `json:"start_date"`
	EndDate   time.Time  `json:"end_date"`
	CreatedAt *time.Time `json:"create_at"`
	UpdatedAt *time.Time `json:"update_at"`
	UpdatedBy *string    `json:"update_by"`
}
