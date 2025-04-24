package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        *string  `gorm:"unique; not null"`
	Description *string  `gorm:"not null"`
	Price       *float64 `gorm:"not null"`
	Category    *string  `gorm:"not null"`
	Currency    *string  `gorm:"default:'KZT'"`
	InStock     *uint    `gorm:"not null"`
	Status      *string  `gorm:"not null"`
}
