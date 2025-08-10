package repository

import (
	"fmt"

	"github.com/fiveret/crm-golang/internal/models"
)

func (repo *leadRepo) AddProducts(id uint32, product_id []string) (*models.Lead, error) {
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
