package helpers

import (
	"os"

	"gopkg.in/yaml.v2"
)

const path = "/app/config/config.yaml"

type writer struct {
	Service service `yaml:"lead-service"`
}

type service struct {
	Writer kafkaTopics `yaml:"kafka-writer"`
}
type kafkaTopics struct {
	Topics []string `yaml:"topics"`
}

func GetTopic(v int) (string, error) {
	writer := new(writer)
	body, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	err = yaml.Unmarshal(body, writer)
	if err != nil {

	}
	return writer.Service.Writer.Topics[v], nil
}
