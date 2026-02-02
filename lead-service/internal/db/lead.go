package db

import (
	"fmt"
	"strconv"

	"github.com/fiveret/crm-golang/internal/models"
)

func (db *DBConnection) FindLeadById(id uint32) (*models.Lead, error) {
	var lead models.Lead
	err := db.db.Preload("Products").First(&lead, id).Error
	if err != nil {
		return nil, err
	}
	return &lead, nil
}

func (db *DBConnection) FindLeads() []*models.Lead {
	var leads []*models.Lead
	db.db.Find(&leads)
	return leads
}

func (db *DBConnection) DeleteLead(id uint) error {
	var lead *models.Lead
	err := db.db.First(&lead, id).Error
	if err != nil {
		return err
	}
	if lead.Name == "" {
		return fmt.Errorf("no lead found with id: %d", id)
	}
	err = db.db.Delete(&lead, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *DBConnection) SaveLead(lead *models.Lead) error {
	err := db.db.Create(&lead).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *DBConnection) UpdateLead(name, email, company, phone, id string) error {
	uintID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return err
	}
	err = db.db.Model(&models.Lead{}).Where("id = ?", uintID).Updates(models.Lead{
		Name:    name,
		Email:   email,
		Phone:   phone,
		Company: company,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
