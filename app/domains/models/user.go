package models

import (
	"time"
)

type User struct {
	Name     string
	Email    string
	IsActive bool
	Password string
	CreateAt *time.Time
	UpdateAt *time.Time
}
