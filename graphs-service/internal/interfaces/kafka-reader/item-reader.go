package reader

import (
	"context"
	"encoding/json"
	"errors"
	domain "graphs-service/internal/entities"
	models "graphs-service/internal/interfaces/dto"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

type itemReader struct {
	reader *kafka.Reader
	logger *slog.Logger
}

type ItemReader interface {
	Start(ctx context.Context, itemChan chan<- domain.Event)
}

func (i *itemReader) Start(ctx context.Context, itemChan chan<- domain.Event) {
	for {
		msg, err := i.reader.FetchMessage(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return
			}
			i.logger.Error("kafka item fetch error", "error", err)
			continue
		}
		eventDomain, err := parse(msg)
		if err != nil {
			i.logger.Error("item parse error", "error", err)
			continue
		}
		select {
		case <-ctx.Done():
			return
		case itemChan <- *eventDomain:
		}
		if err := i.reader.CommitMessages(ctx, msg); err != nil {
			i.logger.Error("commit error", "error", err)
		}
	}
}

func parse(msg kafka.Message) (*domain.Event, error) {
	eventModels := new(models.Event)
	err := json.Unmarshal(msg.Value, eventModels)
	if err != nil {
		return nil, err
	}
	return &domain.Event{
		EventType: domain.EventType(eventModels.EventType),
		EventID:   eventModels.EventID,
		Payload:   eventModels.Payload,
	}, nil
}
