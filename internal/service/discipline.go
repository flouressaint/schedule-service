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

func (s *DisciplineService) Creatediscipline(discipline entity.Discipline) (uint, error) {
	return s.disciplineRepo.Creatediscipline(discipline)
}

func (s *DisciplineService) Getdisciplines() ([]entity.Discipline, error) {
	return s.disciplineRepo.Getdisciplines()
}

func (s *DisciplineService) Deletediscipline(id uint) error {
	return s.disciplineRepo.Deletediscipline(id)
}

func (s *DisciplineService) Updatediscipline(id uint, newdiscipline entity.Discipline) error {
	return s.disciplineRepo.Updatediscipline(id, newdiscipline)
}
