package repository

import (
	"fmt"
	"time"

	"github.com/fiveret/item-service/internal/db"
	"github.com/fiveret/item-service/internal/domain"
	"github.com/fiveret/item-service/internal/repository/mapper"
)

type ItemRepository interface {
	GetItem(*uint32) (*domain.Item, error)
	GetItems() ([]*domain.Item, error)
	DeleteItem(*uint32) (string, error)
	NewItem(*domain.Item) (*time.Time, error)
	PutItem(*uint32, *domain.UpdateItem) (*domain.Item, *time.Time, *time.Time, error)
}

type itemRepo struct {
	db *db.DB
}

func NewItemRepo(db *db.DB) ItemRepository {
	return &itemRepo{db: db}
}

func (r *itemRepo) DeleteItem(id *uint32) (string, error) {
	item, err := r.db.FindItem(*id)
	if err != nil {
		return "", fmt.Errorf("error finding item for deleting")
	}
	err = r.db.DeleteItem(*id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("item %s has successfully deleted", *item.Name), nil
}

func (r *itemRepo) NewItem(item *domain.Item) (*time.Time, error) {
	err := r.db.CreateItem(mapper.DomainToDB(item))
	if err != nil {
		return nil, err
	}

	newItem, err := r.db.FindItemByName(item.Name)
	if err != nil {
		return nil, err
	}
	return &newItem.CreatedAt, nil
}

func (r *itemRepo) GetItem(id *uint32) (*domain.Item, error) {
	item, err := r.db.FindItem(*id)
	if err != nil {
		return nil, err
	}
	return mapper.DBToDomain(item), nil
}

func (r *itemRepo) GetItems() ([]*domain.Item, error) {
	items, err := r.db.FindItems()
	if err != nil {
		return nil, err
	}
	serviceItem := []*domain.Item{}
	for _, i := range items {
		serviceItem = append(serviceItem, mapper.DBToDomain(i))
	}
	return serviceItem, nil
}

func (r *itemRepo) PutItem(id *uint32, item *domain.UpdateItem) (*domain.Item, *time.Time, *time.Time, error) {
	itemReq, err := r.db.UpdateItem(*id, mapper.DomainUpdateToDB(item))
	if err != nil {
		return nil, nil, nil, err
	}
	return mapper.DBToDomain(itemReq), &item.CreatedAt, &item.UpdatedAt, nil
}
