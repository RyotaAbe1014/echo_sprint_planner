package repositories

import (
	"echo_sprint_planner/app/domains/models"
	"echo_sprint_planner/app/domains/repositories"

	"gorm.io/gorm"
)

type sprintRepository struct {
	db *gorm.DB
}

func NewSprintRepository(db *gorm.DB) repositories.ISprintRepository {
	return &sprintRepository{db}
}

func (sr *sprintRepository) SprintCreate(sprint *models.Sprint) (err error) {
	return nil
}
