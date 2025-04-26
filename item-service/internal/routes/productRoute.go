package routes

import (
	"github.com/fiveret/product-service/internal/db"
	"github.com/fiveret/product-service/internal/repository"
	"github.com/gofiber/fiber"
)

func ItemRoutes(app *fiber.App, db *db.DB) {
	app.Post("/items", repository.CreateItem(db))
	app.Get("/items/:id", repository.GetItem(db))
	app.Get("/items", repository.GetItems(db))
	app.Put("/items/:id", repository.UpdateItem(db))
	app.Delete("/items/:id", repository.DeleteItem(db))
}
