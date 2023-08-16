package repositories

import (
	"echo_sprint_planner/app/domains/models"
	"echo_sprint_planner/app/domains/repositories"
	"errors"

	dbModel "echo_sprint_planner/app/infra/db/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type sprintRepository struct {
	db *gorm.DB
}

func NewSprintRepository(db *gorm.DB) repositories.ISprintRepository {
	return &sprintRepository{db}
}

func (sr *sprintRepository) SprintCreate(sprint *models.Sprint) (err error) {
	dbSprint := &dbModel.Sprint{
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

func (sr *sprintRepository) SprintList() (sprints []*models.Sprint, err error) {
	dbSprints := []dbModel.Sprint{}
	if err := sr.db.Select("ID", "Name", "StartDate", "EndDate", "CreatedAt", "UpdatedAt", "UpdatedBy").Find(&dbSprints).Error; err != nil {
		return nil, err
	}

	for _, dbSprint := range dbSprints {
		sprints = append(sprints, &models.Sprint{
			Name:      dbSprint.Name,
			StartDate: dbSprint.StartDate,
			EndDate:   dbSprint.EndDate,
			UpdatedBy: dbSprint.UpdatedBy,
		})
	}
	return sprints, nil
}

func (sr *sprintRepository) SprintUpdate(sprint *models.Sprint) (err error) {
	dbSprint := &dbModel.Sprint{
		ID:        sprint.ID,
		Name:      sprint.Name,
		StartDate: sprint.StartDate,
		EndDate:   sprint.EndDate,
		UpdatedBy: sprint.UpdatedBy,
	}

	if err := sr.db.Select("Name", "StartDate", "EndDate", "UpdatedBy").Updates(&dbSprint).Error; err != nil {
		return err
	}

	return nil
}

func (sr *sprintRepository) SprintDelete(sprintId uuid.UUID) (err error) {
	if tx := sr.db.Select("ID").Delete(&dbModel.Sprint{}, sprintId); tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return nil
}
