package repository

import (
	"time"

	"github.com/fiveret/product-service/internal/db"
	"github.com/fiveret/product-service/internal/models"
)

type ItemRepository interface {
	NewItem(*models.Item) (*time.Time, error)
}

type itemRepo struct {
	db *db.DB
}

func NewItemRepo(db *db.DB) ItemRepository {
	return &itemRepo{db: db}
}

func (r *itemRepo) NewItem(item *models.Item) (*time.Time, error) {
	err := r.db.CreateItem(&models.Item{
		Name:        item.Name,
		Description: item.Description,
		Category:    item.Category,
		Price:       item.Price,
		InStock:     item.InStock,
	})
	if err != nil {
		return nil, err
	}

	newItem, err := r.db.FindItemByName(*item.Name)
	if err != nil {
		return nil, err
	}
	return &newItem.CreatedAt, nil
}

func GetItem(grpcModels *models.Item)
