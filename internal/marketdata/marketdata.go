package marketdata

import "low-latency-quant-engine/internal/types"

// Feed defines a generic market data source.
//
// Concrete implementations may use WebSocket streams,
// REST polling, or hybrid recovery mechanisms.
type Feed interface {
    // Connect initializes the data connection.
    Connect() error

    // Disconnect gracefully shuts down the feed.
    Disconnect()

    // Subscribe begins streaming orderbook data
    // for a specific trading symbol.
    Subscribe(symbol string) error

    // OnDelta registers a callback invoked
    // for every orderbook update received.
    OnDelta(handler func(delta OrderBookDelta))
}

// OrderBookDelta represents incremental book updates.
type OrderBookDelta struct {
    Bids []types.PriceLevel
    Asks []types.PriceLevel

    // Sequence number used to detect gaps
    // and trigger recovery mechanisms.
    Sequence int64
}
