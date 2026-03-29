package models

import "time"

// TODO сделать json теги
type DTOLead struct {
	ID               uint `json:""`
	Name             string
	Email            string
	Phone            string
	Company          string
	Products         []*DTOProduct
	Visits           uint
	LastVisit        *time.Time
	TotalSales       float64
	LastPurchaseDays *uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type DTOProduct struct {
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
