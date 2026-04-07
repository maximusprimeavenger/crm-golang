package aggregator

import (
	"context"
	domain "graphs-service/internal/entities"
	"log/slog"
	"sync"
)

type aggregator struct {
	mu       sync.Mutex
	logger   *slog.Logger
	Items    map[uint]*domain.ItemAnalytics
	Leads    map[uint]*domain.LeadState
	handlers map[domain.EventType]EventHandler
}

func NewAggregator(log *slog.Logger) *aggregator {
	a := &aggregator{Leads: make(map[uint]*domain.LeadState),
		Items:    make(map[uint]*domain.ItemAnalytics),
		logger:   log,
		handlers: make(map[domain.EventType]EventHandler),
	}
	a.handlers[domain.ItemCreated] = a.eventItemCreated
	return a
}

func (a *aggregator) Run(ctx context.Context, eventChan chan domain.Event) {
	for {
		select {
		case <-ctx.Done():
			return
		case event := <-eventChan:
			a.apply(event)
		}
	}
}

type EventHandler func(e domain.Event)

func (a *aggregator) apply(event domain.Event) {
	handlers, ok := a.handlers[event.EventType]
	if !ok {
		a.logger.Warn("unknown event type", "type", event.EventType)
		return
	}
	handlers(event)
}

func (a *aggregator) GetItemAnalytics() map[uint]*domain.ItemAnalytics {
	a.mu.Lock()
	defer a.mu.Unlock()
	copyMap := make(map[uint]*domain.ItemAnalytics)
	for k, v := range a.Items {
		copyMap[k] = v
	}
	return copyMap
}

func (a *aggregator) GetSalesAndRevenueByID(itemID uint) (map[string]int, map[string]float64, string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	analytics := a.Items[itemID]
	return analytics.SalesByDay, analytics.RevenueByDay, a.Items[itemID].Name
}
