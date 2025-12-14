package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/fiveret/api-gateway/internal/gateway"
	"github.com/fiveret/api-gateway/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"google.golang.org/grpc"
)

func main() {
	logger, err := loadLogger()
	if err != nil {
		log.Fatal("Logger is not set")
		os.Exit(1)
	}

	port := os.Getenv("PORT_GATEWAY")
	if port == "" {
		port = "9090"
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
func loadLogger() (*slog.Logger, error) {
	env := os.Getenv("ENV")
	var handler *slog.TextHandler
	switch env {
	case "dev":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	case "test":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case "prod":
		handler = (*slog.TextHandler)(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}))
	}
	return slog.New(handler), nil
}
