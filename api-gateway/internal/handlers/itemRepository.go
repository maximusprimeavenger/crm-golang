package handlers

import (
	proto "github.com/fiveret/api-gateway/grpc/item-grpc"
	"github.com/fiveret/api-gateway/internal/repository"
	"github.com/gofiber/fiber"
)

func ItemRoutes(app *fiber.App, c proto.ItemServiceClient) {
	app.Group("/items")
	{
		app.Patch("/:id", repository.UpdateItem())
	}

}
