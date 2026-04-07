package domain

type EventType string

const (
	// Lead lifecycle
	LeadCreated   EventType = "lead.created"
	LeadQualified EventType = "lead.qualified"
	LeadConverted EventType = "lead.converted"
	LeadLost      EventType = "lead.lost"

	// Item lifecycle
	ItemCreated      EventType = "item.created"
	ItemPriceChanged EventType = "item.price_changed"
	ItemArchived     EventType = "item.archived"

	// Sales
	ItemSold EventType = "item.sold"

	// Relations (lead ↔ item)
	LeadItemAdded   EventType = "lead-product.added"
	LeadItemRemoved EventType = "lead-product.removed"
)

type Event struct {
	EventID   string
	EventType EventType
	Payload   []byte
}

type ItemAnalytics struct {
	ID     uint
	Name   string
	Price  float64
	Status string

	SalesByDay   map[string]int
	RevenueByDay map[string]float64
	PriceHistory map[string]float64
}

type LeadState struct {
	TotalLeadsByDay  map[string]int
	RevenueByItem    map[string]float64
	SalesCountByItem map[string]int
}
