package models

import "time"

type Lead struct {
	ID               uint       `json:"lead_id"`
	Name             string     `json:"name"`
	Email            string     `json:"email"`
	Phone            string     `json:"phone"`
	Company          string     `json:"company"`
	Products         []*Product `json:"products"`
	Visits           uint       `json:"visits"`
	LastVisit        *time.Time `json:"last_visit"`
	TotalSales       float64    `json:"total_sales"`
	LastPurchaseDays *uint      `json:"last_purchase_days"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

type Product struct {
	ID          uint      `json:"product_id"`
	Name        string    `json:"product_name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category    string    `json:"category"`
	Currency    string    `json:"currency"`
	InStock     uint      `json:"in_stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type LeadCreatedPayload struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
