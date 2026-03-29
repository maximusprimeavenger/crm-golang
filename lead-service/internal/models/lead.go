package models

import (
	"encoding/json"
	"time"
)

type Lead struct {
	ID               uint       `gorm:"primaryKey" json:"lead_id"`
	Name             string     `gorm:"varchar(50);not null" json:"name"`
	Email            string     `gorm:"varchar(100);not null;unique" json:"email"`
	Phone            string     `gorm:"varchar(15);not null;unique" json:"phone"`
	Company          string     `gorm:"varchar(50);not null" json:"company"`
	Products         []*Product `gorm:"many2many:lead_products;constraint:OnDelete:CASCADE" json:"products"`
	Visits           uint       `gorm:"default:0" json:"visits"`
	LastVisit        *time.Time `json:"last_visit"`
	TotalSales       float64    `gorm:"default:0" json:"total_sales"`
	LastPurchaseDays *uint      `json:"last_purchase_days"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"product_id"`
	Name        string    `gorm:"not null" json:"product_name"`
	Description string    `gorm:"not null" json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
	Category    string    `gorm:"not null" json:"category"`
	Currency    string    `gorm:"default:'KZT'" json:"currency"`
	InStock     uint      `gorm:"not null" json:"in_stock"`
	Status      string    `gorm:"not null" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Event struct {
	EventID    string          `json:"event_id"`
	EventType  string          `json:"event_type"`
	OccurredAt time.Time       `json:"occurred_at"`
	Payload    json.RawMessage `json:"payload"`
}

type OutboxEvent struct {
	ID            uint            `gorm:"primaryKey;autoIncrement"`
	AggregateID   uint            `gorm:"not null"`                   // ID лида
	AggregateType string          `gorm:"not null"`                   // "lead"
	EventType     string          `gorm:"not null"`                   // "lead.created"
	Payload       json.RawMessage `gorm:"type:jsonb;not null"`        //json event
	Status        string          `gorm:"not null;default:'pending'"` // pending / processing / sent / failed
	RetryCount    int             `gorm:"default:0"`
	CreatedAt     time.Time       `gorm:"autoCreateTime"`
	ProcessedAt   time.Time
}
