package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "low-latency-quant-engine/internal/execution"
    "low-latency-quant-engine/internal/marketdata"
    "low-latency-quant-engine/internal/microstructure"
    "low-latency-quant-engine/internal/orderbook"
    "low-latency-quant-engine/internal/persistence"
    "low-latency-quant-engine/internal/risk"
    "low-latency-quant-engine/internal/strategies"
)

func main() {
    log.Println("starting low-latency trading engine")

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // --- core components (implementations omitted intentionally)

    var feed marketdata.Feed
    var book orderbook.OrderBook
    var analyzer microstructure.Analyzer
    var strategy strategies.Strategy
    var executor execution.Executor
    var riskManager risk.Manager
    var store persistence.Store

    _ = feed
    _ = book
    _ = analyzer
    _ = strategy
    _ = executor
    _ = riskManager
    _ = store

    // --- graceful shutdown handling

    shutdown := make(chan os.Signal, 1)
    signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

    <-shutdown

    log.Println("shutdown signal received")
    log.Println("flushing state and stopping engine")

    cancel()
    time.Sleep(500 * time.Millisecond)

    log.Println("engine stopped cleanly")
}
