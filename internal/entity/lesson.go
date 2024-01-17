package entity

import (
	"time"

	"github.com/google/uuid"
)

type Lesson struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	Date     time.Time `json:"date" validate:"required" gorm:"type:timestamp"`
	Duration int       `json:"duration" validate:"required" gorm:"type:int;not null"`

	DisciplineID  uint `json:"discipline_id" validate:"required" gorm:"type:int;not null"`
	Discipline    Discipline
	TeacherUserID uuid.UUID `json:"teacher_user_id" validate:"required" gorm:"type:uuid"`
	StudyGroupID  uint      `json:"study_group_id" validate:"required" gorm:"type:int;not null"`
	AuditoriumID  uint      `json:"auditorium_id" validate:"required" gorm:"type:int;not null"`
	Auditorium    Auditorium
	HometaskID    uint `json:"hometask_id" validate:"required" gorm:"type:int;not null"`
	Hometask      Hometask
}
