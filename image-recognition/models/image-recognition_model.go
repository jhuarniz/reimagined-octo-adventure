package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ImageRecognition struct {
	gorm.Model

	ID       uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4()"`
	Email    string    `gorm:"type:varchar(150);not null"`
	Filters  string    `gorm:"type:varchar(100);not null"`
	ImageUrl string    `gorm:"type:varchar(200);not null"`
}
