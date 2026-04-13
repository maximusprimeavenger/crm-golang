package helpers

import (
	grpcModels "github.com/fiveret/item-service/grpc/models"
	models "github.com/fiveret/item-service/internal/domain"
)

func ModelsToGRPC(item *models.Item) *grpcModels.ItemResponse {
	return &grpcModels.ItemResponse{
		ProductId:   uint32(item.ID),
		Price:       item.Price,
		Name:        item.Name,
		Currency:    item.Currency,
		Category:    item.Category,
		Description: item.Description,
		Status:      item.Status,
		InStock:     item.InStock,
	}
}

func GRPCToModels(id uint32, item *grpcModels.ItemRequest) *models.Item {
	return &models.Item{
		ID:          uint(id),
		Name:        item.Name,
		Category:    item.Category,
		Currency:    item.Currency,
		Description: item.Description,
		Status:      item.Status,
		InStock:     item.InStock,
		Price:       item.Price,
	}
}

func GRPCToModelsUpdate(id uint32, item *grpcModels.ItemRequest) *models.UpdateItem {
	return &models.UpdateItem{
		ID:          uint(id),
		Name:        &item.Name,
		Category:    &item.Category,
		Currency:    &item.Currency,
		Description: &item.Description,
		Status:      &item.Status,
		InStock:     &item.InStock,
		Price:       &item.Price,
	}
}
