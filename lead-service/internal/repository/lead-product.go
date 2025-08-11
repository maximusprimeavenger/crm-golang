package repository

import (
	"fmt"

	"github.com/fiveret/crm-golang/internal/models"
)

func (repo *leadRepo) AddProducts(id uint32, product_id []uint32) (*models.Lead, error) {
	products, err := repo.db.GetProducts(product_id)
	if err != nil {
		return nil, err
	}
	lead, err := repo.db.FindLeadById(id)
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		if product != nil {
			lead.Products = append(lead.Products, *product)
		} else if product.ID != 0 && product.Name == nil {
			repo.logger.Warn(fmt.Sprintf("the product with id: %d is null", product.ID))
		}
	}
	return lead, nil
}

func (repo *leadRepo) DeleteLeadProducts(id uint32) (string, error) {
	lead, err := repo.db.FindLeadById(id)
	if err != nil {
		return "failure", err
	}
	lead.Products = []models.Product{}
	return "lead's products have been successfully deleted", nil
}

func (repo *leadRepo) DeleteLeadProduct(id, productId uint32) (string, error) {
	lead, err := repo.db.FindLeadById(id)
	if err != nil {
		return "", err
	}

	var (
		deletedProductName string
		newProducts        []models.Product
	)

	for _, p := range lead.Products {
		if p.ID == uint(productId) {
			deletedProductName = *p.Name
			continue
		}
		newProducts = append(newProducts, p)
	}

	if deletedProductName == "" {
		return "", fmt.Errorf("product with id %d not found", productId)
	}

	lead.Products = newProducts

	return fmt.Sprintf("lead's product %s has successfully deleted", deletedProductName), nil
}
