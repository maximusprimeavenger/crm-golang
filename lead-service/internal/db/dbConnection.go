package db

import (
	"fmt"
	"os"

	"github.com/fiveret/crm-golang/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DBConnection struct {
	db *gorm.DB
}

func Init() (*DBConnection, error) {
	dbPath := os.Getenv("SQLITE_PATH")
	if dbPath == "" {
		return nil, fmt.Errorf("error finding the SQLITE_PATH")
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
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
