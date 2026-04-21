package helpers

import (
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(path string) (*Config, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	expanded := os.ExpandEnv(string(body))
	var cfg Config
	if err := yaml.Unmarshal([]byte(expanded), &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

type Config struct {
	ItemService struct {
		Port        int `yaml:"grpc-port"`
		KafkaWriter struct {
			Topic string `yaml:"topic"`
		} `yaml:"kafka-writer"`
	} `yaml:"item-service"`
}
