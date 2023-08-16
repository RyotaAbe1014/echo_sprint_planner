package repositories

import (
	"echo_sprint_planner/app/domains/models"

	"github.com/google/uuid"
)

type ISprintRepository interface {
	SprintCreate(sprint *models.Sprint) (err error)
	SprintList() (sprints []*models.Sprint, err error)
	SprintUpdate(sprint *models.Sprint) (err error)
	SprintDelete(sprintId uuid.UUID) (err error)
}
