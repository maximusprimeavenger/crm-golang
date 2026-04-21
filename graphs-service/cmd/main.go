package main

import (
	"context"
	"fmt"
	"graphs-service/internal/application/aggregator"
	linechart "graphs-service/internal/application/use_cases/line_chart"
	domain "graphs-service/internal/entities"
	reader "graphs-service/internal/transport/kafka-reader"
	"log/slog"
	"os"
	"strconv"

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
	eventsChan := make(chan domain.Event, 1000)

	itemReader, leadReader, leadProductReader, err := reader.LoadReaders(itemTopic, leadTopics, address, logger)
	if err != nil {
		logger.Error(fmt.Sprintf("Error creating one of the readers: %v", err))
		os.Exit(1)
	}

	go itemReader.Start(ctx, eventsChan)
	go leadReader.Start(ctx, eventsChan)
	go leadProductReader.Start(ctx, eventsChan)

	agg := aggregator.NewAggregator(logger)
	go agg.Run(ctx, eventsChan)
	r := gin.Default()
	r.GET("/line-chart/:item_id", func(c *gin.Context) {
		if !agg.Ready() {
			logger.Error("aggregator is null")
			c.JSON(503, gin.H{"error": "not ready"})
			return
		}
		itemID := c.Param("item_id")
		id, err := strconv.ParseUint(itemID, 10, 0)
		if err != nil {
			logger.Error("error", "error parsing uint during request line-chart/:item_id", err)
			return
		}
		linechart.DrawLineChartByItem(agg.GetSalesAndRevenueByID(uint(id)))
	})
	//routes.RouteManager(r)

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
