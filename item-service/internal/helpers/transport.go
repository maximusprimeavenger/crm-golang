package helpers

import (
	grpcModels "github.com/fiveret/product-service/grpc/models"
	"github.com/fiveret/product-service/internal/models"
)

func ConvertModelsToGRPC(item *models.Item) *grpcModels.Item {
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

func ConvertGRPCToModels(item *grpcModels.Item) *models.Item {
	return &models.Item{
		Name:        &item.Name,
		Category:    &item.Category,
		Currency:    &item.Currency,
		Description: &item.Description,
		Status:      &item.Status,
		InStock:     &item.InStock,
		Price:       &item.Price,
	}
}
