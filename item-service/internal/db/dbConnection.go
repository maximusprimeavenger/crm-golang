package db

import (
	"fmt"
	"os"

	"github.com/fiveret/item-service/internal/repository/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	err := godotenv.Load("/app/.env")
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("PORT_SQL"))
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = conn.AutoMigrate(&models.Item{}, &models.OutboxEvent{})
	if err != nil {
		return nil, fmt.Errorf("error migrating")
	}
	fmt.Println("Connected and Migrated!")
	return conn, nil
}
