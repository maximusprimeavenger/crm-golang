package routes

import (
	"github.com/fiveret/crm-golang/internal/db"
	"github.com/fiveret/crm-golang/internal/repository"
	"github.com/gofiber/fiber"
)

func LeadRouteManager(app *fiber.App, db *db.DBConnection) {
	app.Get("/lead/:id", repository.GetLead(db))
	app.Get("/lead", repository.GetLeads(db))
	app.Delete("/lead/:id", repository.DeleteLead(db))
	app.Post("/lead", repository.NewLead(db))
	app.Put("/lead/:id", repository.UpdateLead(db))
}
