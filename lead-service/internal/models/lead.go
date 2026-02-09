package models

import (
	"time"
)

type Lead struct {
	ID               uint       `gorm:"primaryKey"`
	Name             string     `gorm:"varchar(50);not null"`
	Email            string     `gorm:"varchar(100);not null;unique"`
	Phone            string     `gorm:"varchar(15);not null;unique"`
	Company          string     `gorm:"varchar(50);not null"`
	Products         []*Product `gorm:"many2many:lead_products;constraint:OnDelete:CASCADE"`
	Visits           uint       `gorm:"default:0"`
	LastVisit        *time.Time
	TotalSales       float64 `gorm:"default:0"`
	LastPurchaseDays *uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Category    string  `gorm:"not null"`
	Currency    string  `gorm:"default:'KZT'"`
	InStock     uint    `gorm:"not null"`
	Status      string  `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
