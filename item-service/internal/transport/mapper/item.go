package mapper

import (
	models "github.com/fiveret/item-service/grpc/models"
	"github.com/fiveret/item-service/internal/domain"
)

func ProtoToDomain(item *models.ItemRequest) *domain.Item {
	return &domain.Item{
		Name:        item.Name,
		Category:    item.Category,
		Description: item.Description,
		Currency:    item.Currency,
		InStock:     item.InStock,
		Price:       item.Price,
		Status:      item.Status,
	}
}
