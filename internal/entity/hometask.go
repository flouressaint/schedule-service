package entity

import (
	"gorm.io/gorm"
)

type Hometask struct {
	gorm.Model
	Description string `json:"description" validate:"required" gorm:"type:varchar(255);not null"`
	Attachment  string `json:"attachment" validate:"required" gorm:"type:varchar(255);not null"`
}
