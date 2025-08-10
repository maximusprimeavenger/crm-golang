package service

import (
	"fmt"

	"github.com/fiveret/crm-golang/internal/repository"
)

type LeadProductService interface {
	AddProductsToLead(id uint32, product_id []string) (string, error)
}

type productLeadService struct {
	repo repository.LeadRepo
}

func NewLeadProductService(repo repository.LeadRepo) LeadProductService {
	return &productLeadService{repo: repo}
}

func (service *productLeadService) AddProductsToLead(id uint32, product_id []string) (string, error) {
	user, err := service.repo.AddProducts(id, product_id)
	if err != nil {
		return "failure", err
	}
	if len(user.Products) == 0 {
		return "failure", fmt.Errorf("couldn't add products to lead")
	}
	return "success", nil
}
