package mapper

import (
	"time"

	"github.com/fiveret/item-service/internal/domain"
	"github.com/fiveret/item-service/internal/repository/models"
	serviceModels "github.com/fiveret/item-service/internal/service/models"
)

func OutboxEventToEvent(event *models.OutboxEvent) *serviceModels.Event {
	return &serviceModels.Event{
		EventID:    event.EventID,
		EventType:  event.EventType,
		OccurredAt: time.Now(),
		Payload:    event.Payload,
	}
}

func EventToOutboxEvent(event *serviceModels.Event, item *domain.Item, topic string) *models.OutboxEvent {
	return &models.OutboxEvent{
		EventID:       event.EventID,
		AggregateID:   item.ID,
		AggregateType: topic,
		EventType:     event.EventType,
		Payload:       event.Payload,
		ProcessedAt:   event.OccurredAt,
	}
}
