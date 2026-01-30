package risk

import "low-latency-quant-engine/internal/types"

// Manager enforces exposure and loss constraints.
type Manager interface {
    Validate(order types.Order) bool
    OnFill(fill types.Order)
    KillSwitchActive() bool
}
