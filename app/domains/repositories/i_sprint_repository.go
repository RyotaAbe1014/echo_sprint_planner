package repositories

import (
	"echo_sprint_planner/app/domains/models"
)

type ISprintRepository interface {
	SprintCreate(sprint *models.Sprint) (err error)
	SprintList() (sprints []*models.Sprint, err error)
}
