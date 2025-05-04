package models

import "gorm.io/gorm"

type Item struct {
	Name        *string  `gorm:"unique; not null" json:"name"`
	Description *string  `gorm:"not null" json:"description"`
	Price       *float64 `gorm:"not null" json:"price"`
	Category    *string  `gorm:"not null" json:"category"`
	InStock     *uint32  `gorm:"not null" json:"in_stock"`
	Currency    *string  `gorm:"default:KZT" json:"currency"`
	Status      *string  `gorm:"default:in stock" json:"status"`
	gorm.Model
}
