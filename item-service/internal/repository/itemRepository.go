package repository

import (
	"time"

	"github.com/fiveret/product-service/internal/db"
	"github.com/fiveret/product-service/internal/models"
)

type ItemRepository interface {
	NewItem(*models.Item) (*time.Time, error)
	GetItem(*uint32) (*models.Item, error)
	PutItem(*uint32, *models.Item) (*models.Item, *time.Time, *time.Time, error)
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

func (r *itemRepo) GetItem(id *uint32) (*models.Item, error) {
	item, err := r.db.FindItem(*id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (r *itemRepo) PutItem(id *uint32, item *models.Item) (*models.Item, *time.Time, *time.Time, error) {
	item, err := r.db.UpdateItem(*id, item)
	if err != nil {
		return nil, nil, nil, err
	}
	return item, &item.CreatedAt, &item.UpdatedAt, nil
}
