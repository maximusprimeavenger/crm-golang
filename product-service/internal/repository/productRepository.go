package repository

import (
	"fmt"

	"github.com/fiveret/product-service/internal/db"
	"github.com/fiveret/product-service/internal/models"
	"github.com/gofiber/fiber"
)

func CreateProduct(db *db.DB) fiber.Handler {
	return func(c *fiber.Ctx) {
		product := new(models.Product)
		err := c.BodyParser(product)
		if err != nil {
			c.Send(fmt.Sprintf("error parsing body: %v", err))
			return
		}
		db.CreateProduct(product)
	}
}

func GetProduct(db *db.DB) fiber.Handler {
	return func(c *fiber.Ctx) {
		id := c.Params("id")
		if id == "" {
			c.Status(500).Send("empty id")
			return
		}
		product, err := db.FindProduct(id)
		if err != nil {
			c.Status(500).Send(err)
			return
		}
		c.Status(201).JSON(fiber.Map{"message": "Product found!",
			"Product": product,
		})
	}
}

func GetProducts(db *db.DB) fiber.Handler {
	return func(c *fiber.Ctx) {
		products := db.FindProducts()
		c.Status(200).JSON(fiber.Map{
			"message": "Products found!",
			"List":    products,
		})
	}
}

func UpdateProduct(db *db.DB) fiber.Handler {
	return func(c *fiber.Ctx) {
		id := c.Params("id")
		if id == "" {
			c.Status(500).Send("empty id")
			return
		}
		product := new(models.Product)
		err := c.BodyParser(product)
		if err != nil {
			c.Status(500).Send(fmt.Sprintf("error parsing body: %v", err))
			return
		}
		err = db.UpdateProduct(id, product.Name, product.Description, product.Currency, product.Category, product.Status, product.Price, product.InStock)
		if err != nil {
			c.Status(500).Send(err)
			return
		}
		updatedProd, err := db.FindProduct(id)
		if err != nil {
			c.Status(500).Send("Erorr finding product")
			return
		}
		c.Status(200).JSON(fiber.Map{"message": "Product updated successfully!",
			"Product": updatedProd,
		})
	}
}

func DeleteProduct(db *db.DB) fiber.Handler {
	return func(c *fiber.Ctx) {

	}
}
