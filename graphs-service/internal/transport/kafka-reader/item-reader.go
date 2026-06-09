package reader

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	domain "graphs-service/internal/entities"
	dto "graphs-service/internal/transport/dto"
	models "graphs-service/internal/transport/kafka-reader/models"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

type itemReader struct {
	reader *kafka.Reader
	logger *slog.Logger
}

type ItemReader interface {
	Start(ctx context.Context, ch chan<- models.Envelope)
	Reader() *kafka.Reader
}

func (i *itemReader) Reader() *kafka.Reader {
	return i.reader
}

func (i *itemReader) Start(ctx context.Context, ch chan<- models.Envelope) {
	for {
		msg, err := i.reader.FetchMessage(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return
			}
			i.logger.Error("fetch error", "error", err)
			continue
		}
		fmt.Println("GOT MESSAGE: ", string(msg.Value))
		event, err := parse(msg)
		if err != nil {
			i.logger.Error("parse error", "error", err)
			continue
		}
		fmt.Println("Created new event: ", event.EventType)
		select {
		case <-ctx.Done():
			return
		case ch <- models.Envelope{
			Event: *event,
			Msg:   msg,
		}:
		}
	}
}
func parse(msg kafka.Message) (*domain.Event, error) {
	e := new(dto.Event)

	if err := json.Unmarshal(msg.Value, e); err != nil {
		return nil, err
	}

	return &domain.Event{
		EventType: domain.EventType(e.EventType),
		EventID:   e.EventID,
		Payload:   e.Payload,
	}, nil
}
