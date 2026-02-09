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
	err = repo.db.Model(lead).Association("Products").Clear()
	if err != nil {
		return "", err
	}
	return "lead's products have been successfully deleted", nil
}

func (repo *leadRepo) DeleteLeadProduct(leadID, productID uint32) (string, error) {
	lead, err := repo.GetLead(leadID)
	if err != nil {
		return "", err
	}

	var product *models.Product
	for _, p := range lead.Products {
		if p.ID == uint(productID) {
			product = p
			break
		}
	}

	if product == nil {
		return "", fmt.Errorf("product with id %d not found", productID)
	}

	if err := repo.db.Model(lead).Association("Products").Delete(product); err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"lead's product %s has successfully deleted",
		product.Name,
	), nil
}
