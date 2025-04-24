package models

import "gorm.io/gorm"

type Lead struct {
	gorm.Model
	Name    string `gorm:"varchar(50);not null" json:"name"`
	Company string `gorm:"varchar(50);not null" json:"company"`
	Phone   string `gorm:"varchar(15);not null; unique" json:"phone"`
	Email   string `gorm:"varchar(100);not null; unique" json:"email"`
}
