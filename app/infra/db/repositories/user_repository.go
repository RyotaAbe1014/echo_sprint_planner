package repositories

import (
	"echo_sprint_planner/app/domains/models"
	"echo_sprint_planner/app/domains/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) UserCreate(user *models.User) (err error) {
	if err := ur.db.Select("Name", "Email", "IsActive", "Password", "CreateAt").Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserList() ([]*models.User, error) {
	var users []*models.User
	if err := ur.db.Select("ID", "Name", "Email", "IsActive", "CreateAt").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) UserUpdate(user *models.User) (err error) {
	if err := ur.db.Select("Name", "Email", "IsActive", "UpdateAt").Updates(&user).Error; err != nil {
		return err
	}
	return nil
}
func (ur *userRepository) UserDelete(id uuid.UUID) (err error) {
	if err := ur.db.Select("ID").Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
