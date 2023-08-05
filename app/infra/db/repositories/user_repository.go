package repositories

import (
	"echo_sprint_planner/app/domains/models"
	"echo_sprint_planner/app/domains/repositories"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) UserCreate(user *models.User) (err error) {
	if err := ur.db.Select("Name", "Email", "IsActive").Create(&user).Error; err != nil {
		return err
	}
	return nil
}
