package handlers

import (
	proto "github.com/fiveret/api-gateway/grpc/lead-grpc"
	"github.com/fiveret/api-gateway/internal/repository"
	"github.com/gofiber/fiber"
)

func LeadRouteManager(app *fiber.App, c proto.LeadServiceClient) {
	app.Group("/lead")
	{
		//app.Get("/:id", repository.GetLead(c))
		//app.Get("", repository.GetLeads(c))
		//app.Delete("/:id", repository.DeleteLead(c))
		//app.Post("", repository.NewLead(c))
		//TODO
		app.Patch("/:id", repository.UpdateLead(c))
	}

}

func LeadProductManager(app *fiber.App, c proto.LeadProductServiceClient) {
	// app.Group("/lead/:id/products")
	// {
	// 	app.Post("", repository.AddProductsToLead(c))
	// 	app.Get("", repository.GetLeadProducts(c))
	// 	app.Delete("", repository.DeleteLeadProducts(c))
	// 	app.Put("", repository.PutProductsLead(c))
	// }
	// app.Get("/lead/products", repository.GetLeadsProducts(c))
}
