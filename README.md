<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go" />
  <img src="https://img.shields.io/badge/Latency-2–5ms-critical?style=for-the-badge" />
  <img src="https://img.shields.io/badge/Architecture-Low%20Latency-blueviolet?style=for-the-badge" />
  <img src="https://img.shields.io/badge/Platform-Linux-black?style=for-the-badge" />
</p>

<p align="center">
  <h1 align="center">High-Performance Real-Time Trading Infrastructure</h1>
  <p align="center">
    <strong>Focused on latency, execution quality, and market microstructure.</strong>
  </p>
</p>

---

## Overview

This project is a low-latency quantitative trading engine written in **Go**, designed to process live exchange orderbooks and execute trades with **millisecond-level tick-to-decision latency**.

Unlike typical trading bots or academic demos, this system focuses on the infrastructure layer used in professional trading environments.

### Core Design Principles
* **Execution quality** over prediction
* **Structure** over complexity
* **Orderflow** over indicators
* **Deterministic behavior** under load
* **Fail safely**, recover automatically

---

## High-Level Architecture

```text
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

Key Capabilities
Market Data Engine
WebSocket-first ingestion pipeline

REST snapshot recovery

Sequence gap detection

Normalized delta feeds

Orderbook Engine
In-memory reconstruction

Constant-time best bid/ask access

Depth aggregation

Spread computation

Microstructure Intelligence
Bid–ask imbalance calculation

Liquidity clustering

Thin-liquidity zone detection

Abnormal size detection

Adaptive Execution
Limit-order–based execution

Continuous repricing

Queue-priority management

Protection from unstable liquidity

Predictive Exit Control
Profit targeting

Break-even defense

Loss minimization

Forward liquidity awareness

Terminal Dashboard (TUI)
50–100 Hz refresh rate

Zero GC latency spikes

Live orderbook ladder

Execution and system metrics

Fault Tolerance & Coordination
Crash-safe persistence and automatic recovery

Safe restart without trade loss

Distributed leader election

Cooldown locks and self-competition prevention

Perfomance

Metric,Value
Tick to Decision Latency,~2–5 ms
Orderbook Update Sub-millisecond
TUI Refresh Rate 50–100 Hz
Runtime          Linux
Architecture     Single-process in-memory


Visual Architecture Overview
System Flow Diagram
<p align="center"> <img src="docs/images/architecture-flow.png" width="900" alt="Architecture Flow" /> </p>

This diagram illustrates the complete lifecycle of the trading engine — from market data ingestion and normalization through execution handling and safety fallback paths. The focus is on control flow and system resilience rather than trading logic.

Live Runtime Dashboard (TUI)
<p align="center"> <img src="docs/images/runtime-dashboard.png" width="900" alt="Runtime Dashboard" /> </p>

Snapshot of the terminal-based runtime dashboard captured during live engine operation. The interface provides real-time observability into:

Live orderbook ladder and depth movement

Bid–ask pressure and liquidity shifts

Active execution state and strategy mode transitions

Capital allocation tracking

Loop health, latency, and engine status

Structured system logs and alerts

Note: All identifiers, balances, and parameters have been intentionally anonymized. This interface is designed for system observability and debugging, not for presenting trading performance.


Repository Structure

cmd/engine      → Engine bootstrap  
internal/       → Core interfaces  
docs/           → Architecture and system design  
examples/       → Safe mock demonstrations  
notes/          → Design references

Source Code Notice
This repository intentionally omits:

Proprietary strategy logic

Execution heuristics

Thresholds and parameters

Exchange credentials

The goal is to demonstrate engineering depth, not to distribute a trading bot.

Project Purpose
This system was built to:

Deeply understand exchange microstructure

Design low-latency backend systems

Explore execution mechanics

Practice production-grade system design

It reflects real-world engineering constraints, not theoretical trading models.

<p align="center"> <em>This repository focuses on infrastructure engineering.</em>


<em>Markets change. <strong>Execution quality does not.</strong></em> </p>

