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
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("PORT_SQL"))
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = conn.AutoMigrate(&models.Item{})
	if err != nil {
		return nil, fmt.Errorf("error migrating")
	}
	fmt.Println("Connected and Migrated!")
	return &DB{db: conn}, nil
}

func (db *DB) CreateItem(item *models.Item) error {
	err := db.db.Create(&item).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) UpdateItem(id uint32, item *models.Item) (*models.Item, error) {
	err := db.db.Model(&models.Item{}).Where("id = ?", id).Updates(models.Item{
		Name:        item.Name,
		Description: item.Description,
		Category:    item.Category,
		Currency:    item.Currency,
		Price:       item.Price,
		InStock:     item.InStock,
		Status:      item.Status,
	}).Error
	if err != nil {
		return nil, fmt.Errorf("error updating item: %v", err)
	}
	db.db.First(&item)
	return item, nil
}

func (db *DB) FindItem(id uint32) (*models.Item, error) {
	product := new(models.Item)
	err := db.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (db *DB) FindItems() []*models.Item {
	var items []*models.Item
	db.db.First(&items)
	return items
}

func (db *DB) DeleteItem(id uint32) error {
	var item *models.Item
	err := db.db.First(&item, id).Error
	if err != nil {
		return fmt.Errorf("error finding item: %v", err)
	}
	err = db.db.Delete(&item).Error
	if err != nil {
		return fmt.Errorf("error deleting item: %v", err)
	}
	return nil
}

func (db *DB) FindItemByName(name string) (*models.Item, error) {
	item := new(models.Item)
	err := db.db.First(&item, name).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}
