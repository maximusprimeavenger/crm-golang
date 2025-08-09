package service

import "github.com/fiveret/crm-golang/internal/repository"

type LeadService interface {
}

type leadService struct {
	repo repository.LeadRepo
}

func NewLeadService(repo repository.LeadRepo) LeadService {
	return &leadService{repo: repo}
}
