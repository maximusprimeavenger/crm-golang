package handlers

import (
	"github.com/fiveret/api-gateway/internal/repository"
	"github.com/gofiber/fiber"
)

func ItemRoutes(app *fiber.App) {
	app.Group("/items")
	{
		app.Post("", repository.CreateItem())
		app.Get("/:id", repository.GetItem())
		app.Get("", repository.GetItems())
		app.Put("/:id", repository.UpdateItem())
		app.Delete("/:id", repository.DeleteItem())
	}

}
