package main

import (
    "fmt"

    "low-latency-quant-engine/internal/marketdata"
)

// Demonstrates how a market data feed
// would be consumed by the engine.
//
// No real exchange connections are used.
func main() {
    var feed marketdata.Feed

    feed.OnDelta(func(delta marketdata.OrderBookDelta) {
        fmt.Println("received delta sequence:", delta.Sequence)
    })

    fmt.Println("mock market data feed initialized")
}
