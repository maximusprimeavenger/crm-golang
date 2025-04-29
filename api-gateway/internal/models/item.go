package models

type Item struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Category    *string  `json:"category"`
	Currency    *string  `json:"currency"`
	InStock     *uint    `json:"in_stock"`
	Status      *string  `json:"status"`
}
