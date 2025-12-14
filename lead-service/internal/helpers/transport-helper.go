package helpers

import (
	"sync"

	grpcModels "github.com/fiveret/crm-golang/grpc/models"
	"github.com/fiveret/crm-golang/internal/models"
	"golang.org/x/text/encoding/charmap"
)

func CP1251ToUTF8(s []byte) string {
	utf8Bytes, err := charmap.Windows1251.NewDecoder().Bytes(s)
	if err != nil {
		return string(s)
	}
	return string(utf8Bytes)
}
func LeadGRPCToModels(lead *grpcModels.Lead) *models.Lead {
	if lead == nil {
		return nil
	}

	convertedProducts := make([]*models.Product, len(lead.Products))
	for i, p := range lead.Products {
		instock := uint(p.InStock)
		leadID := uint(p.LeadID)
		convertedProducts[i] = &models.Product{
			Name:        &p.Name,
			Description: &p.Description,
			Price:       &p.Price,
			Category:    &p.Category,
			Currency:    &p.Currency,
			InStock:     &instock,
			Status:      &p.Status,
			LeadID:      &leadID,
		}
	}
	lastPurchaseDays := uint(lead.LastPurchaseDays)
	return &models.Lead{
		Name:             CP1251ToUTF8([]byte(lead.Name)),
		Email:            CP1251ToUTF8([]byte(lead.Email)),
		Phone:            CP1251ToUTF8([]byte(lead.Phone)),
		Company:          CP1251ToUTF8([]byte(lead.Company)),
		Products:         convertedProducts,
		Visits:           uint(lead.Visits),
		LastVisit:        &lead.LastVisit,
		TotalSales:       lead.TotalSales,
		LastPurchaseDays: &lastPurchaseDays,
	}
}

func ModelsToLeadGRPC(lead *models.Lead) *grpcModels.Lead {
	if lead == nil {
		return nil
	}

	var convertedProducts []*grpcModels.Product
	if len(lead.Products) > 0 {
		convertedProducts = make([]*grpcModels.Product, len(lead.Products))
		var wg sync.WaitGroup
		for i, product := range lead.Products {
			wg.Add(1)
			go func(idx int, p *models.Product) {
				defer wg.Done()
				var name, desc, category, currency, status string
				var price float64
				var instock, leadId uint32

				if p.Name != nil {
					name = *p.Name
				}
				if p.Description != nil {
					desc = *p.Description
				}
				if p.Category != nil {
					category = *p.Category
				}
				if p.Currency != nil {
					currency = *p.Currency
				}
				if p.Status != nil {
					status = *p.Status
				}
				if p.Price != nil {
					price = *p.Price
				}
				if p.InStock != nil {
					instock = uint32(*p.InStock)
				}
				if p.LeadID != nil {
					leadId = uint32(*p.LeadID)
				}

				convertedProducts[idx] = &grpcModels.Product{
					Name:        name,
					Description: desc,
					Price:       price,
					Category:    category,
					Currency:    currency,
					InStock:     instock,
					Status:      status,
					LeadID:      leadId,
				}
			}(i, product)
		}
		wg.Wait()
	}
	lastDays := *lead.LastPurchaseDays
	return &grpcModels.Lead{
		Name:             lead.Name,
		Email:            lead.Email,
		Phone:            lead.Phone,
		Company:          lead.Company,
		Products:         convertedProducts,
		Visits:           uint32(lead.Visits),
		TotalSales:       lead.TotalSales,
		LastPurchaseDays: uint32(lastDays),
		LastVisit:        *lead.LastVisit,
	}
}
