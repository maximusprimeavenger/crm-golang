package worker

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	producer "github.com/fiveret/crm-golang/internal/kafka"
	"github.com/fiveret/crm-golang/internal/models"
	"github.com/fiveret/crm-golang/internal/repository"
)

func (w worker) StartWorker(ctx context.Context) {
	w.logger.Info("Worker started")
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			w.logger.Info("Worker stopping")
			return
		case <-ticker.C:
			events, err := w.repo.GetEvents()
			if err != nil {
				w.logger.Error(fmt.Sprintf("error getting events: %v", err))
				continue
			}
			var wg sync.WaitGroup
			for _, event := range events {
				e := event
				if e.Status == "sent" {
					continue
				}
				wg.Add(1)
				go func(ev *models.OutboxEvent) {
					defer wg.Done()
					w.mu.Lock()
					ev.Status = "processing"
					if err := w.repo.UpdateEvent(ev); err != nil {
						w.logger.Error(fmt.Sprintf("error updating event: %v", err))
						w.mu.Unlock()
						return
					}
					w.mu.Unlock()
					retries := ev.RetryCount
					if retries <= 0 {
						retries = 1
					}
					for i := 1; i <= retries; i++ {
						select {
						case <-ctx.Done():
							return
						default:
						}

						err := w.publisher.Publish(ctx, ev)
						if err != nil {
							if i == retries {
								w.mu.Lock()
								ev.Status = "failed"
								if err2 := w.repo.UpdateEvent(ev); err2 != nil {
									w.logger.Error(fmt.Sprintf("error updating event: %v", err2))
								}
								w.mu.Unlock()
							} else {
								time.Sleep(time.Second * time.Duration(i))
							}
							continue
						}

						w.mu.Lock()
						ev.Status = "sent"
						ev.ProcessedAt = time.Now()
						if err := w.repo.UpdateEvent(ev); err != nil {
							w.logger.Error(fmt.Sprintf("error updating event: %v", err))
						}
						w.mu.Unlock()
						break
					}
				}(e)
			}
			wg.Wait()
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
	mu        sync.Mutex
}

func NewWorker(repo repository.EventRepo, logger *slog.Logger, p producer.EventPublisher) Worker {
	return worker{logger: logger, repo: repo, publisher: p}
}
