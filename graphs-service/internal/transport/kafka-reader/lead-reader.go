package reader

import (
	"context"
	"errors"
	domain "graphs-service/internal/entities"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

type leadReader struct {
	reader *kafka.Reader
	logger *slog.Logger
}
type leadProductReader struct {
	reader *kafka.Reader
	logger *slog.Logger
}

type LeadReader interface {
	Start(ctx context.Context, leadChan chan<- domain.Event)
}
type LeadProductReader interface {
	Start(ctx context.Context, leadProductChan chan<- domain.Event)
}

func LoadReaders(itemTopic string, leadTopics, address []string, l *slog.Logger) (ItemReader, LeadReader, LeadProductReader, error) {
	itemR := kafka.NewReader(kafka.ReaderConfig{
		Brokers: address,
		Topic:   itemTopic,
	})
	leadR := kafka.NewReader(kafka.ReaderConfig{
		Brokers: address,
		Topic:   leadTopics[0],
	})
	leadProductR := kafka.NewReader(kafka.ReaderConfig{
		Brokers: address,
		Topic:   leadTopics[1],
	})
	return &itemReader{reader: itemR, logger: l}, &leadReader{reader: leadR}, &leadProductReader{reader: leadProductR}, nil
}

func (l *leadReader) Start(ctx context.Context, leadChan chan<- domain.Event) {
	for {
		msg, err := l.reader.FetchMessage(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return
			}
			l.logger.Error("kafka item fetch error", "error", err)
			continue
		}
		parsed, err := parse(msg)
		if err != nil {
			l.logger.Error("lead parse error", "error", err)
			continue
		}
		select {
		case <-ctx.Done():
			return
		case leadChan <- *parsed:
		}
		if err := l.reader.CommitMessages(ctx); err != nil {
			l.logger.Error("commit error", "error", err)
		}
	}
}

func (l *leadProductReader) Start(ctx context.Context, leadProductChan chan<- domain.Event) {
	for {
		msg, err := l.reader.FetchMessage(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return
			}
			l.logger.Error("kafka item fetch error", "error", err)
			continue
		}
		parsed, err := parse(msg)
		if err != nil {
			l.logger.Error("lead parse error", "error", err)
			continue
		}
		select {
		case <-ctx.Done():
			return
		case leadProductChan <- *parsed:
		}
		if err := l.reader.CommitMessages(ctx); err != nil {
			l.logger.Error("commit error", "error", err)
		}
	}
}
