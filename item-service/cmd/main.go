package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	proto "github.com/fiveret/product-service/grpc/item-grpc"
	"github.com/fiveret/product-service/internal/db"
	"github.com/fiveret/product-service/internal/helpers"
	"github.com/fiveret/product-service/internal/repository"
	"github.com/fiveret/product-service/internal/service"
	"github.com/fiveret/product-service/internal/transport"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
)

const path = "/app/config/conf.yaml"

func main() {
	logger, err := loadLogger(path)
	if err != nil {
		logger.Info("error setting up the right env for logger")
	}
	port, err := helpers.GetPort()
	if err != nil {
		logger.Error("error getting port", "error", err)
		os.Exit(1)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logger.Error("error listening", "port", *port)
		os.Exit(1)
	}
	s := grpc.NewServer()

	dbConn, err := db.Init()
	if err != nil {
		logger.Error("error initializing db", "error", err)
		os.Exit(1)
	}
	repo := repository.NewItemRepo(dbConn)
	svc := service.NewItemService(repo)
	handler := transport.NewGRPCHandler(svc)

	proto.RegisterItemServiceServer(s, handler)
	logger.Info("server is running", "port", port)
	if err := s.Serve(lis); err != nil {
		logger.Error("error serving", "port", port)
		os.Exit(1)
	}
}

func loadLogger(path string) (*slog.Logger, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	logger := new(yamLogger)
	err = yaml.Unmarshal(body, &logger)
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

type yamLogger struct {
	env string `yaml:"env"`
}
