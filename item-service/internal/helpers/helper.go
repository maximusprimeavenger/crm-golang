package helpers

import (
	"os"

	"errors"

	"gopkg.in/yaml.v3"
)

const path = "../config/conf.yaml"

func GetPort() (*int, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	service := new(itemService)
	err = yaml.Unmarshal(body, &service)
	if err != nil {
		return nil, err
	}
	if service.port <= 0 {
		return nil, errors.New("wrong value for item-service port")
	}
	return &service.port, nil
}

type itemService struct {
	port int `yaml:"grpc-port"`
}
