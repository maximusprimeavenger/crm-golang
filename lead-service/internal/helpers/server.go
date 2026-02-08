package helpers

import (
	grpcModels "github.com/fiveret/crm-golang/grpc/models"
	"github.com/fiveret/crm-golang/internal/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GRPCProductToModels(leadID uint, product *grpcModels.ItemResponse, createdAt, updatedAt timestamppb.Timestamp) *models.Product {
	return &models.Product{
		ID:          uint(product.ProductId),
		LeadID:      leadID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
		Currency:    product.Currency,
		InStock:     uint(product.InStock),
		Status:      product.Status,
		CreatedAt:   createdAt.AsTime(),
		UpdatedAt:   updatedAt.AsTime(),
	}
}
