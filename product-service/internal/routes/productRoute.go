package routes

import (
	"github.com/fiveret/product-service/internal/db"
	"github.com/fiveret/product-service/internal/repository"
	"github.com/gofiber/fiber"
)

func ProductRoutes(app *fiber.App, db *db.DB) {
	app.Post("/products", repository.CreateProduct(db))
	app.Get("/products/:id", repository.GetProduct(db))
	app.Get("/products", repository.GetProducts(db))
	app.Put("/products/:id", repository.UpdateProduct(db))
	app.Delete("/products/:id", repository.DeleteProduct(db))
}
