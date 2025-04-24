package main

import (
	"log"

	"github.com/fiveret/product-service/internal/helpers"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	port, err := helpers.GetPort()
	if err != nil {
		log.Fatal(err)
	}
	//ROUTES
	app.Listen(port)
}
