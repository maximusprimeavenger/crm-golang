package service

import "github.com/fiveret/crm-golang/internal/repository"

type LeadProductService interface {
}

type productLeadService struct {
	repo repository.LeadProductRepo
}

func NewLeadProductService(repo repository.LeadProductRepo) LeadProductService {
	return &productLeadService{repo: repo}
}
