package repository

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/fiveret/crm-golang/internal/models"
	"gorm.io/gorm"
)

type LeadRepo interface {
	CreateLead(lead *models.Lead) (string, *time.Time, error)
	DeleteLead(id uint32) (string, error)
	DeleteLeadProduct(id, productId uint32) (string, error)
	DeleteLeadProducts(id uint32) (string, error)
	GetLead(id uint32) (*models.Lead, error)
	GetLeads() []*models.Lead
	UpdateLead(lead *models.Lead) (*models.Lead, error)
}
type leadRepo struct {
	logger *slog.Logger
	db     *gorm.DB
}

func NewLeadRepository(db *gorm.DB, log *slog.Logger) LeadRepo {
	return &leadRepo{db: db, logger: log}
}
func (repo *leadRepo) CreateLead(lead *models.Lead) (string, *time.Time, error) {
	err := repo.db.Create(lead).Error
	if err != nil {
		return "", nil, err
	}
	foundLead, err := repo.GetLead(uint32(lead.ID))
	if err != nil {
		return "", nil, err
	}
	return fmt.Sprintf("Lead %s has successfully created!", lead.Name), &foundLead.CreatedAt, err
}

func (repo *leadRepo) DeleteLead(id uint32) (string, error) {
	lead, err := repo.GetLead(id)
	if err != nil {
		return "", err
	}
	err = repo.db.Delete(&lead, id).Error
	if err != nil {
		return "failure", err
	}
	return fmt.Sprintf("lead %s has been successfully deleted", lead.Name), nil
}

func (repo *leadRepo) GetLead(id uint32) (*models.Lead, error) {
	lead := &models.Lead{}
	err := repo.db.Preload("Products").First(lead, id).Error
	if err != nil {
		return nil, err
	}
	return lead, nil
}

func (repo *leadRepo) GetLeads() []*models.Lead {
	leads := []*models.Lead{}
	repo.db.Find(&leads)
	return leads
}

func (repo *leadRepo) UpdateLead(lead *models.Lead) (*models.Lead, error) {
	existingLead, err := repo.GetLead(uint32(lead.ID))
	if err != nil {
		return nil, err
	}

	if lead.Name != "" {
		existingLead.Name = lead.Name
	}
	if lead.Email != "" {
		existingLead.Email = lead.Email
	}
	if lead.Phone != "" {
		existingLead.Phone = lead.Phone
	}
	if lead.Company != "" {
		existingLead.Company = lead.Company
	}
	if lead.Visits != 0 {
		existingLead.Visits = lead.Visits
	}
	if lead.LastVisit != nil {
		existingLead.LastVisit = lead.LastVisit
	}
	if lead.TotalSales != 0 {
		existingLead.TotalSales = lead.TotalSales
	}
	if lead.LastPurchaseDays != nil {
		existingLead.LastPurchaseDays = lead.LastPurchaseDays
	}
	if len(lead.Products) > 0 {
		existingLead.Products = lead.Products
	}

	err = repo.db.Save(existingLead).Error
	if err != nil {
		return nil, err
	}

	return existingLead, nil
}
