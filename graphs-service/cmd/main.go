package main

import (
	"context"
	"fmt"
	"graphs-service/internal/application/aggregator"
	"graphs-service/internal/application/dispatcher"
	"graphs-service/internal/transport/http/routes"
	reader "graphs-service/internal/transport/kafka-reader"
	"graphs-service/internal/transport/kafka-reader/models"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-yaml"
)

const path = "/app/config/config.yaml"

var address = []string{"kafka:9092"}

func main() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	logger := slog.New(handler)
	logger.Info("Graphs-service has been started")
	port := os.Getenv("PORT_GRAPHS")
	if port == "" {
		logger.Error("Error, port is empty")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	leadTopics, itemTopic, err := loadTopics(path)
	if err != nil {
		logger.Error(fmt.Sprintf("Error loading topics for leads: %v", err))
		os.Exit(1)
	}
	eventsChan := make(chan models.Envelope, 1000)

	itemReader, _, _, err := reader.LoadReaders(itemTopic, leadTopics, address, logger)
	if err != nil {
		logger.Error(fmt.Sprintf("Error creating one of the readers: %v", err))
		os.Exit(1)
	}

	go itemReader.Start(ctx, eventsChan)
	//go leadReader.Start(ctx, eventsChan)
	//go leadProductReader.Start(ctx, eventsChan)

	agg := aggregator.NewAggregator(logger)
	go dispatcher.Dispatcher(ctx, agg, itemReader.Reader(), eventsChan)
	r := gin.Default()
	routes.RouteManager(r, agg)

	r.Run(":" + port)
}

func loadTopics(path string) ([]string, string, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, "", err
	}
	reader := new(ultimateReader)
	err = yaml.Unmarshal(body, reader)
	if err != nil {
		return nil, "", err
	}
	return reader.Lead.KafkaLeadWriter.Topics, reader.Item.KafkaItemWriter.Topic, nil
}

type ultimateReader struct {
	Lead struct {
		KafkaLeadWriter struct {
			Topics []string `yaml:"topics"`
		} `yaml:"kafka-writer"`
	} `yaml:"lead-service"`
	Item struct {
		KafkaItemWriter struct {
			Topic string `yaml:"topic"`
		} `yaml:"kafka-writer"`
	} `yaml:"item-service"`
}
