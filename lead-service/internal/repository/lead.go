package repository

import (
	"log/slog"

	"github.com/fiveret/crm-golang/internal/db"
	"github.com/fiveret/crm-golang/internal/models"
)

type LeadRepo interface {
	AddProducts(id uint32, product_id []string) (*models.Lead, error)
	CreateLead()
	DeleteLead()
	GetLead()
	GetLeads()
	UpdateLead()
}
type leadRepo struct {
	logger *slog.Logger
	db     *db.DBConnection
}

func NewLeadRepository(db *db.DBConnection, log *slog.Logger) LeadRepo {
	return &leadRepo{db: db, logger: log}
}
func (repo *leadRepo) CreateLead()
func (repo *leadRepo) DeleteLead()
func (repo *leadRepo) GetLead()
func (repo *leadRepo) GetLeads()
func (repo *leadRepo) UpdateLead()
