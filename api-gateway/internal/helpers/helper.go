package helpers

import (
	"os"

	"gopkg.in/yaml.v3"
)

func GetPort(path string) (*int, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	api := new(service)
	err = yaml.Unmarshal(body, api)
	if err != nil {
		return nil, err
	}
	return &api.ApiGateway.Port, nil
}

type service struct {
	ApiGateway ApiGateway `yaml:"api-gateway"`
}

type ApiGateway struct {
	Port int `yaml:"http-port"`
}
