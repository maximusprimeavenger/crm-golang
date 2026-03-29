package mapper

import (
	"encoding/json"
	"graphs-service/internal/domain"
	models "graphs-service/internal/interfaces/dto"
)

func ItemDTOtoDomain(payload json.RawMessage) (*domain.Item, error) {
	item, err := parsePayload(payload)
	if err != nil {
		return nil, err
	}
	return &domain.Item{
		ID:          item.ID,
		Name:        *item.Name,
		Description: *item.Description,
		Currency:    *item.Currency,
		Status:      *item.Status,
		Price:       *item.Price,
		InStock:     *item.InStock,
		Category:    *item.Category,
		UpdatedAt:   item.UpdatedAt,
		CreatedAt:   item.CreatedAt,
	}, nil
}

func parsePayload(payload json.RawMessage) (*models.Item, error) {
	item := new(models.Item)
	err := json.Unmarshal(payload, item)
	if err != nil {
		return nil, err
	}
	return item, nil
}
