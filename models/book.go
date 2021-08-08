package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	MediumPrice float32        `json:"medium_price"`
	Author      string         `json:"author"`
	ImageURl    string         `json:"img_url"`
	CreatedAt   time.Time      `json:"created"`
	UpdateAt    time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"indec" json:"deleted"`
}
