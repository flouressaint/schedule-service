package repo

import (
	"log"

	"github.com/flouressaint/schedule-service/internal/entity"
	"gorm.io/gorm"
)

type AuditoriumRepo struct {
	db *gorm.DB
}

func NewAuditoriumRepo(db *gorm.DB) *AuditoriumRepo {
	// migrate the database
	err := db.AutoMigrate(&entity.Auditorium{})
	if err != nil {
		log.Fatal("Failed to migrate auditorium in the Database")
	}
	log.Println("Migration auditorium complete")
	return &AuditoriumRepo{db}
}

func (r *AuditoriumRepo) CreateAuditorium(auditorium entity.Auditorium) (uint, error) {
	if err := r.db.Create(&auditorium).Error; err != nil {
		return auditorium.ID, err
	}
	return auditorium.ID, nil
}

func (r *AuditoriumRepo) GetAuditoriums() ([]entity.Auditorium, error) {
	var auditoriums []entity.Auditorium
	if err := r.db.Find(&auditoriums).Error; err != nil {
		return auditoriums, err
	}
	return auditoriums, nil
}

func (r *AuditoriumRepo) DeleteAuditorium(id uint) error {
	var auditorium entity.Auditorium
	if err := r.db.First(&auditorium, id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&auditorium).Error; err != nil {
		return err
	}
	return nil
}

func (r *AuditoriumRepo) UpdateAuditorium(id uint, newauditorium entity.Auditorium) error {
	var auditorium entity.Auditorium
	if err := r.db.First(&auditorium, id).Error; err != nil {
		return err
	}
	if err := r.db.Model(&auditorium).Updates(newauditorium).Error; err != nil {
		return err
	}
	return nil
}
