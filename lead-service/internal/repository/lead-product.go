package repository

import "github.com/fiveret/crm-golang/internal/db"

type LeadProductRepo interface{}

type leadProductRepo struct {
	db *db.DBConnection
}

func NewLeadProductRepository(db *db.DBConnection) LeadProductRepo {
	return &leadProductRepo{db: db}
}
