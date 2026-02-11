package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/fiveret/crm-golang/internal/models"
	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
)

type EventPublisher interface {
	Publish(ctx context.Context, topic, key, typeEvent string, createdAt *time.Time, payload any) error
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

func (p *kafkaPublisher) Publish(ctx context.Context, topic, key, typeEvent string, createdAt *time.Time, payload any) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("couldn't marshal payload: %v", err)
	}
	event := &models.Event{
		EventID:    uuid.NewString(),
		EventType:  typeEvent,
		OccurredAt: createdAt.UTC(),
		Payload:    payloadBytes,
	}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}
	msg := kafka.Message{
		Topic: topic,
		Value: eventBytes,
		Key:   []byte(key),
	}
	return p.writer.WriteMessages(ctx, msg)
}
