package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

const path = "../.env"

func GetPort() (string, error) {
	err := godotenv.Load(path)
	if err != nil {
		return "", err
	}
	return os.Getenv("PORT_PRODUCT"), nil
}
