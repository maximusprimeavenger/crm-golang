package aggregator

import (
	"errors"
	"fmt"
	domain "graphs-service/internal/entities"
	"log/slog"
	"sync"
)

type Aggregator struct {
	mu       sync.RWMutex
	logger   *slog.Logger
	Items    map[uint]*domain.ItemAnalytics
	Leads    map[uint]*domain.LeadState
	handlers map[domain.EventType]EventHandler
}

func NewAggregator(log *slog.Logger) *Aggregator {
	a := &Aggregator{Leads: make(map[uint]*domain.LeadState),
		Items:    make(map[uint]*domain.ItemAnalytics),
		logger:   log,
		handlers: make(map[domain.EventType]EventHandler),
	}
	a.handlers[domain.ItemCreated] = a.eventItemCreated
	return a
}

type EventHandler func(e domain.Event)

func (a *Aggregator) Apply(event domain.Event) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	fmt.Println("Aggregator received the event: ", event.EventType)
	handler, ok := a.handlers[event.EventType]
	if !ok {
		return errors.New("unknown event type")
	}

	handler(event)
	return nil
}
func (a *Aggregator) GetItemAnalytics() map[uint]*domain.ItemAnalytics {
	a.mu.Lock()
	defer a.mu.Unlock()
	copyMap := make(map[uint]*domain.ItemAnalytics)
	for k, v := range a.Items {
		copyMap[k] = v
	}
	return copyMap
}

func (a *Aggregator) GetSalesAndRevenueByID(itemID uint) (map[string]int, map[string]float64, string) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	item, ok := a.Items[itemID]
	if !ok || item == nil {
		return map[string]int{}, map[string]float64{}, ""
	}

	return item.SalesByDay, item.RevenueByDay, item.Name
}

func (a *Aggregator) Ready(itemID uint) bool {
	a.mu.RLock()
	defer a.mu.RUnlock()

	item, ok := a.Items[itemID]
	if !ok || item == nil {
		return false
	}

	return len(item.SalesByDay) > 0 || len(item.RevenueByDay) > 0
}
