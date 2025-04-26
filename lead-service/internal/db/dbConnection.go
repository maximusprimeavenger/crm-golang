package db

import (
	"fmt"
	"strconv"

	"github.com/fiveret/crm-golang/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DBConnection struct {
	db *gorm.DB
}

func Init() (*DBConnection, error) {
	db, err := gorm.Open(sqlite.Open("lead.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database")
	}
	fmt.Println("Successfully connected!")
	if err := db.AutoMigrate(&models.Lead{}, &models.Product{}); err != nil {
		return nil, fmt.Errorf("migration failed: %v", err)
	}
	fmt.Println("Successfully migrated!")
	return &DBConnection{db: db}, nil
}

func (db *DBConnection) FindLeadById(id string) (*models.Lead, error) {
	uintID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return nil, err
	}
	var lead models.Lead
	db.db.First(&lead, uintID)
	return &lead, nil
}

func (db *DBConnection) FindLeads() []models.Lead {
	var leads []models.Lead
	db.db.Find(&leads)
	return leads
}

func (db *DBConnection) DeleteLead(id string) error {
	uintID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return err
	}
	var lead *models.Lead
	err = db.db.First(&lead, uintID).Error
	if err != nil {
		return err
	}
	if lead.Name == "" {
		return fmt.Errorf("no lead found with id: %d", uintID)
	}
	err = db.db.Delete(&lead, uintID).Error
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
