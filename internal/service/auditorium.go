package service

import (
	"github.com/flouressaint/schedule-service/internal/entity"
	"github.com/flouressaint/schedule-service/internal/repo"
)

type AuditoriumService struct {
	auditoriumRepo repo.Auditorium
}

func NewAuditoriumService(auditoriumRepo repo.Auditorium) *AuditoriumService {
	return &AuditoriumService{
		auditoriumRepo: auditoriumRepo,
	}
}

func (s *AuditoriumService) CreateAuditorium(auditorium entity.Auditorium) (uint, error) {
	return s.auditoriumRepo.CreateAuditorium(auditorium)
}

func (s *AuditoriumService) GetAuditoriums() ([]entity.Auditorium, error) {
	return s.auditoriumRepo.GetAuditoriums()
}

func (s *AuditoriumService) DeleteAuditorium(id uint) error {
	return s.auditoriumRepo.DeleteAuditorium(id)
}

func (s *AuditoriumService) UpdateAuditorium(id uint, newAuditorium entity.Auditorium) error {
	return s.auditoriumRepo.UpdateAuditorium(id, newAuditorium)
}
