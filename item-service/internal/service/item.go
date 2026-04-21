package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	grpcModels "github.com/fiveret/item-service/grpc/models"
	"github.com/fiveret/item-service/internal/domain"
	"github.com/fiveret/item-service/internal/helpers"
	"github.com/fiveret/item-service/internal/repository"
	"github.com/fiveret/item-service/internal/service/mapper"
	"github.com/fiveret/item-service/internal/service/models"
	"github.com/google/uuid"
)

type ItemService interface {
	CreateItem(context.Context, *domain.Item) (*time.Time, error)
	DeleteItem(context.Context, *uint32) (string, error)
	GetItem(context.Context, *uint32) (*domain.Item, error)
	GetItems(context.Context) ([]*grpcModels.ItemResponse, error)
	PutItem(context.Context, *uint32, *domain.UpdateItem) (*domain.Item, error)
}

type itemService struct {
	repo  repository.ItemRepository
	topic string
}

func NewItemService(repo repository.ItemRepository, topic string) ItemService {
	return &itemService{repo: repo, topic: topic}
}

func (s *itemService) CreateItem(ctx context.Context, item *domain.Item) (*time.Time, error) {
	if item.Name == "" {
		return nil, errors.New("item name is required")
	}
	if item.Price < 0 {
		return nil, errors.New("price cannot be negative")
	}
	if item.InStock == 0 {
		return nil, errors.New("instock must be greater than 0")
	}

	payload, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	event := &models.Event{
		EventID:    uuid.New().String(),
		EventType:  fmt.Sprintf("%s.created", s.topic),
		OccurredAt: time.Now(),
		Payload:    payload,
	}
	return s.repo.NewItem(item, mapper.EventToOutboxEvent(event, item, s.topic))
}

func (s *itemService) DeleteItem(ctx context.Context, id *uint32) (string, error) {
	if id == nil {
		return "", errors.New("id is empty")
	}
	return s.repo.DeleteItem(id)
}

func (s *itemService) GetItem(ctx context.Context, id *uint32) (*domain.Item, error) {
	if id == nil {
		return nil, errors.New("id is empty")
	}
	return s.repo.GetItem(id)
}

func (s *itemService) GetItems(ctx context.Context) ([]*grpcModels.ItemResponse, error) {
	items, err := s.repo.GetItems()
	if err != nil {
		return nil, err
	}
	returnItems := make([]*grpcModels.ItemResponse, len(items))
	var wg sync.WaitGroup
	for i, item := range items {
		wg.Add(1)
		go func(i int, item *domain.Item) {
			defer wg.Done()
			returnItems[i] = helpers.ModelsToGRPC(item)
		}(i, item)
	}

	wg.Wait()
	return returnItems, nil
}

func (s *itemService) PutItem(ctx context.Context, id *uint32, item *domain.UpdateItem) (*domain.Item, error) {
	if id == nil {
		return nil, errors.New("id is empty")
	}

	err := helpers.CheckItem(item)
	if err != nil {
		return nil, err
	}
	newItem, err := s.repo.PutItem(id, item)
	if err != nil {
		return nil, err
	}
	return newItem, nil
}
