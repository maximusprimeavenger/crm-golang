package helpers

import (
	grpcModels "github.com/fiveret/product-service/grpc/models"
	"github.com/fiveret/product-service/internal/models"
)

func ConvertItem(item *models.Item) *grpcModels.Item {
	return &grpcModels.Item{
		Price:       *item.Price,
		Name:        *item.Name,
		Currency:    *item.Currency,
		Category:    *item.Category,
		Description: *item.Description,
		Status:      *item.Status,
		InStock:     *item.InStock,
	}
}
