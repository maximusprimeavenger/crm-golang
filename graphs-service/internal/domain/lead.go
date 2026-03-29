package domain

import "time"

type Lead struct {
	ID               uint
	Name             string
	Email            string
	Phone            string
	Company          string
	Products         []*Product
	Visits           uint
	LastVisit        *time.Time
	TotalSales       float64
	LastPurchaseDays *uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Product struct {
	ID          uint
	Name        string
	Description string
	Price       float64
	Category    string
	Currency    string
	InStock     uint
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
