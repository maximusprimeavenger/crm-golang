package repository

import (
	"time"

	"github.com/fiveret/product-service/internal/db"
	"github.com/fiveret/product-service/internal/models"
)

func NewItem(grpcModels *models.Item, db *db.DB) (*time.Time, error) {
	err := db.CreateItem(&models.Item{
		Name:        grpcModels.Name,
		Description: grpcModels.Description,
		Category:    grpcModels.Category,
		Price:       grpcModels.Price,
		InStock:     grpcModels.InStock,
	})
	if err != nil {
		return nil, err
	}

	item, err := db.FindItemByName(*grpcModels.Name)
	if err != nil {
		return nil, err
	}
	return &item.CreatedAt, nil
}
