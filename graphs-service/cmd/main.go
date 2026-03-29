package main

import (
	"context"
	"fmt"
	"graphs-service/internal/domain"
	reader "graphs-service/internal/interfaces/kafka-reader"
	"log/slog"
	"os"

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
	eventsChan := make(chan domain.Event, 100)
	//leadProductChan := make(chan domain.LeadProduct, 100)

	itemReader, leadReader, leadProductReader, err := reader.LoadReaders(itemTopic, leadTopics, address, logger)
	if err != nil {
		logger.Error(fmt.Sprintf("Error creating one of the readers: %v", err))
		os.Exit(1)
	}

	go itemReader.Start(ctx, eventsChan)
	/*go leadReader.Start(ctx, leadChan)
	go leadProductReader.Start(ctx, eventsChan)*/

	//go aggregator.NewAggregator().Run(ctx, eventsChan)
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
