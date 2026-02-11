package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fiveret/api-gateway/internal/gateway"
	"github.com/fiveret/api-gateway/internal/handlers"
	"github.com/fiveret/api-gateway/internal/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"google.golang.org/grpc"
)

const path = "/app/config/config.yaml"

func main() {
	logger, err := helpers.LoadLogger(path)
	if err != nil {
		log.Fatalf("Logger is not set: %v", err)
		os.Exit(1)
	}

	port, err := helpers.GetPort(path)
	if err != nil {
		logger.Error(fmt.Sprintf("error getting port: %v", err))
	}
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	ctx := context.Background()
	mux := runtime.NewServeMux()

	err = gateway.RegisterHandlers(
		ctx,
		mux,
		opts,
	)
	if err != nil {
		logger.Error("failed to register handlers", "error", err)
		os.Exit(1)
	}

	logger.Info("gRPC Gateway handlers registered")

	app := fiber.New()
	handler := fasthttpadaptor.NewFastHTTPHandler(mux)
	handlers.PythonHandler(app)
	app.All("/v1/*", func(c *fiber.Ctx) error {
		handler(c.Context())
		return nil
	})

	logger.Info("Starting API Gateway", "port", port)
	log.Fatal(app.Listen(":" + port))
}
