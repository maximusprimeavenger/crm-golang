package aggregator

import (
	"context"
	"fmt"
	"graphs-service/internal/domain"
	"graphs-service/internal/interfaces/mapper"
	"log/slog"
)

type aggregator struct {
	logger *slog.Logger
	Items  map[uint]*domain.ItemState
	Leads  map[uint]*domain.LeadState
}

func NewAggregator(log *slog.Logger) *aggregator {
	return &aggregator{Leads: make(map[uint]*domain.LeadState), Items: make(map[uint]*domain.ItemState), logger: log}
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

type itemEvent struct {
	ID string `json:"item_id"`
}

type leadEvent struct {
	ID string `json:"lead_id"`
}

func (a *aggregator) apply(event domain.Event) {

	switch event.EventType {
	case domain.ItemCreated:
		itemPayload, err := mapper.ItemDTOtoDomain(event.Payload)
		if err != nil {
			a.logger.Error(fmt.Sprintf("error mapping item with id: %d", itemPayload.ID), "error", err)
		}
		state, ok := a.Items[itemPayload.ID]
		if !ok {
			state = &domain.ItemState{
				ID:           itemPayload.ID,
				Name:         itemPayload.Name,
				SalesByDay:   make(map[string]int),
				RevenueByDay: make(map[string]float64),
			}
			a.Items[itemPayload.ID] = state
		}
	}
}
