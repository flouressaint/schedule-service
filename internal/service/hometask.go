package service

import (
	"github.com/flouressaint/schedule-service/internal/entity"
	"github.com/flouressaint/schedule-service/internal/repo"
)

type HometaskService struct {
	hometaskRepo repo.Hometask
}

func NewHometaskService(hometaskRepo repo.Hometask) *HometaskService {
	return &HometaskService{
		hometaskRepo: hometaskRepo,
	}
}

func (s *HometaskService) CreateHometask(hometask entity.Hometask) (uint, error) {
	return s.hometaskRepo.CreateHometask(hometask)
}

func (s *HometaskService) GetHometasks() ([]entity.Hometask, error) {
	return s.hometaskRepo.GetHometasks()
}

func (s *HometaskService) DeleteHometask(id uint) error {
	return s.hometaskRepo.DeleteHometask(id)
}

func (s *HometaskService) UpdateHometask(id uint, newhometask entity.Hometask) error {
	return s.hometaskRepo.UpdateHometask(id, newhometask)
}
