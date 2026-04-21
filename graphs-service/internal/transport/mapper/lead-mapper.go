package mapper

import (
	"encoding/json"
	domain "graphs-service/internal/entities"
	models "graphs-service/internal/transport/dto"
)

func LeadDTOToDomain(payload json.RawMessage) (*domain.Lead, error) {
	lead, err := parseLeadPayload(payload)
	if err != nil {
		return nil, err
	}
	return &domain.Lead{
		ID:               lead.ID,
		Name:             lead.Name,
		Email:            lead.Email,
		Phone:            lead.Phone,
		Company:          lead.Company,
		Visits:           lead.Visits,
		Products:         dtoProductsToDomain(lead.Products),
		LastVisit:        lead.LastVisit,
		TotalSales:       lead.TotalSales,
		LastPurchaseDays: lead.LastPurchaseDays,
		CreatedAt:        lead.CreatedAt,
		UpdatedAt:        lead.UpdatedAt,
	}, nil
}

func parseLeadPayload(payload json.RawMessage) (*models.Lead, error) {
	item := new(models.Lead)
	err := json.Unmarshal(payload, item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func dtoProductsToDomain(dtoProducts []*models.Product) []*domain.Product {
	domainProducts := []*domain.Product{}
	for _, dtoProduct := range dtoProducts {
		domainProduct := &domain.Product{
			ID:          dtoProduct.ID,
			Name:        dtoProduct.Name,
			Description: dtoProduct.Description,
			Price:       dtoProduct.Price,
			Category:    dtoProduct.Category,
			Currency:    dtoProduct.Currency,
			InStock:     dtoProduct.InStock,
			Status:      dtoProduct.Status,
			CreatedAt:   dtoProduct.CreatedAt,
			UpdatedAt:   dtoProduct.UpdatedAt,
		}
		domainProducts = append(domainProducts, domainProduct)
	}
	return domainProducts
}
