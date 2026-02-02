package service

import (
	"time"

	"github.com/fiveret/crm-golang/internal/helpers"
	"github.com/fiveret/crm-golang/internal/models"
	"github.com/fiveret/crm-golang/internal/repository"
)

type LeadService interface {
	NewLead(lead *models.Lead) (*time.Time, error)
	AddProducts(leadID uint32, productIDs []uint32) (*models.Lead, error)
	DeleteLeadProduct(leadID, productID uint32) (string, error)
	DeleteLeadProducts(leadID uint32) (string, error)
	DeleteLead(leadID uint32) (string, error)
	UpdateLead(lead *models.Lead) (*models.Lead, error)
	GetLead(leadID uint32) (*models.Lead, error)
	GetLeads() []*models.Lead
}

type leadService struct {
	repo repository.LeadRepo
}

func NewLeadService(repo repository.LeadRepo) LeadService {
	return &leadService{repo: repo}
}

func (s *leadService) NewLead(lead *models.Lead) (*time.Time, error) {
	if err := helpers.ValidateNewLead(lead); err != nil {
		return nil, err
	}
	return s.repo.CreateLead(lead)
}

func (s *leadService) AddProducts(leadID uint32, productIDs []uint32) (*models.Lead, error) {
	return s.repo.AddProducts(leadID, productIDs)
}

func (s *leadService) DeleteLeadProduct(leadID, productID uint32) (string, error) {
	return s.repo.DeleteLeadProduct(leadID, productID)
}

func (s *leadService) DeleteLeadProducts(leadID uint32) (string, error) {
	return s.repo.DeleteLeadProducts(leadID)
}

func (s *leadService) DeleteLead(leadID uint32) (string, error) {
	return s.repo.DeleteLead(leadID)
}

func (s *leadService) UpdateLead(lead *models.Lead) (*models.Lead, error) {
	return s.repo.UpdateLead(lead)
}

func (s *leadService) GetLead(leadID uint32) (*models.Lead, error) {
	lead, err := s.repo.GetLead(leadID)
	if err != nil {
		return nil, err
	}
	return lead, nil
}

func (s *leadService) GetLeads() []*models.Lead {
	return s.repo.GetLeads()
}
