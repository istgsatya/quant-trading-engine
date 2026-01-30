package exchange

import "low-latency-quant-engine/internal/types"

// Gateway abstracts exchange connectivity.
type Gateway interface {
    Send(order types.Order) error
    Cancel(orderID string) error
    Replace(orderID string, price float64) error
}
