package main

import (
	"context"
	"log"
	"log/slog"
	"net"
	"os"

	itemproto "github.com/fiveret/crm-golang/grpc/item-grpc"
	proto "github.com/fiveret/crm-golang/grpc/lead-grpc"
	"github.com/fiveret/crm-golang/internal/db"
	producer "github.com/fiveret/crm-golang/internal/kafka"
	"github.com/fiveret/crm-golang/internal/repository"
	"github.com/fiveret/crm-golang/internal/service"
	"github.com/fiveret/crm-golang/internal/transport"
	"github.com/fiveret/crm-golang/internal/worker"
	"google.golang.org/grpc"
)

func main() {
	logger, err := loadLogger()
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT_LEAD")
	if port == "" {
		logger.Error("error setting the port", "message", "port is nil")
		os.Exit(1)
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Error("error connecting with port :" + port)
		os.Exit(1)
	}
	s := grpc.NewServer()
	dbConn, err := db.Init()
	if err != nil {
		logger.Error("error connecting to db")
		os.Exit(1)
	}
	repo := repository.NewLeadRepository(dbConn, logger)
	repoEvent := repository.NewEventRepo(dbConn)
	conn, err := grpc.Dial("item-service:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect item-service: %v", err)
	}
	defer conn.Close()

	p := producer.NewKafkaPublisher([]string{"kafka:9092"})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	w := worker.NewWorker(repoEvent, logger, p)
	go w.StartWorker(ctx)

	itemClient := itemproto.NewItemServiceClient(conn)
	publisher := producer.NewKafkaPublisher([]string{"kafka:9092"})
	serv1 := service.NewLeadService(repo, publisher)
	serv2 := service.NewLeadProductService(repo, logger, itemClient)
	handler := transport.NewGRPCHandler(serv1, serv2)

	proto.RegisterLeadServiceServer(s, handler)
	proto.RegisterLeadProductServiceServer(s, handler)
	logger.Info("server is running", "port", port)
	err = s.Serve(lis)
	if err != nil {
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
