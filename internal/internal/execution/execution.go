package execution

import "low-latency-quant-engine/internal/types"

// Executor handles order placement and lifecycle management.
type Executor interface {
    Place(order types.Order) error
    Cancel(orderID string) error
    Replace(orderID string, price float64) error
    ActiveOrders() []types.Order
}
