package repository

import (
	"time"

	grpc "github.com/fiveret/product-service/grpc/models"
	"github.com/fiveret/product-service/internal/db"
	"github.com/fiveret/product-service/internal/models"
)

func NewItem(grpcItem *grpc.Item, db *db.DB) (*time.Time, error) {
	err := db.CreateItem(&models.Item{
		Name:        &grpcItem.Name,
		Description: &grpcItem.Description,
		Category:    &grpcItem.Category,
		Price:       &grpcItem.Price,
		InStock:     &grpcItem.InStock,
	})
	if err != nil {
		return nil, err
	}

	item, err := db.FindItemByName(grpcItem.Name)
	if err != nil {
		return nil, err
	}
	return &item.CreatedAt, nil
}

func GetItem(grpcModels *models.Item)
