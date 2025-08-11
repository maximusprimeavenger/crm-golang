package service

import (
	"fmt"

	"github.com/fiveret/crm-golang/internal/repository"
)

type LeadProductService interface {
	AddProductsToLead(id uint32, product_id []uint32) (string, error)
	DeleteLeadProduct(id, productId uint32) (string, error)
	DeleteLeadProducts(id uint32) (string, error)
}

type productLeadService struct {
	repo repository.LeadRepo
}

func NewLeadProductService(repo repository.LeadRepo) LeadProductService {
	return &productLeadService{repo: repo}
}

func (service *productLeadService) AddProductsToLead(id uint32, product_id []uint32) (string, error) {
	user, err := service.repo.AddProducts(id, product_id)
	if err != nil {
		return "failure", err
	}
	if len(user.Products) == 0 {
		return "failure", fmt.Errorf("couldn't add products to lead")
	}
	return "success", nil
}

func (service *productLeadService) DeleteLeadProduct(id, productId uint32) (string, error) {
	return service.repo.DeleteLeadProduct(id, productId)
}

func (service *productLeadService) DeleteLeadProducts(id uint32) (string, error) {
	return service.repo.DeleteLeadProducts(id)
}
