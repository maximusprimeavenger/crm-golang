package main

import (
	"log"

	"github.com/fiveret/api-gateway/grpc/clients"
	"github.com/fiveret/api-gateway/internal/handlers"
	"github.com/fiveret/api-gateway/internal/helpers"
	"github.com/gofiber/fiber"
)

func main() {
	port, err := helpers.GetPort()
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	grpcClients, err := clients.InitClients()
	if err != nil {
		log.Fatal(err)
	}
	handlers.ItemRoutes(app, grpcClients.ItemClient)
	handlers.LeadRouteManager(app, grpcClients.LeadClient)
	handlers.LeadProductManager(app, grpcClients.LeadProductClient)
	app.Listen(port)
}
