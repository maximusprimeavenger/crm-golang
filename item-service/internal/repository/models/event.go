package models

import (
	"encoding/json"
	"time"
)

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
