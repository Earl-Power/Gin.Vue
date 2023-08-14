package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        int
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(110);not null"`
	Password  string `gorm:"size:255;not null"`
}
