package service

import (
	"time"

	"github.com/fiveret/crm-golang/internal/helpers"
	"github.com/fiveret/crm-golang/internal/models"
	"github.com/fiveret/crm-golang/internal/repository"
)

type LeadService interface {
	NewLead(lead *models.Lead) (*time.Time, error)
}

type leadService struct {
	repo repository.LeadRepo
}

func NewLeadService(repo repository.LeadRepo) LeadService {
	return &leadService{repo: repo}
}

func (s *leadService) NewLead(lead *models.Lead) (*time.Time, error) {
	err := helpers.ValidateNewLead(lead)
	if err != nil {
		return nil, err
	}
	return s.repo.CreateLead(lead)
}
