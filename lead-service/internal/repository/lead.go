package repository

import (
	"fmt"
	"log/slog"

	"github.com/fiveret/crm-golang/internal/db"
	"github.com/fiveret/crm-golang/internal/models"
)

type LeadRepo interface {
	AddProducts(id uint32, product_id []uint32) (*models.Lead, error)
	CreateLead(lead *models.Lead) (*models.Lead, error)
	DeleteLead(id uint32) (string, error)
	DeleteLeadProduct(id, productId uint32) (string, error)
	DeleteLeadProducts(id uint32) (string, error)
	GetLead(id uint32) (*models.Lead, error)
	GetLeads() []*models.Lead
	UpdateLead(lead *models.Lead) (*models.Lead, error)
}
type leadRepo struct {
	logger *slog.Logger
	db     *db.DBConnection
}

func NewLeadRepository(db *db.DBConnection, log *slog.Logger) LeadRepo {
	return &leadRepo{db: db, logger: log}
}
func (repo *leadRepo) CreateLead(lead *models.Lead) (*models.Lead, error) {
	err := repo.db.SaveLead(lead)
	if err != nil {
		return nil, err
	}
	foundLead, err := repo.db.FindLeadById(uint32(lead.ID))
	if err != nil {
		return nil, err
	}

	return foundLead, err
}

func (repo *leadRepo) DeleteLead(id uint32) (string, error) {
	lead, err := repo.db.FindLeadById(id)
	if err != nil {
		return "", err
	}
	err = repo.db.DeleteLead(lead.ID)
	if err != nil {
		return "failure", err
	}
	return fmt.Sprintf("lead %s has been successfully deleted", lead.Name), nil
}
func (repo *leadRepo) GetLead(id uint32) (*models.Lead, error) {
	lead, err := repo.db.FindLeadById(id)
	if err != nil {
		return nil, err
	}
	return lead, nil
}

func (repo *leadRepo) GetLeads() []*models.Lead {
	return repo.db.FindLeads()
}

func (repo *leadRepo) UpdateLead(lead *models.Lead) (*models.Lead, error) {
	existingLead, err := repo.db.FindLeadById(uint32(lead.ID))
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

	err = repo.db.SaveLead(existingLead)
	if err != nil {
		return nil, err
	}

	return existingLead, nil
}
