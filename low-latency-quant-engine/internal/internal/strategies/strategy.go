package strategies

import (
    "low-latency-quant-engine/internal/microstructure"
    "low-latency-quant-engine/internal/types"
)

// Strategy defines a pluggable alpha module.
// Implementation details are intentionally omitted.
type Strategy interface {
    Name() string
    OnSignal(metrics microstructure.Metrics)
    OnFill(fill types.Order)
}
