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

func (db *DB) CreateProduct(product *models.Product) error {
	err := db.db.Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) UpdateProduct(id string, name, description, currency, category, status *string, price *float64, instock *uint) error {
	err := db.db.Model(&models.Product{}).Where("id = ?", id).Updates(models.Product{
		Name:        name,
		Description: description,
		Category:    category,
		Currency:    currency,
		Price:       price,
		InStock:     instock,
		Status:      status,
	}).Error
	if err != nil {
		return fmt.Errorf("error updating product: %v", err)
	}
	return nil
}

func (db *DB) FindProduct(id string) (*models.Product, error) {
	product := new(models.Product)
	err := db.db.First(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (db *DB) FindProducts() []*models.Product {
	var products []*models.Product
	db.db.First(&products)
	return products
}
