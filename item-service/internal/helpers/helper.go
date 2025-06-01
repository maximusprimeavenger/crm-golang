package helpers

import (
	"os"

	"errors"

	"gopkg.in/yaml.v3"
)

const path = "/app/config/conf.yaml"

func GetPort() (*int, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	service := new(service)
	err = yaml.Unmarshal(body, service)
	if err != nil {
		return nil, err
	}
	if service.ItemService.Port <= 0 {
		return nil, errors.New("wrong value for item-service port")
	}
	return &service.ItemService.Port, nil
}

type service struct {
	ItemService ItemService `yaml:"item-service"`
}

type ItemService struct {
	Port int `yaml:"grpc-port"`
}
