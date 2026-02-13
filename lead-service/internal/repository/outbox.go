package repository

import (
	"github.com/fiveret/crm-golang/internal/models"
	"gorm.io/gorm"
)

type EventRepo interface {
	GetEvents() ([]*models.OutboxEvent, error)
	UpdateEvent(event *models.OutboxEvent) error
}

type eventRepo struct {
	db *gorm.DB
}

func NewEventRepo(dbConn *gorm.DB) EventRepo {
	return eventRepo{db: dbConn}
}

func (r eventRepo) GetEvents() ([]*models.OutboxEvent, error) {
	events := new([]*models.OutboxEvent)
	if err := r.db.Find(&events).Where("status = ?", "pending").Error; err != nil {
		return nil, err
	}
	return *events, nil
}

func (r eventRepo) UpdateEvent(event *models.OutboxEvent) error {
	if err := r.db.Save(event).Error; err != nil {
		return err
	}
	return nil
}
