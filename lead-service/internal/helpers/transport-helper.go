package helpers

import (
	"time"

	grpcModels "github.com/fiveret/crm-golang/grpc/models"
	"github.com/fiveret/crm-golang/internal/models"
	"golang.org/x/text/encoding/charmap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CP1251ToUTF8(s []byte) string {
	utf8Bytes, err := charmap.Windows1251.NewDecoder().Bytes(s)
	if err != nil {
		return string(s)
	}
	return string(utf8Bytes)
}

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

	var lastPurchaseDays *uint
	if lead.LastPurchaseDays > 0 {
		v := uint(lead.LastPurchaseDays)
		lastPurchaseDays = &v
	}
	var lastVisit *time.Time
	if lead.LastVisit != nil {
		t := lead.LastVisit.AsTime()
		lastVisit = &t
	}

	return &models.Lead{
		Name:             CP1251ToUTF8([]byte(lead.Name)),
		Email:            CP1251ToUTF8([]byte(lead.Email)),
		Phone:            CP1251ToUTF8([]byte(lead.Phone)),
		Company:          CP1251ToUTF8([]byte(lead.Company)),
		Products:         products,
		Visits:           uint(lead.Visits),
		LastVisit:        lastVisit,
		TotalSales:       lead.TotalSales,
		LastPurchaseDays: lastPurchaseDays,
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
	var lastPurchaseDays uint32
	if lead.LastPurchaseDays != nil {
		lastPurchaseDays = uint32Ptr(lead.LastPurchaseDays)
	}
	return &grpcModels.LeadResponse{
		Id:               uint32(lead.ID),
		Name:             lead.Name,
		Email:            lead.Email,
		Phone:            lead.Phone,
		Company:          lead.Company,
		Products:         products,
		Visits:           uint32(lead.Visits),
		TotalSales:       lead.TotalSales,
		LastPurchaseDays: lastPurchaseDays,
		LastVisit:        lastVisit,
	}
}

func uint32Ptr(u *uint) uint32 {
	return uint32(*u)
}
