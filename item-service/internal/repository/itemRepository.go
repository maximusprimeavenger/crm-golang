package repository

import (
	"fmt"
	"time"

	"github.com/fiveret/item-service/internal/domain"
	"github.com/fiveret/item-service/internal/repository/mapper"
	"github.com/fiveret/item-service/internal/repository/models"
	"gorm.io/gorm"
)

type ItemRepository interface {
	GetItem(*uint32) (*domain.Item, error)
	GetItems() ([]*domain.Item, error)
	DeleteItem(*uint32) (string, error)
	NewItem(item *domain.Item, event *models.OutboxEvent) (*time.Time, error)
	PutItem(*uint32, *domain.UpdateItem) (*domain.Item, error)
}

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) ItemRepository {
	return &itemRepo{db: db}
}

func (r *itemRepo) DeleteItem(id *uint32) (string, error) {
	var name string
	err := r.db.Raw("DELETE FROM items WHERE id = ? RETURNING name", *id).Scan(&name).Error
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("item %s has successfully deleted", name), nil
}

func (r *itemRepo) NewItem(item *domain.Item, event *models.OutboxEvent) (*time.Time, error) {
	dbItem := mapper.DomainToDB(item)
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&dbItem).Error; err != nil {
			return err
		}
		if err := tx.Create(&event).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &dbItem.CreatedAt, nil
}

func (r *itemRepo) GetItem(id *uint32) (*domain.Item, error) {
	item := new(models.Item)
	err := r.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return mapper.DBToDomain(item), nil
}

func (r *itemRepo) GetItems() ([]*domain.Item, error) {
	items := []*models.Item{}
	err := r.db.Find(&items).Error
	if err != nil {
		return nil, err
	}
	serviceItem := []*domain.Item{}
	for _, i := range items {
		serviceItem = append(serviceItem, mapper.DBToDomain(i))
	}
	return serviceItem, nil
}

func (r *itemRepo) PutItem(id *uint32, item *domain.UpdateItem) (*domain.Item, error) {
	err := r.db.Model(&models.Item{}).
		Where("id = ?", *id).
		Updates(map[string]interface{}{
			"name":        item.Name,
			"description": item.Description,
			"category":    item.Category,
			"currency":    item.Currency,
			"old_price":   gorm.Expr("new_price"),
			"new_price":   item.Price,
			"in_stock":    item.InStock,
			"status":      item.Status,
		}).Error

	if err != nil {
		return nil, fmt.Errorf("error updating item: %w", err)
	}

	var updated models.Item
	if err := r.db.First(&updated, *id).Error; err != nil {
		return nil, err
	}

	return mapper.DBToDomain(&updated), nil
}
