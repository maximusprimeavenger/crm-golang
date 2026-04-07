package main

import (
	"context"
	"fmt"
	"graphs-service/internal/application/aggregator"
	linechart "graphs-service/internal/application/use_cases/line_chart"
	domain "graphs-service/internal/entities"
	reader "graphs-service/internal/interfaces/kafka-reader"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	itemTopic, err := loadItemTopic(path)
	if err != nil {
		logger.Error(fmt.Sprintf("Error loading topic for item: %v", err))
		os.Exit(1)
	}
	leadTopics, err := loadLeadTopics(path)
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
		itemID := c.Param("item_id")
		id, err := strconv.ParseUint(itemID, 10, 0)
		if err != nil {
			logger.Error("error", "error parsing uint during request line-chart/:item_id", err)
			return
		}
		linechart.DrawLineChartByItem(agg.GetSalesAndRevenueByID(uint(id)))
	})
}

func loadLeadTopics(path string) ([]string, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lead := new(leadReader)
	err = yaml.Unmarshal(body, lead)
	if err != nil {
		return nil, err
	}
	return lead.Lead.Writer.Topics, nil
}

func loadItemTopic(path string) (string, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	item := new(itemReader)
	err = yaml.Unmarshal(body, item)
	if err != nil {
		return "", err
	}
	return item.Item.Writer.Topic, err
}

type leadReader struct {
	Lead lead `yaml:"lead-service"`
}

type lead struct {
	Writer kafkaLeadWriter `yaml:"kafka-writer"`
}
type kafkaLeadWriter struct {
	Topics []string `yaml:"topics"`
}

type itemReader struct {
	Item item `yaml:"item-service"`
}

type item struct {
	Writer kafkaItemWriter `yaml:"kafka-writer"`
}

type kafkaItemWriter struct {
	Topic string `yaml:"item"`
}
