package helpers

import (
	"os"

	"github.com/badoux/checkmail"
	"gopkg.in/yaml.v3"
)

func FindPort(path string) (*int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lead := new(leadStruct)
	err = yaml.Unmarshal(data, &lead)
	if err != nil {
		return nil, err
	}
	return lead.Lead.Port, nil
}

func IsValidEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	return err == nil
}

type lead struct {
	Port *int `yaml:"grpc-port"`
}

type leadStruct struct {
	Lead lead `yaml:"lead-service"`
}
