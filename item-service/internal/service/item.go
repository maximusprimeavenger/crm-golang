package service

import (
	"context"
	"errors"
	"time"

	"github.com/fiveret/product-service/internal/models"
	"github.com/fiveret/product-service/internal/repository"
)

type ItemService interface {
	CreateItem(ctx context.Context, item *models.Item) (*time.Time, error)
	GetItem(ctx context.Context, id *uint32) (*models.Item, error)
}

type itemService struct {
	repo repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) ItemService {
	return &itemService{repo: repo}
}

func (s *itemService) CreateItem(ctx context.Context, item *models.Item) (*time.Time, error) {
	if item.Name == nil || *item.Name == "" {
		return nil, errors.New("item name is required")
	}
	if item.Price != nil && *item.Price < 0 {
		return nil, errors.New("price cannot be negative")
	}
	if item.InStock == nil {
		return nil, errors.New("zero instock value")
	}
	return s.repo.NewItem(item)
}

func (s *itemService) GetItem(ctx context.Context, id *uint32) (*models.Item, error) {
	if id == nil {
		return nil, errors.New("id is empty")
	}
	return s.repo.GetItem(id)
}
