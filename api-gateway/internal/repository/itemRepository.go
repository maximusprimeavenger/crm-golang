package repository

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func CreateItem() fiber.Handler {
	return func(c *fiber.Ctx) {
		product := new(models.Item)
		err := c.BodyParser(product)
		if err != nil {
			c.Send(fmt.Sprintf("error parsing body: %v", err))
			return
		}
		db.CreateItem(product)
	}
}

func GetItem() fiber.Handler {
	return func(c *fiber.Ctx) {
		id := c.Params("id")
		if id == "" {
			c.Status(500).Send("empty id")
			return
		}
		product, err := db.FindItem(id)
		if err != nil {
			c.Status(500).Send(err)
			return
		}
		c.Status(201).JSON(fiber.Map{"message": "Product found!",
			"Product": product,
		})
	}
}

func GetItems() fiber.Handler {
	return func(c *fiber.Ctx) {
		products := db.FindItems()
		c.Status(200).JSON(fiber.Map{
			"message": "Products found!",
			"List":    products,
		})
	}
}

func UpdateItem() fiber.Handler {
	return func(c *fiber.Ctx) {
		id := c.Params("id")
		if id == "" {
			c.Status(500).Send("empty id")
			return
		}
		product := new(models.Item)
		err := c.BodyParser(product)
		if err != nil {
			c.Status(500).Send(fmt.Sprintf("error parsing body: %v", err))
			return
		}
		err = db.UpdateItem(id, product.Name, product.Description, product.Currency, product.Category, product.Status, product.Price, product.InStock)
		if err != nil {
			c.Status(500).Send(err)
			return
		}
		updatedProd, err := db.FindItem(id)
		if err != nil {
			c.Status(500).Send("Erorr finding item")
			return
		}
		c.Status(200).JSON(fiber.Map{"message": "Item updated successfully!",
			"Product": updatedProd,
		})
	}
}

func DeleteItem() fiber.Handler {
	return func(c *fiber.Ctx) {
		id := c.Params("id")
		if id == "" {
			c.Status(500).Send("id is missing")
			return
		}
		err := db.DeleteItem(id)
		if err != nil {
			c.Status(500).Send(err)
			return
		}
		c.Status(201).Send("Item successfully deleted!")
	}
}
