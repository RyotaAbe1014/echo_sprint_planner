package services

import (
	"echo_sprint_planner/app/domains/models"
	"echo_sprint_planner/app/domains/repositories"
	"time"
)

type ISprintService interface {
	SprintCreate(name string, startDate time.Time, endDate time.Time, updatedBy string) (err error)
}

type sprintService struct {
	ur repositories.ISprintRepository
}

// constructorを使用して、serviceの構造体を生成
func NewSprintService(sr repositories.ISprintRepository) ISprintService {
	return &sprintService{sr}
}

func (us *sprintService) SprintCreate(name string, startDate time.Time, endDate time.Time, updatedBy string) (err error) {
	sprint := &models.Sprint{
		Name:      &name,
		StartDate: startDate,
		EndDate:   endDate,
		UpdatedBy: &updatedBy,
	}

	if err := us.ur.SprintCreate(sprint); err != nil {
		return err
	}
	return nil
}
