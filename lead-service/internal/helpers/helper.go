package helpers

import (
	"fmt"
	"os"

	"github.com/badoux/checkmail"
	"github.com/joho/godotenv"
)

func FindPort() (string, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return "", err
	}
	if os.Getenv("PORT") == "" {
		return "", fmt.Errorf("provided no port")
	}
	return os.Getenv("PORT"), nil
}

func IsValidEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	return err == nil
}
