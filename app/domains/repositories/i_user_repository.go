package repositories

import "echo_sprint_planner/app/domains/models"

type IUserRepository interface {
	UserCreate(user *models.User) (err error)
	GetUserList() ([]*models.User, error)
}
