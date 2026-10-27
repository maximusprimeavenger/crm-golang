package helpers

import (
	"time"

	grpcModels "github.com/fiveret/crm-golang/grpc/models"
	"github.com/fiveret/crm-golang/internal/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func LeadRequest(lead *grpcModels.LeadRequest) *models.Lead {
	if lead == nil {
		return nil
	}

	products := make([]*models.Product, 0, len(lead.Products))
	for _, p := range lead.Products {
		products = append(products, &models.Product{
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Category:    p.Category,
			Currency:    p.Currency,
			InStock:     uint(p.InStock),
			Status:      p.Status,
		})
	}

	var lastPurchaseAt *time.Time
	if lead.LastPurchaseAt != nil {
		t := lead.LastPurchaseAt.AsTime()
		lastPurchaseAt = &t
	}
	var lastVisit *time.Time
	if lead.LastVisit != nil {
		t := lead.LastVisit.AsTime()
		lastVisit = &t
	}

	return &models.Lead{
		Name:           lead.Name,
		Email:          lead.Email,
		Phone:          lead.Phone,
		Company:        lead.Company,
		Products:       products,
		Visits:         uint(lead.Visits),
		LastVisit:      lastVisit,
		TotalSales:     lead.TotalSales,
		LastPurchaseAt: lastPurchaseAt,
	}
}

func LeadResponse(lead *models.Lead) *grpcModels.LeadResponse {
	products := make([]*grpcModels.ItemResponse, 0)
	if lead == nil {
		return nil
	}
	if lead.Products != nil {
		products = make([]*grpcModels.ItemResponse, len(lead.Products))
		for i, p := range lead.Products {
			products[i] = &grpcModels.ItemResponse{
				ProductId:   uint32(p.ID),
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
				Category:    p.Category,
				Currency:    p.Currency,
				InStock:     uint32(p.InStock),
				Status:      p.Status,
			}
		}
	}
	var lastVisit *timestamppb.Timestamp
	if lead.LastVisit != nil {
		lastVisit = timestamppb.New(*lead.LastVisit)
	}
	var lastPurchaseAt *timestamppb.Timestamp
	if lead.LastPurchaseAt != nil {
		lastPurchaseAt = timestamppb.New(*lead.LastPurchaseAt)
	}
	return &grpcModels.LeadResponse{
		Id:             uint32(lead.ID),
		Name:           lead.Name,
		Email:          lead.Email,
		Phone:          lead.Phone,
		Company:        lead.Company,
		Products:       products,
		Visits:         uint32(lead.Visits),
		TotalSales:     lead.TotalSales,
		LastPurchaseAt: lastPurchaseAt,
		LastVisit:      lastVisit,
	}
}

func uint32Ptr(u *uint) uint32 {
	return uint32(*u)
}
