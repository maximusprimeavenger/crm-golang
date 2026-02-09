package helpers

import (
	"time"

	grpcModels "github.com/fiveret/crm-golang/grpc/models"
	"github.com/fiveret/crm-golang/internal/models"
)

func GRPCProductToModels(product *grpcModels.ItemResponse, createdAt, updatedAt time.Time) *models.Product {
	return &models.Product{
		ID:          uint(product.ProductId),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
		Currency:    product.Currency,
		InStock:     uint(product.InStock),
		Status:      product.Status,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
