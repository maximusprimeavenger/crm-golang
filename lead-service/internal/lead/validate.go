package lead

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/fiveret/crm-golang/internal/helpers"
	"github.com/fiveret/crm-golang/internal/models"
)

func CheckEmail(email string) error {
	if !helpers.IsValidEmail(email) {
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

func NewLead(name, email, phone, company string) (*models.Lead, error) {
	lead := &models.Lead{
		Email: email,
		Phone: phone,
	}

	err := CheckEmail(email)
	if err != nil {
		return nil, err
	}
	err = CheckPhone(phone)
	if err != nil {
		return nil, err
	}
	lead.Name, err = CheckName(name)
	if err != nil {
		return nil, err
	}
	lead.Company, err = CheckCompany(company)
	if err != nil {
		return nil, err
	}
	return lead, nil
}
