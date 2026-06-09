package models

import (
	domain "graphs-service/internal/entities"

	"github.com/segmentio/kafka-go"
)

type Envelope struct {
	Event domain.Event
	Msg   kafka.Message
}
