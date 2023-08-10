package repositories

import (
	"echo_sprint_planner/app/domains/models"

	"github.com/google/uuid"
)

type IUserRepository interface {
	UserCreate(user *models.User) (err error)
	GetUserList() ([]*models.User, error)
	UserUpdate(user *models.User) (err error)
	UserDelete(id uuid.UUID) (err error)
}
