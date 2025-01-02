package types

// Money type
type Money int64

// Currency type
type Currency string

// Currency code
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN is card no.
type PAN string

// Card - information about payment card
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

type Payment struct {
	ID     int
	Amount Money
}