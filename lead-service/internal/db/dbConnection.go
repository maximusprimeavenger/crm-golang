package db

import (
	"fmt"
	"os"

	"github.com/fiveret/crm-golang/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("PORT_SQL"))
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = conn.AutoMigrate(&models.Lead{}, &models.OutboxEvent{})
	if err != nil {
		return nil, fmt.Errorf("error migrating")
	}
	fmt.Println("Connected and Migrated!")
	return conn, nil
}
