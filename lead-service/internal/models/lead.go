package models

import "gorm.io/gorm"

type Lead struct {
	Name     string    `gorm:"varchar(50);not null" json:"name"`
	Email    string    `gorm:"varchar(100);not null; unique" json:"email"`
	Phone    string    `gorm:"varchar(15);not null; unique" json:"phone"`
	Company  string    `gorm:"varchar(50);not null" json:"company"`
	Products []Product `gorm:"foreignKey:LeadID"`
	LeadID   uint
	gorm.Model
}

type Product struct {
	Name        *string  `gorm:"unique; not null" json:"name"`
	Description *string  `gorm:"not null" json:"description"`
	Price       *float64 `gorm:"not null" json:"price"`
	Category    *string  `gorm:"not null" json:"category"`
	Currency    *string  `gorm:"default:'KZT'" json:"currency"`
	InStock     *uint    `gorm:"not null" json:"in_stock"`
	Status      *string  `gorm:"not null" json:"status"`
	gorm.Model
}
