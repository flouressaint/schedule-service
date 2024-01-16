package repo

import (
	"log"

	"github.com/flouressaint/schedule-service/internal/entity"
	"gorm.io/gorm"
)

type DisciplineRepo struct {
	db *gorm.DB
}

func NewDisciplineRepo(db *gorm.DB) *DisciplineRepo {
	// migrate the database
	err := db.AutoMigrate(&entity.Discipline{})
	if err != nil {
		log.Fatal("Failed to migrate discipline in the Database")
	}
	log.Println("Migration discipline complete")
	return &DisciplineRepo{db}
}

func (r *DisciplineRepo) CreateDiscipline(discipline entity.Discipline) (uint, error) {
	if err := r.db.Create(&discipline).Error; err != nil {
		return discipline.ID, err
	}
	return discipline.ID, nil
}

func (r *DisciplineRepo) GetDisciplines() ([]entity.Discipline, error) {
	var disciplines []entity.Discipline
	if err := r.db.Find(&disciplines).Error; err != nil {
		return disciplines, err
	}
	return disciplines, nil
}

func (r *DisciplineRepo) DeleteDiscipline(id uint) error {
	var discipline entity.Discipline
	if err := r.db.First(&discipline, id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&discipline).Error; err != nil {
		return err
	}
	return nil
}

func (r *DisciplineRepo) UpdateDiscipline(id uint, newDiscipline entity.Discipline) error {
	var discipline entity.Discipline
	if err := r.db.First(&discipline, id).Error; err != nil {
		return err
	}
	if err := r.db.Model(&discipline).Updates(newDiscipline).Error; err != nil {
		return err
	}
	return nil
}
