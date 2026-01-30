# Low-Latency Quant Trading Engine

<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go" />
  <img src="https://img.shields.io/badge/Latency-2–5ms-critical?style=for-the-badge" />
  <img src="https://img.shields.io/badge/Architecture-Low%20Latency-blueviolet?style=for-the-badge" />
  <img src="https://img.shields.io/badge/Platform-Linux-black?style=for-the-badge" />
</p>

<p align="center">
  <strong>
    High-performance real-time trading infrastructure focused on latency,
    execution quality, and market microstructure.
  </strong>
</p>

---

## Overview

This repository contains the infrastructure of a low-latency quantitative
trading engine designed for real-time operation under live exchange
conditions.

The project focuses on engineering reliability, execution quality,
and market microstructure behavior rather than trading strategy logic.

It is intended to demonstrate production-grade backend system design,
not to function as a distributable trading bot. 

---

## High-Level Architecture

```
Market Data
↓
Orderbook
↓
Microstructure
↓
Strategy Interface
↓
Execution Engine
↓
Risk Manager
↓
Exchange Gateway
↓
Persistence
```

---

## Key Capabilities

### Market Data Engine

- WebSocket-first ingestion pipeline
- REST snapshot recovery
- sequence gap detection
- normalized delta feeds

### Orderbook Engine

- in-memory reconstruction
- constant-time best bid / ask access
- depth aggregation
- spread computation

### Microstructure Intelligence

- bid–ask imbalance
- liquidity clustering
- thin-liquidity zones
- abnormal size detection

### Adaptive Execution

- limit-order–based execution
- continuous repricing
- queue-priority management
- protection from unstable liquidity

### Predictive Exit Control

- profit targeting
- break-even defense
- loss minimization
- forward liquidity awareness

### Terminal Dashboard (TUI)

- 50–100 Hz refresh rate
- zero GC latency spikes
- live orderbook ladder
- execution and system metrics

### Fault Tolerance

- crash-safe persistence
- automatic recovery
- safe restart without trade loss

### Distributed Coordination

- leader election
- cooldown locks
- self-competition prevention

---

## Performance

| Metric                   | Value                    |
|--------------------------|--------------------------|
| Tick to Decision Latency | ~2–5 ms                  |
| Orderbook Update         | Sub-millisecond          |
| TUI Refresh Rate         | 100+ Hz                  |
| Runtime                  | Any OS                   |
| Architecture             | Single-process in-memory |


Those numbers on metrics are verified via logs and counting tick times in certain loops
---

## Repository Structure

cmd/engine/      → engine bootstrap  
internal/        → core interfaces  
docs/            → architecture and system design  
examples/        → safe mock demonstrations  
notes/           → design references  

---

## Source Code Notice

This repository intentionally omits:

- proprietary strategy logic
- execution heuristics
- thresholds and parameters
- exchange credentials

The goal is to demonstrate engineering depth, not to distribute a trading bot.

---

## Visual Architecture Overview

System Flow Diagram

<p align="center">
  <img src="docs/images/architecture-flow.png" width="900" />
</p>

This diagram illustrates the complete lifecycle of the trading engine —
from market data ingestion and normalization through execution handling
and safety fallback paths.

The focus is on control flow and system resilience rather than trading logic.

---

## Live Runtime Dashboard (TUI)

<p align="center">
  <img src="docs/images/runtime-dashboard.png" width="900" />
</p>

Snapshot of the terminal-based runtime dashboard captured during live
engine operation.

The interface provides real-time observability into:

- live orderbook ladder and depth movement
- bid–ask pressure and liquidity shifts
- active execution state
- strategy mode transitions
- capital allocation tracking
- loop health, latency, and engine status
- structured system logs and alerts

All identifiers, balances, and parameters have been intentionally anonymized.

This interface is designed for system observability and debugging,
not for presenting trading performance.

---

## Project Purpose

This system was built to:

- deeply understand exchange microstructure
- design low-latency backend systems
- explore execution mechanics
- practice production-grade system design

It reflects real-world engineering constraints, not theoretical models.

---

## What This Project Demonstrates

- real-time low-latency backend engineering
- exchange-style WebSocket data pipelines
- orderbook reconstruction from deltas
- market microstructure analysis
- execution-priority optimization
- concurrency-safe Go systems
- terminal-based observability tooling
- crash-safe persistence design
- multi-instance coordination

---

## Final Note

This repository focuses on infrastructure engineering.

Markets change.  
Execution quality does not.
