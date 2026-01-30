package microstructure

import "low-latency-quant-engine/internal/orderbook"

// Analyzer processes live orderbook state
// and extracts microstructure metrics.
type Analyzer interface {
    OnBookUpdate(book orderbook.OrderBook)
    Metrics() Metrics
}

// Metrics represents derived market signals.
// Computation logic is intentionally abstracted.
type Metrics struct {
    Spread         float64
    Imbalance      float64
    LiquidityDelta float64
    WhaleActivity  bool
}
