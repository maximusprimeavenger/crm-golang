package helpers

import (
	"errors"

	models "github.com/fiveret/item-service/internal/domain"
)

func CheckItem(item *models.UpdateItem) error {
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
