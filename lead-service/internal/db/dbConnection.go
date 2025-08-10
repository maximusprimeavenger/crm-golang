package db

import (
	"fmt"

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
