package repository

import "github.com/fiveret/crm-golang/internal/db"

type LeadRepo interface {
	CreateLead()
	DeleteLead()
	GetLead()
	GetLeads()
	UpdateLead()
}
type leadRepo struct {
	db *db.DBConnection
}

func NewLeadRepository(db *db.DBConnection) LeadRepo {
	return &leadRepo{db: db}
}
func (repo *leadRepo) CreateLead()
func (repo *leadRepo) DeleteLead()
func (repo *leadRepo) GetLead()
func (repo *leadRepo) GetLeads()
func (repo *leadRepo) UpdateLead()
