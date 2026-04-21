package main

import (
	"log"
	"log/slog"
	"net"
	"os"
	"strconv"

	proto "github.com/fiveret/item-service/grpc/item-grpc"
	"github.com/fiveret/item-service/internal/db"
	"github.com/fiveret/item-service/internal/helpers"
	"github.com/fiveret/item-service/internal/repository"
	"github.com/fiveret/item-service/internal/service"
	"github.com/fiveret/item-service/internal/transport"
	"google.golang.org/grpc"
)

const path = "/app/config/config.yaml"

func main() {
	logger, err := loadLogger()
	if err != nil {
		log.Fatal("logger is nil, error:", err)
	}
	config, err := helpers.LoadConfig(path)
	if err != nil {
		log.Fatal("config is nil, error:", err)
	}
	port := config.ItemService.Port
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		logger.Error("error listening", "port", port)
		os.Exit(1)
	}
	s := grpc.NewServer()

	dbConn, err := db.Init()
	if err != nil {
		logger.Error("error initializing db", "error", err)
		os.Exit(1)
	}
	repo := repository.NewItemRepo(dbConn)
	svc := service.NewItemService(repo, config.ItemService.KafkaWriter.Topic)
	handler := transport.NewGRPCHandler(svc)

	proto.RegisterItemServiceServer(s, handler)
	logger.Info("server is running", "port", port)
	if err := s.Serve(lis); err != nil {
		logger.Error("error serving", "port", port)
		os.Exit(1)
	}
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
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}
	return slog.New(handler), nil
}
