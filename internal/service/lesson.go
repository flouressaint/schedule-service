package service

import (
	"github.com/flouressaint/schedule-service/internal/entity"
	"github.com/flouressaint/schedule-service/internal/repo"
)

type LessonService struct {
	lessonRepo repo.Lesson
}

func NewLessonService(lessonRepo repo.Lesson) *LessonService {
	return &LessonService{
		lessonRepo: lessonRepo,
	}
}

func (s *LessonService) CreateLesson(lesson entity.Lesson) (uint, error) {
	return s.lessonRepo.CreateLesson(lesson)
}

func (s *LessonService) GetLessons() ([]entity.Lesson, error) {
	return s.lessonRepo.GetLessons()
}

func (s *LessonService) DeleteLesson(id uint) error {
	return s.lessonRepo.DeleteLesson(id)
}

func (s *LessonService) UpdateLesson(id uint, newlesson entity.Lesson) error {
	return s.lessonRepo.UpdateLesson(id, newlesson)
}
