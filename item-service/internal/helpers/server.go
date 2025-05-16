package helpers

import (
	"errors"

	"github.com/fiveret/product-service/internal/models"
)

func CheckItem(item *models.Item) error {
	switch {
	case item.Name == nil:
		return errors.New("name is null")
	case item.Category == nil:
		return errors.New("category is null")
	case item.Description == nil:
		return errors.New("description is null")
	case item.InStock == nil:
		return errors.New("in stock value is null")
	case item.Price == nil:
		return errors.New("price is null")
	}
	return nil
}
