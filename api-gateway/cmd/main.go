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

	clients, err := clients.InitClients()
	if err != nil {
		log.Fatal(err)
	}
	handlers.ItemRoutes(app, clients)
	handlers.LeadRouteManager(app, clients)
	handlers.LeadProductManager(app, clients)
	app.Listen(port)
}
