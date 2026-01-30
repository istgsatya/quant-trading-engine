package orderbook

import "low-latency-quant-engine/internal/types"

// OrderBook represents an in-memory reconstructed orderbook.
//
// Implementations are optimized for fast delta application
// and constant-time access to top-of-book levels.
type OrderBook interface {

    // ApplyDelta mutates the book state using
    // incremental exchange updates.
    ApplyDelta(delta Delta)

    // Snapshot returns a read-only view
    // of the current book state.
    Snapshot() Snapshot

    // BestBid returns the highest bid level.
    BestBid() *types.PriceLevel

    // BestAsk returns the lowest ask level.
    BestAsk() *types.PriceLevel

    // Spread returns current bid–ask spread.
    Spread() float64
}

// Delta represents normalized book updates.
type Delta struct {
    Bids []types.PriceLevel
    Asks []types.PriceLevel
}

// Snapshot represents the current visible depth.
type Snapshot struct {
    Bids []types.PriceLevel
    Asks []types.PriceLevel
}

