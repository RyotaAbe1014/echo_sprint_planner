package model

import (
	"time"
)

// User ユーザー情報
type User struct {
	ID uint `gorm:"primary_key"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Name string `gorm:"size:255"`
}
