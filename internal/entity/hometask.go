package entity

type Hometask struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Description string `json:"description" validate:"required" gorm:"type:varchar(255);not null"`
	Attachment  string `json:"attachment" gorm:"type:varchar(255)"`
}
