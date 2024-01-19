package repo

import (
	"log"

	"github.com/flouressaint/schedule-service/internal/entity"
	"gorm.io/gorm"
)

type LessonRepo struct {
	db *gorm.DB
}

func NewLessonRepo(db *gorm.DB) *LessonRepo {
	// migrate the database
	err := db.AutoMigrate(&entity.Lesson{})
	if err != nil {
		log.Fatal("Failed to migrate lesson in the Database")
	}
	log.Println("Migration lesson complete")
	return &LessonRepo{db}
}

func (r *LessonRepo) CreateLesson(lesson entity.Lesson) (uint, error) {
	if err := r.db.Preload("Discipline").Preload("Hometask").Preload("Auditorium").Create(&lesson).Error; err != nil {
		return lesson.ID, err
	}
	return lesson.ID, nil
}

func (r *LessonRepo) GetLessons() ([]entity.Lesson, error) {
	var lessons []entity.Lesson
	if err := r.db.Preload("Discipline").Preload("Hometask").Preload("Auditorium").Find(&lessons).Error; err != nil {
		return lessons, err
	}
	return lessons, nil
}

func (r *LessonRepo) DeleteLesson(id uint) error {
	var lesson entity.Lesson
	if err := r.db.First(&lesson, id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&lesson).Error; err != nil {
		return err
	}
	return nil
}

func (r *LessonRepo) UpdateLesson(id uint, newLesson entity.Lesson) error {
	var lesson entity.Lesson
	if err := r.db.First(&lesson, id).Error; err != nil {
		return err
	}
	if err := r.db.Model(&lesson).Updates(newLesson).Error; err != nil {
		return err
	}
	return nil
}
