package service

import (
	"context"
	"fmt"
	"time"

	"github.com/fiveret/crm-golang/internal/helpers"
	producer "github.com/fiveret/crm-golang/internal/kafka"
	"github.com/fiveret/crm-golang/internal/models"
	"github.com/fiveret/crm-golang/internal/repository"
)

type LeadService interface {
	NewLead(ctx context.Context, lead *models.Lead) (string, *time.Time, error)
	DeleteLead(leadID uint32) (string, error)
	UpdateLead(lead *models.Lead) (*models.Lead, error)
	GetLead(leadID uint32) (*models.Lead, error)
	GetLeads() []*models.Lead
}

type leadService struct {
	repo      repository.LeadRepo
	publisher producer.EventPublisher
}

func NewLeadService(r repository.LeadRepo, p producer.EventPublisher) LeadService {
	return &leadService{repo: r, publisher: p}
}

func (s *leadService) NewLead(ctx context.Context, lead *models.Lead) (string, *time.Time, error) {
	topic, err := helpers.GetTopic(0)
	if err != nil {
		return "", nil, fmt.Errorf("error getting topic: %v", err)
	}
	if err := helpers.ValidateNewLead(lead); err != nil {
		return "", nil, err
	}
	name, createdAt, err := s.repo.CreateLead(lead, topic)
	if err != nil {
		return "", nil, err
	}
	return name, createdAt, nil
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
