package worker

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	producer "github.com/fiveret/crm-golang/internal/kafka"
	"github.com/fiveret/crm-golang/internal/repository"
)

func (w worker) StartWorker(ctx context.Context) {
	w.logger.Info("Worker started")
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ctx.Done():
			w.logger.Info("Worker stopping")
			return
		case <-ticker.C:
			events, err := w.repo.GetEvents()
			if err != nil {
				w.logger.Error(fmt.Sprintf("error getting events: %v", err))
				return
			}
			for _, event := range events {
				for i := 1; i <= event.RetryCount; i++ {
					err := w.publisher.Publish(ctx, event)
					if err != nil {
						if i == event.RetryCount {
							event.Status = "failed"
							err = w.repo.UpdateEvent(event)
							if err != nil {
								w.logger.Error(fmt.Sprintf("error updating event: %v", err))
								continue
							}
						}
						continue
					}
					event.Status = "sent"
					err = w.repo.UpdateEvent(event)
					if err != nil {
						w.logger.Error(fmt.Sprintf("error updating event: %v", err))
						continue
					}
				}
			}
		}
	}
}

type Worker interface {
	StartWorker(ctx context.Context)
}

type worker struct {
	publisher producer.EventPublisher
	logger    *slog.Logger
	repo      repository.EventRepo
}

func NewWorker(repo repository.EventRepo, logger *slog.Logger, p producer.EventPublisher) Worker {
	return worker{logger: logger, repo: repo, publisher: p}
}
