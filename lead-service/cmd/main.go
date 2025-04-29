package main

import (
	"log"

	"github.com/fiveret/crm-golang/internal/db"
	"github.com/fiveret/crm-golang/internal/helpers"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	port, err := helpers.FindPort()
	if err != nil {
		log.Fatal(err)
	}
	conn, err := db.Init()
	if err != nil {
		log.Fatal("Error with connecting to database")
	}
	app.Listen(port)

}
