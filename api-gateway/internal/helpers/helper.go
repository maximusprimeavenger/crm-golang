package helpers

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v2"
)

type yamLogger struct {
	Env string `yaml:"env"`
}

func GetPort(path string) (string, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	api := new(gatewayStruct)
	err = yaml.Unmarshal(body, api)
	if err != nil {
		return "", err
	}
	return os.ExpandEnv(api.Service.Port), err
}

type gatewayStruct struct {
	Service gatewayService `yaml:"api-gateway"`
}

type gatewayService struct {
	Port string `yaml:"http-port"`
}

func LoadLogger(path string) (*slog.Logger, error) {
	env, err := getEnv(path)
	if err != nil {
		return nil, err
	}
	var handler slog.Handler
	switch env {
	case "dev":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	case "test":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case "prod":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn})
	}
	return slog.New(handler), nil
}

func getEnv(path string) (string, error) {
	env := new(yamLogger)
	body, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	err = yaml.Unmarshal(body, env)
	if err != nil {
		return "", err
	}
	return env.Env, nil
}
