package entity

import (
	"gorm.io/gorm"
)

type Auditorium struct {
	gorm.Model
	Name string `json:"name" validate:"required" gorm:"type:varchar(255);not null"`
}
