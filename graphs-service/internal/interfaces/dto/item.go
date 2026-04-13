package models

import "time"

type Item struct {
	ID          uint      `json:"item_id"`
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	OldPrice    *float64  `json:"old_price"`
	NewPrice    *float64  `json:"new_price"`
	Category    *string   `json:"category"`
	InStock     *uint32   `json:"in_stock"`
	Currency    *string   `json:"currency"`
	Status      *string   `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// item.created
type ItemCreatedPayload struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// item.price_changed
type ItemPriceChangedPayload struct {
	ID        uint      `json:"id"`
	OldPrice  float64   `json:"old_price"`
	NewPrice  float64   `json:"new_price"`
	ChangedAt time.Time `json:"changed_at"`
}

// item.sold
type ItemSoldPayload struct {
	ID       uint      `json:"item_id"`
	Quantity int       `json:"quantity"`
	Price    float64   `json:"price"`
	Total    float64   `json:"total"`
	SoldAt   time.Time `json:"sold_at"`
}
