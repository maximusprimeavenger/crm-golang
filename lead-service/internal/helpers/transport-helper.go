package helpers

import (
	"sync"

	grpcModels "github.com/fiveret/crm-golang/grpc/models"
	"github.com/fiveret/crm-golang/internal/models"
)

func LeadGRPCToModels(lead *grpcModels.Lead) *models.Lead {
	if lead == nil {
		return nil
	}
	convertedProducts := make([]*models.Product, 0)
	if len(lead.Products) > 0 {
		convertedProducts = make([]*models.Product, len(lead.Products))
		var wg sync.WaitGroup
		for i, product := range lead.Products {
			wg.Add(1)
			go func(m int, prod *grpcModels.Product) {
				defer wg.Done()
				instock := product.InStock
				uintStock := uint(instock)
				leadId := product.LeadID
				uintId := uint(leadId)
				convertedProduct := &models.Product{
					Name:        &product.Name,
					Description: &product.Description,
					Price:       &product.Price,
					Category:    &product.Category,
					Currency:    &product.Currency,
					InStock:     &uintStock,
					Status:      &product.Status,
					LeadID:      &uintId,
				}
				convertedProducts[i] = convertedProduct
			}(i, product)
		}
		wg.Wait()
	}
	return &models.Lead{
		Name:     lead.Name,
		Email:    lead.Email,
		Phone:    lead.Phone,
		Company:  lead.Company,
		Products: convertedProducts,
	}
}
