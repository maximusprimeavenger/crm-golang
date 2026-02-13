package models

import (
	"encoding/json"
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
	EventType     string          `gorm:"not null"`                   // "LeadCreated"
	Payload       json.RawMessage `gorm:"type:jsonb;not null"`        //json event
	Status        string          `gorm:"not null;default:'pending'"` // pending / processing / sent / failed
	RetryCount    int             `gorm:"default:0"`
	CreatedAt     time.Time       `gorm:"autoCreateTime"`
	ProcessedAt   time.Time
}
