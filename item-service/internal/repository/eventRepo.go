package repository

import (
	"github.com/fiveret/item-service/internal/repository/models"
	"gorm.io/gorm"
)

type EventRepo interface {
	GetEvents() ([]*models.OutboxEvent, error)
	UpdateEvent(*models.OutboxEvent) error
}

type eventRepo struct {
	db *gorm.DB
}

func NewEventRepo(db *gorm.DB) EventRepo {
	return &eventRepo{db: db}
}

func (e *eventRepo) GetEvents() ([]*models.OutboxEvent, error) {
	events := []*models.OutboxEvent{}
	if err := e.db.Find(&events).Where("status = ?", "pending").Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (e *eventRepo) UpdateEvent(event *models.OutboxEvent) error {
	if err := e.db.Save(event).Error; err != nil {
		return err
	}
	return nil
}
