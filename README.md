# ⚡ Low-Latency Quant Trading Engine

<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go" />
  <img src="https://img.shields.io/badge/Latency-2–5ms-critical?style=for-the-badge" />
  <img src="https://img.shields.io/badge/Architecture-Low%20Latency-blueviolet?style=for-the-badge" />
  <img src="https://img.shields.io/badge/Platform-Linux-black?style=for-the-badge" />
</p>

<p align="center">
  <b>High-performance real-time trading infrastructure focused on latency, execution quality, and market microstructure.</b>
</p>

---

## 🚀 Overview

This project is a **low-latency quantitative trading engine** written in **Go**, designed to process live exchange orderbooks and execute trades with **millisecond-level tick-to-decision latency**.

Unlike typical trading bots or academic demos, this system focuses on the **infrastructure layer** used in professional trading environments:

- real-time market data ingestion  
- in-memory orderbook reconstruction  
- microstructure-based intelligence  
- execution priority management  
- latency-aware risk control  
- crash-safe recovery  
- multi-instance coordination  

---

## 🧠 Core Design Principles

- **execution quality over prediction**
- **structure over complexity**
- **orderflow over indicators**
- **deterministic behavior under load**
- **fail safely, recover automatically**

---

## 🧩 High-Level Architecture

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

yaml
Copy code

---

## ⚙️ Key Capabilities

### 📡 Market Data Engine
- WebSocket-first ingestion pipeline  
- REST snapshot recovery  
- sequence gap detection  
- normalized delta feeds  

### 📘 Orderbook Engine
- in-memory reconstruction  
- constant-time best bid / ask  
- depth aggregation  
- spread computation  

### 🔍 Microstructure Intelligence
- bid-ask imbalance  
- liquidity clustering  
- thin-liquidity zones  
- abnormal size detection  

### ⚔️ Adaptive Execution
- limit-order based execution  
- continuous repricing  
- queue-priority management  
- protection from unstable depth  

### 🧠 Predictive Exit Control
- profit targeting  
- break-even defense  
- loss minimization  
- forward liquidity awareness  

### 🖥 Terminal Dashboard (TUI)
- 50–100 Hz refresh rate  
- zero GC latency spikes  
- live orderbook ladder  
- execution and system metrics  

### 🧯 Fault Tolerance
- crash-safe persistence  
- automatic recovery  
- safe restart without trade loss  

### 🧬 Distributed Coordination
- leader election  
- cooldown locks  
- self-competition prevention  

---

## ⏱️ Performance

| Metric | Value |
|------|------|
| Tick → Decision Latency | ~2–5 ms |
| Orderbook Update | Sub-millisecond |
| TUI Refresh | 50–100 Hz |
| Runtime | Linux |
| Architecture | Single-process in-memory |

---

## 🔒 Source Code Notice

This repository intentionally omits:

- proprietary strategy logic  
- execution heuristics  
- thresholds and parameters  
- exchange credentials  

The goal is to demonstrate **engineering depth**, not distribute a trading bot.

---

## 📂 Repository Structure

cmd/engine → engine bootstrap
internal/ → core interfaces
docs/ → architecture & system design
examples/ → safe mock demonstrations
notes/ → design references

yaml
Copy code

---

## 🎯 Project Purpose

This system was built to:

- deeply understand exchange microstructure  
- design low-latency backend systems  
- explore execution mechanics  
- practice production-grade system design  

It reflects **real-world engineering constraints**, not theoretical trading models.

---

<p align="center">
<b>⚡ Built for speed. Designed for chaos. Engineered for reality.</b>
</p>
