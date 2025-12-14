package handlers

import (
	"github.com/fiveret/api-gateway/internal/repository"
	"github.com/gofiber/fiber/v2"
)

const crmAnalyticsBaseURL = "http://crm-analytics:9092"

func PythonHandler(app *fiber.App) {

	app.Get("/sales-forecast", func(c *fiber.Ctx) error {
		return repository.ProxyToCRM(c, "/sales-forecast")
	})

	app.Get("/customer-clusters", func(c *fiber.Ctx) error {
		return repository.ProxyToCRM(c, "/customer-clusters")
	})

	app.Get("/gemma-analysis", func(c *fiber.Ctx) error {
		return repository.ProxyToCRM(c, "/gemma-analysis")
	})
}
