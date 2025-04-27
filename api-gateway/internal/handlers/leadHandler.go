package handlers

import (
	"github.com/fiveret/api-gateway/internal/repository"
	"github.com/gofiber/fiber"
)

func LeadRouteManager(app *fiber.App) {
	app.Group("/lead")
	{
		app.Get("/:id", repository.GetLead())
		app.Get("", repository.GetLeads())
		app.Delete("/:id", repository.DeleteLead())
		app.Post("", repository.NewLead())
		app.Put("/:id", repository.UpdateLead())
	}

}

func LeadProductManager(app *fiber.App) {
	app.Group("/lead/:id/products")
	{
		app.Post("", repository.AddProductToLead())
		app.Get("", repository.GetLeadProducts())
		app.Delete("", repository.DeleteLeadProducts())
		app.Put("", repository.PutProductsLead())
	}
	app.Get("/lead/products", repository.GetLeadsProducts())
}
