package service

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	itemproto "github.com/fiveret/crm-golang/grpc/item-grpc"
	"github.com/fiveret/crm-golang/internal/helpers"
	"github.com/fiveret/crm-golang/internal/repository"
)

type LeadProductService interface {
	AddProductsToLead(leadID uint32, productIDs []uint32) (string, error)
	DeleteLeadProduct(id, productId uint32) (string, error)
	DeleteLeadProducts(id uint32) (string, error)
}

func NewLeadProductService(repo repository.LeadRepo, logger *slog.Logger, itemClient itemproto.ItemServiceClient) LeadProductService {
	return &leadProductService{logger: logger, itemClient: itemClient, repo: repo}
}

type leadProductService struct {
	logger     *slog.Logger
	itemClient itemproto.ItemServiceClient
	repo       repository.LeadRepo
}

func (s *leadProductService) AddProductsToLead(leadID uint32, productIDs []uint32) (string, error) {
	lead, err := s.repo.GetLead(leadID)
	if err != nil {
		return "failure", err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(len(productIDs))

	for _, id := range productIDs {
		go func(id uint32) {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			resp, err := s.itemClient.GetItem(ctx, &itemproto.GetItemRequest{Id: id})
			if err != nil {
				s.logger.Error(fmt.Sprintf("error getting item %d: %v", id, err))
				return
			}
			createdAt := resp.CreatedAt.AsTime()
			updatedAt := resp.UpdatedAt.AsTime()
			product := helpers.GRPCProductToModels(resp.Item, createdAt, updatedAt)

			mu.Lock()
			lead.Products = append(lead.Products, product)
			mu.Unlock()
		}(id)
	}

	wg.Wait()

	if len(lead.Products) == 0 {
		return "failure", fmt.Errorf("couldn't add products to lead")
	}

	if _, err := s.repo.UpdateLead(lead); err != nil {
		return "failure", err
	}

	return "success", nil
}

func (service *leadProductService) DeleteLeadProduct(id, productId uint32) (string, error) {
	return service.repo.DeleteLeadProduct(id, productId)
}

func (service *leadProductService) DeleteLeadProducts(id uint32) (string, error) {
	return service.repo.DeleteLeadProducts(id)
}
