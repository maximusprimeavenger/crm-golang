package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fiveret/crm-golang/internal/models"
	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
)

type EventPublisher interface {
	Publish(ctx context.Context, event *models.OutboxEvent) error
}

type kafkaPublisher struct {
	writer *kafka.Writer
}

func NewKafkaPublisher(brokers []string) EventPublisher {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
	})
	return &kafkaPublisher{writer: writer}
}

func (p *kafkaPublisher) Publish(ctx context.Context, event *models.OutboxEvent) error {
	payloadBytes, err := json.Marshal(event.Payload)
	if err != nil {
		return fmt.Errorf("couldn't marshal payload: %v", err)
	}
	eventKafka := &models.Event{
		EventID:    uuid.NewString(),
		EventType:  event.EventType,
		OccurredAt: event.CreatedAt.UTC(),
		Payload:    payloadBytes,
	}
	eventBytes, err := json.Marshal(eventKafka)
	if err != nil {
		return err
	}
	keyString := strconv.Itoa(int(event.AggregateID))
	msg := kafka.Message{
		Topic: event.AggregateType,
		Value: eventBytes,
		Key:   []byte(keyString),
	}
	return p.writer.WriteMessages(ctx, msg)
}
