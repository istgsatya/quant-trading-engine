# System Architecture

This document describes the high-level architecture of the **Low-Latency Quant Trading Engine**, focusing on data flow, execution boundaries, and design rationale.

The system is intentionally modular, enabling independent development, testing, and replacement of components without impacting core latency paths.

---

##  Architectural Goals

- deterministic low-latency execution
- predictable memory behavior
- strict separation between strategy and infrastructure
- graceful degradation during exchange instability
- safe recovery from process crashes
- prevention of multi-instance self-competition

---

## High-Level Flow

Exchange Market Data
↓
WebSocket / REST Feed
↓
Market Data Normalizer
↓
In-Memory Orderbook
↓
Microstructure Analyzer
↓
Strategy Interface
↓
Execution Engine
↓
Risk Manager
↓
Exchange Gateway
↓
Persistence Layer


---

## Market Data Layer

**Responsibilities:**

- maintain persistent WebSocket connections
- detect sequence gaps and feed desynchronization
- automatically recover via REST snapshots
- normalize exchange-specific messages

**Design Notes:**

- WebSocket is always preferred for latency
- REST is used strictly for recovery
- all downstream components consume normalized deltas

---

## Orderbook Engine

The orderbook is maintained entirely in memory and acts as the single source of truth for all trading decisions.

**Capabilities:**

- incremental delta application
- constant-time best bid / ask lookup
- depth aggregation
- snapshot generation for analysis layers

The orderbook is isolated from strategy logic to ensure correctness and performance.

---

## Microstructure Intelligence Layer

This layer extracts real-time market signals directly from orderbook behavior.

Examples of observed features:

- bid–ask imbalance
- liquidity clustering
- depth withdrawal
- aggressive absorption
- abnormal size concentration

All computations are derived from **orderflow**, not indicators.

---

## Strategy Interface

Strategies operate exclusively on:

- orderbook-derived metrics
- microstructure signals
- execution feedback

They do not interact directly with:

- exchange APIs
- network layers
- persistence systems

This ensures strategy code remains portable and testable.

---

## Execution Engine

The execution engine manages the complete lifecycle of orders:

- placement
- cancellation
- repricing
- fill handling

Design goals:

- maintain top-of-book priority
- minimize unnecessary cancels
- avoid unstable liquidity
- react deterministically under latency pressure

Execution logic is fully decoupled from alpha generation.

---

## Risk Management

Risk controls are enforced **before** orders reach the exchange.

Examples include:

- maximum position size
- per-trade loss limits
- volatility-aware throttling
- latency-based kill switches
- daily drawdown protection

The risk layer has authority to block or cancel orders at any time.

---

## Exchange Gateway

The gateway abstracts exchange-specific APIs.

Responsibilities:

- order submission
- cancellation
- replace operations
- acknowledgement handling
- error normalization

This design allows the engine to support multiple exchanges without modifying upstream logic.

---

## Persistence & Recovery

Critical runtime state is persisted continuously:

- open orders
- active positions
- execution metadata

On restart:

- state is reloaded
- orphaned orders are reconciled
- trading resumes safely without manual intervention

This prevents catastrophic failures during crashes or restarts.

---

## Multi-Instance Coordination

When multiple engine instances run simultaneously:

- leader election determines execution authority
- cooldown locks prevent order collisions
- self-trading and over-positioning are avoided

This enables horizontal scaling without internal competition.

---

## Latency Considerations

The system is optimized for:

- minimal heap allocations
- short-lived critical paths
- predictable garbage collection
- bounded blocking behavior

Latency is measured across:

- tick arrival
- decision computation
- order transmission
- exchange acknowledgement

---

## Summary

This engine prioritizes:

- execution quality over prediction
- structure over complexity
- observability over opacity

It is designed as a **real trading infrastructure system**, not a demo bot.

> Markets cannot be predicted —  
> they must be handled efficiently.
