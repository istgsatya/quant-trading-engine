package types

import "time"

// Side represents order direction.
type Side int

const (
    Buy Side = iota
    Sell
)

// OrderType defines execution style.
type OrderType int

const (
    Limit OrderType = iota
    Market
)

// PriceLevel represents one orderbook level.
type PriceLevel struct {
    Price float64
    Size  float64
}

// Trade represents an executed trade.
type Trade struct {
    Price     float64
    Size      float64
    Side      Side
    Timestamp time.Time
}

// Order represents an exchange order.
type Order struct {
    ID        string
    Side      Side
    Type      OrderType
    Price     float64
    Quantity  float64
    Timestamp time.Time
}
