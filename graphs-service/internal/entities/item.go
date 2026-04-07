package domain

import "time"

type Item struct {
	ID          uint
	Name        string
	Description string
	Price       float64
	Category    string
	InStock     uint32
	Currency    string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
