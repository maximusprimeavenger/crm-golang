package aggregator

import (
	"encoding/json"
	"fmt"
	domain "graphs-service/internal/entities"
	models "graphs-service/internal/interfaces/dto"
)

func (a *aggregator) eventItemCreated(event domain.Event) {
	itemPayload := new(models.ItemCreatedPayload)
	err := json.Unmarshal(event.Payload, itemPayload)
	if err != nil {
		a.logger.Error("error", "error unmarshalling the item with id: %d", itemPayload.ID)
		return
	}
	_, exists := a.Items[itemPayload.ID]
	if !exists {
		a.Items[itemPayload.ID] = &domain.ItemAnalytics{
			ID:           itemPayload.ID,
			Name:         itemPayload.Name,
			SalesByDay:   make(map[string]int),
			RevenueByDay: make(map[string]float64),
			Status:       itemPayload.Status,
			Price:        itemPayload.Price,
		}
	}
}

func (a *aggregator) eventItemPriceChanged(event domain.Event) {
	itemPayload := new(models.ItemPriceChangedPayload)
	err := json.Unmarshal(event.Payload, itemPayload)
	if err != nil {
		a.logger.Error("error", "error unmarshalling the item with id: %d", itemPayload.ID)
		return
	}

	state, exists := a.Items[itemPayload.ID]
	if !exists {
		a.logger.Error(fmt.Sprintf("item not found for update: %d", itemPayload.ID))
		return
	}
	state.PriceHistory[itemPayload.ChangedAt.Format("2006-01-02")] = itemPayload.NewPrice
}

func (a *aggregator) eventItemSold(event domain.Event) {
	itemPayload := new(models.ItemSoldPayload)
	err := json.Unmarshal(event.Payload, itemPayload)
	if err != nil {
		a.logger.Error("error", "error unmarshalling the item with id: %d", itemPayload.ID)
		return
	}
	state, exists := a.Items[itemPayload.ID]
	if !exists {
		a.logger.Error(fmt.Sprintf("item not found for update(sold): %d", itemPayload.ID))
		return
	}
	state.SalesByDay[itemPayload.SoldAt.Format("2006-01-02")] += itemPayload.Quantity
	state.RevenueByDay[itemPayload.SoldAt.Format("2006-01-02")] += float64(itemPayload.Quantity) * itemPayload.Price
}
