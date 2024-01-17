package entity

type Auditorium struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" validate:"required" gorm:"unique;type:varchar(255);not null"`
}
