package main

import (
    "fmt"
    "time"

    "low-latency-quant-engine/internal/orderbook"
    "low-latency-quant-engine/internal/types"
)

// This example demonstrates how orderbook deltas
// would be applied in a live trading environment.
//
// All logic is mocked and simplified intentionally.
func main() {
    var book orderbook.OrderBook

    delta := orderbook.Delta{
        Bids: []types.PriceLevel{
            {Price: 100.0, Size: 2.5},
            {Price: 99.5, Size: 4.0},
        },
        Asks: []types.PriceLevel{
            {Price: 100.5, Size: 1.8},
            {Price: 101.0, Size: 3.2},
        },
    }

    fmt.Println("applying mock orderbook delta")
    book.ApplyDelta(delta)

    time.Sleep(200 * time.Millisecond)

    bid := book.BestBid()
    ask := book.BestAsk()

    fmt.Println("best bid:", bid)
    fmt.Println("best ask:", ask)
}
