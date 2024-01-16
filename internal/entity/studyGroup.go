package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudyGroup struct {
	ID   uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name" validate:"required" gorm:"type:varchar(255);not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
