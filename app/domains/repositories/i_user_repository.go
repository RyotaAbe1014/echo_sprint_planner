package repositories

import (
	"echo_sprint_planner/app/domains/models"
	"time"

	"github.com/google/uuid"
)

type IUserRepository interface {
	UserCreate(user *models.User) (err error)
	GetUserList() ([]*models.User, error)
	UserFindByEmail(email string) (*models.User, error)
	UserUpdate(user *models.User) (err error)
	UserDelete(id uuid.UUID) (err error)
}

type MockUserRepository struct{}

func NewMockUserRepository() IUserRepository {
	return &MockUserRepository{}
}

func (mur *MockUserRepository) UserCreate(user *models.User) (err error) {
	return nil
}

// UserUpdate
func (mur *MockUserRepository) UserUpdate(user *models.User) (err error) {
	return nil
}

// UserDelete
func (mur *MockUserRepository) UserDelete(id uuid.UUID) (err error) {
	return nil
}

// GetUserList
func (mur *MockUserRepository) GetUserList() ([]*models.User, error) {
	var users []*models.User
	user := &models.User{
		ID:       &uuid.UUID{},
		Name:     "test",
		Email:    "test@test.com",
		IsActive: true,
		CreateAt: &time.Time{},
		UpdateAt: &time.Time{},
	}
	users = append(users, user)
	return users, nil
}

// UserFindByEmail
func (mur *MockUserRepository) UserFindByEmail(email string) (*models.User, error) {
	user := &models.User{
		ID:       &uuid.UUID{},
		Name:     "test",
		Email:    "test@test.com",
		IsActive: true,
		CreateAt: &time.Time{},
		UpdateAt: &time.Time{},
	}
	return user, nil
}
