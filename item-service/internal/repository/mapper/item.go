package mapper

import (
	"github.com/fiveret/item-service/internal/domain"
	"github.com/fiveret/item-service/internal/repository/models"
)

func DomainToDB(item *domain.Item) *models.Item {
	return &models.Item{
		ID:          item.ID,
		Name:        &item.Name,
		Description: &item.Description,
		Category:    &item.Category,
		InStock:     &item.InStock,
		Currency:    &item.Currency,
		Status:      &item.Status,
		NewPrice:    &item.Price,
	}
}

func DBToDomain(item *models.Item) *domain.Item {
	return &domain.Item{
		ID:          item.ID,
		Name:        *item.Name,
		Description: *item.Description,
		Category:    *item.Category,
		InStock:     *item.InStock,
		Currency:    *item.Currency,
		Status:      *item.Status,
		Price:       *item.NewPrice,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
}

func DomainUpdateToDB(item *domain.UpdateItem) *models.Item {
	return &models.Item{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Category:    item.Category,
		InStock:     item.InStock,
		Currency:    item.Currency,
		Status:      item.Status,
		NewPrice:    item.Price,
	}
}

func DBUpdateToDomain(item *models.Item) *domain.UpdateItem {
	return &domain.UpdateItem{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		InStock:     item.InStock,
		Category:    item.Category,
		Currency:    item.Currency,
		Status:      item.Status,
		Price:       item.NewPrice,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
}
