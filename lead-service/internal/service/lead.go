package service

import (
	"log/slog"
	"time"

	itemproto "github.com/fiveret/crm-golang/grpc/item-grpc"
	"github.com/fiveret/crm-golang/internal/helpers"
	"github.com/fiveret/crm-golang/internal/models"
	"github.com/fiveret/crm-golang/internal/repository"
)

type LeadService interface {
	NewLead(lead *models.Lead) (string, *time.Time, error)
	DeleteLead(leadID uint32) (string, error)
	UpdateLead(lead *models.Lead) (*models.Lead, error)
	GetLead(leadID uint32) (*models.Lead, error)
	GetLeads() []*models.Lead
	AddProductsToLead(id uint32, product_id []uint32) (string, error)
	DeleteLeadProduct(id, productId uint32) (string, error)
	DeleteLeadProducts(id uint32) (string, error)
}

type leadService struct {
	repo       repository.LeadRepo
	itemClient itemproto.ItemServiceClient
	logger     *slog.Logger
}

func NewLeadService(repo repository.LeadRepo, logger *slog.Logger, itemClient itemproto.ItemServiceClient) LeadService {
	return &leadService{repo: repo, logger: logger, itemClient: itemClient}
}

func (s *leadService) NewLead(lead *models.Lead) (string, *time.Time, error) {
	if err := helpers.ValidateNewLead(lead); err != nil {
		return "", nil, err
	}
	return s.repo.CreateLead(lead)
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
