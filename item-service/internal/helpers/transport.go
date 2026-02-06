package helpers

import (
	grpcModels "github.com/fiveret/product-service/grpc/models"
	"github.com/fiveret/product-service/internal/models"
)

func ConvertModelsToGRPCResponse(item *models.Item) *grpcModels.ItemResponse {
	return &grpcModels.ItemResponse{
		ProductId:   uint32(item.ID),
		Price:       *item.Price,
		Name:        *item.Name,
		Currency:    *item.Currency,
		Category:    *item.Category,
		Description: *item.Description,
		Status:      *item.Status,
		InStock:     *item.InStock,
	}
}

func ConvertGRPCToModelsRequest(id uint32, item *grpcModels.ItemRequest) *models.Item {
	itemModel := &models.Item{
		Name:        &item.Name,
		Category:    &item.Category,
		Currency:    &item.Currency,
		Description: &item.Description,
		Status:      &item.Status,
		InStock:     &item.InStock,
		Price:       &item.Price,
	}
	itemModel.ID = uint(id)
	return itemModel
}
