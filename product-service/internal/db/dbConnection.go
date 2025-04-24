package db

import (
	"fmt"
	"os"

	"github.com/fiveret/product-service/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func Init() (*DB, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("PORT_PRODUCT"))
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = conn.AutoMigrate(&models.Product{})
	if err != nil {
		return nil, fmt.Errorf("error migrating")
	}
	fmt.Println("Connected and Migrated!")
	return &DB{db: conn}, nil
}
