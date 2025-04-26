package main

import (
	"log"

	"github.com/fiveret/product-service/internal/db"
	"github.com/fiveret/product-service/internal/helpers"
	"github.com/fiveret/product-service/internal/routes"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	port, err := helpers.GetPort()
	if err != nil {
		log.Fatal(err)
	}
	db, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}
	routes.ItemRoutes(app, db)
	app.Listen(port)
}
