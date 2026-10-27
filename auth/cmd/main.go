package main

import (
	"fmt"
	"log"
	"net"
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
	if port == "" {
		log.Fatal("PORT_AUTH is not set")
	}

	if err := router.SetTrustedProxies([]string{}); err != nil {
		log.Fatal("Failed to set trusted proxies:", err)
	}
	err := controllers.AdminPresence()
	if err != nil {
		fmt.Print(err)
	}

	routes.AuthintificateRoute(router)
	routes.UserManager(router)
	routes.AdminManager(router)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error while trying to listen to port - %s", port)
	}
	log.Printf("Trying to run server on %s...", lis.Addr())
	if err := router.RunListener(lis); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
