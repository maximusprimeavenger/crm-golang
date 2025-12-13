package helpers

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/badoux/checkmail"
	"github.com/fiveret/crm-golang/internal/models"
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

func isValidEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	return err == nil
}

type lead struct {
	Port *int `yaml:"grpc-port"`
}

type leadStruct struct {
	Lead lead `yaml:"lead-service"`
}

func CheckEmail(email string) error {
	if !isValidEmail(email) {
		return fmt.Errorf("invalid email format")
	}
	return nil
}
func CheckPhone(phone string) error {
	if len(phone) < 9 {
		return fmt.Errorf("phone number is too short")
	}
	for _, char := range phone {
		if !unicode.IsDigit(char) && string(char) != "-" && string(char) != "+" {
			return fmt.Errorf("phone is not valid")
		}
	}
	return nil
}
func CheckCompany(company string) (string, error) {
	for _, char := range company {
		if unicode.IsDigit(char) {
			return "", fmt.Errorf("company should not contain any digits")
		}
	}
	company = strings.ToUpper(company[:1]) + company[1:]
	return company, nil
}
func CheckName(name string) (string, error) {
	for _, char := range name {
		if unicode.IsDigit(char) {
			return "", fmt.Errorf("name should not contain any digits")
		}
	}
	name = strings.ToUpper(name[:1]) + name[1:]
	return name, nil
}

func ValidateNewLead(lead *models.Lead) error {
	err := CheckEmail(lead.Email)
	if err != nil {
		return err
	}
	err = CheckPhone(lead.Phone)
	if err != nil {
		return err
	}
	lead.Name, err = CheckName(lead.Name)
	if err != nil {
		return err
	}
	lead.Company, err = CheckCompany(lead.Company)
	if err != nil {
		return err
	}
	return nil
}
