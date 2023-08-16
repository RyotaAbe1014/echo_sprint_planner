package services

import (
	"echo_sprint_planner/app/domains/models"
	"echo_sprint_planner/app/domains/repositories"
	"time"

	"github.com/google/uuid"
)

// class SprintUpdateService():
//     sprint_repository: ISprintRepository = inject.attr(ISprintRepository)

//     def delete_sprint(self, sprint_id: SprintId) -> None:
//         self.sprint_repository.delete_sprint(sprint_id=sprint_id)

//     def update_sprint(self, sprint_id: SprintId, name: str, start_date: str, end_date: str, updated_by: str) -> None:
//         sprint: Sprint = Sprint(sprint_id=sprint_id, name=name,
//                                 start_date=start_date, end_date=end_date, updated_by=updated_by, created_at=None, updated_at=None)
//         self.sprint_repository.update_sprint(sprint=sprint)

type ISprintService interface {
	SprintCreate(name string, startDate time.Time, endDate time.Time, updatedBy string) (err error)
	SprintList() ([]*models.Sprint, error)
	SprintUpdate(sprintId uuid.UUID, name string, startDate time.Time, endDate time.Time, updatedBy string) (err error)
	SprintDelete(sprintId uuid.UUID) (err error)
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

func (us *sprintService) SprintList() ([]*models.Sprint, error) {
	return us.ur.SprintList()
}

func (us *sprintService) SprintUpdate(sprintId uuid.UUID, name string, startDate time.Time, endDate time.Time, updatedBy string) (err error) {
	sprint := &models.Sprint{
		ID:        sprintId,
		Name:      &name,
		StartDate: startDate,
		EndDate:   endDate,
		UpdatedBy: &updatedBy,
	}

	if err := us.ur.SprintUpdate(sprint); err != nil {
		return err
	}
	return nil
}

func (us *sprintService) SprintDelete(sprintId uuid.UUID) (err error) {
	if err := us.ur.SprintDelete(sprintId); err != nil {
		return err
	}
	return nil
}
