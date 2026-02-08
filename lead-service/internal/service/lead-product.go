package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	itemproto "github.com/fiveret/crm-golang/grpc/item-grpc"
	"github.com/fiveret/crm-golang/internal/helpers"
)

func (s *leadService) AddProductsToLead(leadID uint32, productIDs []uint32) (string, error) {
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

			product := helpers.GRPCProductToModels(uint(id), resp.Item, *resp.CreatedAt, *resp.UpdatedAt)

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

func (service *leadService) DeleteLeadProduct(id, productId uint32) (string, error) {
	return service.repo.DeleteLeadProduct(id, productId)
}

func (service *leadService) DeleteLeadProducts(id uint32) (string, error) {
	return service.repo.DeleteLeadProducts(id)
}
