package service

import (
	"context"
	"errors"
	"time"

	"github.com/fiveret/product-service/internal/helpers"
	"github.com/fiveret/product-service/internal/models"
	"github.com/fiveret/product-service/internal/repository"
)

type ItemService interface {
	CreateItem(context.Context, *models.Item) (*time.Time, error)
	GetItem(context.Context, *uint32) (*models.Item, error)
	PutItem(context.Context, *uint32, *models.Item) (*models.Item, *time.Time, *time.Time, error)
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

func (s *itemService) PutItem(ctx context.Context, id *uint32, item *models.Item) (*models.Item, *time.Time, *time.Time, error) {
	if id == nil {
		return nil, nil, nil, errors.New("id is empty")
	}

	err := helpers.CheckItem(item)
	if err != nil {
		return nil, nil, nil, err
	}
	return s.repo.PutItem(id, item)
}
