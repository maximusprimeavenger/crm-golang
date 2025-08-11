package db

import "github.com/fiveret/crm-golang/internal/models"

func (db *DBConnection) GetProducts(product_ids []uint32) ([]*models.Product, error) {
	var products []*models.Product
	for _, val := range product_ids {
		var product *models.Product
		err := db.db.Where("id = ?", val).Scan(product).Error
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
