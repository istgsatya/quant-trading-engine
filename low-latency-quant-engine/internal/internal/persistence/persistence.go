package persistence

import "low-latency-quant-engine/internal/types"

// Store provides crash-safe state recovery.
type Store interface {
    SaveOrders([]types.Order) error
    LoadOrders() ([]types.Order, error)
}
