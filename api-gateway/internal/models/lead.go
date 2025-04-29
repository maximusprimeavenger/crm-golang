package models

type Lead struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Company  string `json:"company"`
	Products []Product
}

type Product struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Category    *string  `json:"category"`
	Currency    *string  `json:"currency"`
	InStock     *uint    `json:"in_stock"`
	Status      *string  `json:"status"`
	LeadID      *uint    `json:"lead_id"`
}
