package service

import (
	"github.com/flouressaint/schedule-service/internal/entity"
	"github.com/flouressaint/schedule-service/internal/repo"
)

type DisciplineService struct {
	disciplineRepo repo.Discipline
}

func NewDisciplineService(disciplineRepo repo.Discipline) *DisciplineService {
	return &DisciplineService{
		disciplineRepo: disciplineRepo,
	}
}

func (s *DisciplineService) CreateDiscipline(discipline entity.Discipline) (uint, error) {
	return s.disciplineRepo.CreateDiscipline(discipline)
}

func (s *DisciplineService) GetDisciplines() ([]entity.Discipline, error) {
	return s.disciplineRepo.GetDisciplines()
}

func (s *DisciplineService) DeleteDiscipline(id uint) error {
	return s.disciplineRepo.DeleteDiscipline(id)
}

func (s *DisciplineService) UpdateDiscipline(id uint, newdiscipline entity.Discipline) error {
	return s.disciplineRepo.UpdateDiscipline(id, newdiscipline)
}
