package repo

import (
	"log"

	"github.com/flouressaint/schedule-service/internal/entity"
	"gorm.io/gorm"
)

type HometaskRepo struct {
	db *gorm.DB
}

func NewHometaskRepo(db *gorm.DB) *HometaskRepo {
	// migrate the database
	err := db.AutoMigrate(&entity.Hometask{})
	if err != nil {
		log.Fatal("Failed to migrate hometask in the Database")
	}
	log.Println("Migration hometask complete")
	return &HometaskRepo{db}
}

func (r *HometaskRepo) CreateHometask(hometask entity.Hometask) (uint, error) {
	if err := r.db.Create(&hometask).Error; err != nil {
		return hometask.ID, err
	}
	return hometask.ID, nil
}

func (r *HometaskRepo) GetHometasks() ([]entity.Hometask, error) {
	var hometasks []entity.Hometask
	if err := r.db.Find(&hometasks).Error; err != nil {
		return hometasks, err
	}
	return hometasks, nil
}

func (r *HometaskRepo) DeleteHometask(id uint) error {
	var hometask entity.Hometask
	if err := r.db.First(&hometask, id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&hometask).Error; err != nil {
		return err
	}
	return nil
}

func (r *HometaskRepo) UpdateHometask(id uint, newHometask entity.Hometask) error {
	var hometask entity.Hometask
	if err := r.db.First(&hometask, id).Error; err != nil {
		return err
	}
	if err := r.db.Model(&hometask).Updates(newHometask).Error; err != nil {
		return err
	}
	return nil
}
