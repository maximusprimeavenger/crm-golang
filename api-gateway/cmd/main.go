package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/fiveret/api-gateway/internal/gateway"
	"github.com/fiveret/api-gateway/internal/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v3"
)

const env = "/app/config/conf.yaml"

func main() {
	logger, err := loadLogger(env)
	if err != nil {
		log.Fatal("error, logger is nil: ", err)
	}
	port, err := helpers.GetPort(env)
	if err != nil {
		logger.Error("error getting port", "error", err)
	}
	app := fiber.New()

	mux := runtime.NewServeMux()
	ctx := context.Background()
	credentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	opts := []grpc.DialOption{credentials}

	endpoints := []string{
		os.Getenv("ITEM_SERVICE_URL"),
		os.Getenv("LEAD_SERVICE_URL"),
	}
	err = gateway.RegisterHandlers(ctx, mux, endpoints, opts)
	if err != nil {
		logger.Error("error connecting to item-service or lead-service", "details:", err)
	}
	logger.Info("Registered all handlers!")
	fasthttpHandler := fasthttpadaptor.NewFastHTTPHandler(mux)

	app.All("/v1/*", func(c *fiber.Ctx) error {
		fasthttpHandler(c.Context())
		return nil
	})

	logger.Error("error listening on http server", "details:",
		app.Listen(fmt.Sprintf(":%d", *port)))
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
