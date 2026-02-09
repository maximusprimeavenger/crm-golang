package repository

import (
	"fmt"

	"github.com/fiveret/crm-golang/internal/models"
)

func (repo *leadRepo) DeleteLeadProducts(id uint32) (string, error) {
	lead, err := repo.GetLead(id)
	if err != nil {
		return "failure", err
	}
	if len(lead.Products) == 0 {
		return "nothing to delete", nil
	}
	lead.Products = []*models.Product{}
	return "lead's products have been successfully deleted", nil
}

func (repo *leadRepo) DeleteLeadProduct(id, productId uint32) (string, error) {
	lead, err := repo.GetLead(id)
	if err != nil {
		return "", err
	}

	var (
		deletedProductName string
		newProducts        []*models.Product
	)

	for _, p := range lead.Products {
		if p.ID == uint(productId) {
			deletedProductName = p.Name
			continue
		}
		newProducts = append(newProducts, p)
	}

	if deletedProductName == "" {
		return "", fmt.Errorf("product with id %d not found", productId)
	}

	lead.Products = newProducts
	err = repo.db.Save(lead).Error
	if err != nil {
		return "", fmt.Errorf("error saving updated lead")
	}

	return fmt.Sprintf("lead's product %s has successfully deleted", deletedProductName), nil
}
