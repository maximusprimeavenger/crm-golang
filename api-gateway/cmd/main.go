package main

import (
	"context"
	"crypto/tls"
	"log"
	"log/slog"
	"os"

	"github.com/fiveret/api-gateway/internal/gateway"
	"github.com/gofiber/fiber/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/yaml.v3"
)

const env = "/app/config/conf.yaml"

func main() {
	logger, err := loadLogger(env)
	if err != nil {
		log.Fatal("Logger is not set")
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	itemServiceURL := os.Getenv("ITEM_SERVICE_URL")
	leadServiceURL := os.Getenv("LEAD_SERVICE_URL")

	if itemServiceURL == "" || leadServiceURL == "" {
		log.Fatal("ITEM_SERVICE_URL or LEAD_SERVICE_URL is not set")
	}

	creds := credentials.NewTLS(&tls.Config{})
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	ctx := context.Background()
	mux := runtime.NewServeMux()

	err = gateway.RegisterHandlers(
		ctx,
		mux,
		[]string{
			itemServiceURL,
			leadServiceURL,
		},
		opts,
	)
	if err != nil {
		logger.Error("failed to register handlers", "error", err)
		os.Exit(1)
	}

	logger.Info("gRPC Gateway handlers registered")

	app := fiber.New()
	handler := fasthttpadaptor.NewFastHTTPHandler(mux)

	app.All("/v1/*", func(c *fiber.Ctx) error {
		handler(c.Context())
		return nil
	})

	logger.Info("Starting API Gateway", "port", port)
	log.Fatal(app.Listen(":" + port))
}
func loadLogger(env string) (*slog.Logger, error) {
	body, err := os.ReadFile(env)
	if err != nil {
		return nil, err
	}
	logger := new(yamLogger)
	err = yaml.Unmarshal(body, &logger)
	if err != nil {
		return nil, err
	}
	var handler *slog.TextHandler
	switch logger.Env {
	case "dev":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	case "test":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case "prod":
		handler = (*slog.TextHandler)(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}))
	}
	return slog.New(handler), nil
}

type yamLogger struct {
	Env string `yaml:"env"`
}
