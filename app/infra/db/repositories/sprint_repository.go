package repositories

import (
	"echo_sprint_planner/app/domains/models"
	"echo_sprint_planner/app/domains/repositories"

	db "echo_sprint_planner/app/infra/db/models"

	"gorm.io/gorm"
)

type sprintRepository struct {
	db *gorm.DB
}

func NewSprintRepository(db *gorm.DB) repositories.ISprintRepository {
	return &sprintRepository{db}
}

func (sr *sprintRepository) SprintCreate(sprint *models.Sprint) (err error) {
	dbSprint := &db.Sprint{
		Name:      sprint.Name,
		StartDate: sprint.StartDate,
		EndDate:   sprint.EndDate,
		UpdatedBy: sprint.UpdatedBy,
	}

	if err := sr.db.Select("Name", "StartDate", "EndDate", "UpdatedBy").Create(&dbSprint).Error; err != nil {
		return err
	}
	return nil
}

func (sr *sprintRepository) SprintList() (sprints []models.Sprint, err error) {
	return nil, nil
}
