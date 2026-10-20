package main

import (
	"fmt"
	"log"
	"os"

	"go/auth-service/internal/controllers"
	"go/auth-service/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}

	router := gin.Default()

	port := os.Getenv("PORT_AUTH")

	ip := os.Getenv("IP")
	if ip == "" {
		log.Fatal("No trusted IP address provided")
	}

	if err := router.SetTrustedProxies([]string{ip}); err != nil {
		log.Fatal("Failed to set trusted proxies:", err)
	}
	err := controllers.AdminPresence()
	if err != nil {
		fmt.Print(err)
	}

	routes.AuthintificateRoute(router)
	routes.UserManager(router)
	routes.AdminManager(router)

	addr := fmt.Sprintf("%s:%s", ip, port)
	log.Printf("Trying to run server on %s...", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
