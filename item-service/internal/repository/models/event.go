package models

import (
	"encoding/json"
	"time"
)

type OutboxEvent struct {
	ID            uint            `gorm:"primaryKey;autoIncrement"`   //ID события в базе данных
	EventID       string          `gorm:"not null"`                   //ID события
	AggregateID   uint            `gorm:"not null"`                   // ID продукта
	AggregateType string          `gorm:"not null"`                   // "item"
	EventType     string          `gorm:"not null"`                   // "item.created"
	Payload       json.RawMessage `gorm:"type:jsonb;not null"`        //json event
	Status        string          `gorm:"not null;default:'pending'"` // pending / processing / sent / failed
	RetryCount    int             `gorm:"default:0"`
	CreatedAt     time.Time       `gorm:"autoCreateTime"`
	ProcessedAt   time.Time
}
