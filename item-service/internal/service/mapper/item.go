package mapper

import (
	"time"

	"github.com/fiveret/item-service/internal/repository/models"
	serviceModels "github.com/fiveret/item-service/internal/service/models"
	"github.com/google/uuid"
)

func OutboxEventToEvent(event models.OutboxEvent) *serviceModels.Event {
	return &serviceModels.Event{
		EventID:    uuid.NewString(),
		EventType:  event.EventType,
		OccurredAt: time.Now(),
		Payload:    event.Payload,
	}
}
