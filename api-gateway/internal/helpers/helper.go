package helpers

import (
	"os"

	"gopkg.in/yaml.v3"
)

const path = "../config/conf.yaml"

func GetPort() (*int, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	api := new(apiGateway)
	err = yaml.Unmarshal(body, &api)
	if err != nil {
		return nil, err
	}
	return &api.port, nil
}

type apiGateway struct {
	port int `yaml:"http-port"`
}
