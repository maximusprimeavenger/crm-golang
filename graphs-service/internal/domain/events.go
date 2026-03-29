package domain

type EventType string

const (
	LeadCreated EventType = "lead.created"
	ItemCreated EventType = "item.created"
)

type Event struct {
	EventID   string
	EventType EventType
	Payload   []byte
}

type ItemState struct {
	ID     uint
	Name   string
	Price  float64
	Status string

	SalesByDay   map[string]int
	RevenueByDay map[string]float64
}

type LeadState struct {
	TotalLeadsByDay  map[string]int
	RevenueByItem    map[string]float64
	SalesCountByItem map[string]int
}
