package repositories

import (
	"echo_sprint_planner/app/domains/models"
	"echo_sprint_planner/app/domains/repositories"
	"errors"

	db "echo_sprint_planner/app/infra/db/models"

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
	dbUser := &db.User{
		Name:     user.Name,
		Email:    user.Email,
		IsActive: user.IsActive,
		Password: *user.Password,
	}

	if err := ur.db.Select("Name", "Email", "IsActive", "Password", "CreateAt").Create(&dbUser).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserList() ([]*models.User, error) {
	var users []*models.User
	var dbUsers []*db.User
	if err := ur.db.Select("ID", "Name", "Email", "IsActive", "CreateAt", "UpdateAt").Find(&dbUsers).Error; err != nil {
		return nil, err
	}

	for _, dbUser := range dbUsers {
		user := &models.User{
			ID:       &dbUser.ID,
			Name:     dbUser.Name,
			Email:    dbUser.Email,
			IsActive: dbUser.IsActive,
			CreateAt: &dbUser.CreateAt,
			UpdateAt: &dbUser.UpdateAt,
		}
		users = append(users, user)
	}

	return users, nil
}

func (ur *userRepository) UserFindByEmail(email string) (*models.User, error) {
	var dbUser db.User
	if err := ur.db.Select("ID", "Name", "Email", "Password", "IsActive", "CreateAt", "UpdateAt").Where("email = ?", email).First(&dbUser).Error; err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       &dbUser.ID,
		Name:     dbUser.Name,
		Email:    dbUser.Email,
		Password: &dbUser.Password,
		IsActive: dbUser.IsActive,
		CreateAt: &dbUser.CreateAt,
		UpdateAt: &dbUser.UpdateAt,
	}
	return user, nil
}

func (ur *userRepository) UserFindByID(userID uuid.UUID) (*models.User, error) {
	var dbUser db.User
	if err := ur.db.Select("ID", "Name", "Email", "Password", "IsActive", "CreateAt", "UpdateAt").Where("ID = ?", userID).First(&dbUser).Error; err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       &dbUser.ID,
		Name:     dbUser.Name,
		Email:    dbUser.Email,
		Password: &dbUser.Password,
		IsActive: dbUser.IsActive,
		CreateAt: &dbUser.CreateAt,
		UpdateAt: &dbUser.UpdateAt,
	}
	return user, nil
}

func (ur *userRepository) UserUpdate(user *models.User) (err error) {
	dbUser := &db.User{
		ID:       *user.ID,
		Name:     user.Name,
		Email:    user.Email,
		IsActive: user.IsActive,
	}
	if err := ur.db.Select("Name", "Email", "IsActive", "UpdateAt").Updates(&dbUser).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) UserDelete(id uuid.UUID) (err error) {
	if tx := ur.db.Select("ID").Delete(&db.User{}, id); tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return nil
}
