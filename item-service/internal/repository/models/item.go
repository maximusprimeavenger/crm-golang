package models

import "time"

type Item struct {
	ID          uint      `gorm:"primarykey" json:"item_id"`
	Name        *string   `gorm:"unique; not null" json:"name"`
	Description *string   `gorm:"not null" json:"description"`
	OldPrice    *float64  `json:"old_price"`
	NewPrice    *float64  `gorm:"not null" json:"new_price"`
	Category    *string   `gorm:"not null" json:"category"`
	InStock     *uint32   `gorm:"not null" json:"in_stock"`
	Currency    *string   `gorm:"default:KZT" json:"currency"`
	Status      *string   `gorm:"default:in stock" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
