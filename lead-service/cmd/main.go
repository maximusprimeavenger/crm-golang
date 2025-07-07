package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"

	"github.com/fiveret/crm-golang/internal/db"
	"github.com/fiveret/crm-golang/internal/helpers"
	"github.com/gofiber/fiber"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
)

const path = "/app/config/conf.yaml"

func main() {
	app := fiber.New()
	logger, err := loadLogger(path)
	if err != nil {
		log.Fatal(err)
	}
	port, err := helpers.FindPort(path)
	if err != nil {
		log.Fatal(err)
	}
	dbConn, err := db.Init()
	if err != nil {
		logger.Error("error connecting to db")
		os.Exit(1)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Error(fmt.Sprintf("error connecting with port :%d", *port))
		os.Exit(1)
	}
	s := grpc.NewServer()
}

func loadLogger(path string) (*slog.Logger, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	logger := new(yamlLogger)
	err = yaml.Unmarshal(data, &logger)
	if err != nil {
		return nil, err
	}
	var handler *slog.TextHandler
	switch logger.env {
	case "dev":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	case "test":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case "prod":
		handler = (*slog.TextHandler)(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}))
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}
	return slog.New(handler), nil
}

type yamlLogger struct {
	env string `yaml:"env"`
}
