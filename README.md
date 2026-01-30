<p align="center"> <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go" /> <img src="https://img.shields.io/badge/Latency-2–5ms-critical?style=for-the-badge" /> <img src="https://img.shields.io/badge/Architecture-Low%20Latency-blueviolet?style=for-the-badge" /> <img src="https://img.shields.io/badge/Platform-Linux-black?style=for-the-badge" /> </p> <p align="center"> <b> Real-time trading infrastructure built for speed, stability, and surviving chaotic markets. </b> </p>
🚀 What is this?

This is a low-latency quant trading engine written in Go.

It processes live exchange orderbooks, reacts in milliseconds, and focuses on the part most bots completely ignore:

the infrastructure.

Not indicators.
Not ML predictions.
Not “RSI 14 bro trust me”.

This project is about how trades actually get executed in real systems.

🧠 What this engine actually handles

real-time WebSocket market data

in-memory orderbook reconstruction

microstructure & orderflow analysis

execution priority (top-of-book fights)

latency-aware risk control

crash-safe recovery

multi-instance coordination

No toy logic.
No backtest fantasy land.

🧩 Core mindset

execution > prediction

structure > clever hacks

orderflow > indicators

deterministic behavior under load

if it crashes, it recovers automatically

If it can’t survive chaos, it doesn’t belong in prod.

🧱 High-Level Architecture
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


Everything is isolated.
Everything has boundaries.
Nothing leaks across layers.

⚙️ What’s inside
📡 Market Data

WebSocket-first pipeline

REST snapshot fallback

sequence-gap detection

normalized delta feeds

📘 Orderbook

fully in-memory

constant-time best bid / ask

depth aggregation

real-time spread tracking

🔍 Microstructure

bid–ask imbalance

liquidity clustering

thin-ice zones

aggressive pressure detection

No candles.
No laggy indicators.

⚔️ Execution Engine

limit-order focused

continuous repricing

queue-position awareness

avoids unstable liquidity

This part is about getting filled, not predicting price.

🧠 Exit Logic

profit target handling

break-even defense

controlled loss exits

forward orderbook pressure checks

Exits adapt to liquidity — not vibes.

🖥 Terminal UI (TUI)

runs fully inside terminal

50–100 Hz refresh rate

zero GC latency spikes

live orderbook ladder

engine + execution telemetry

Built for monitoring chaos in real time.

🧯 Safety & Recovery

crash-safe persistence

state reload on restart

no lost positions

no orphaned orders

If the process dies, the engine doesn’t panic.

🧬 Multi-Instance Mode

leader election

cooldown locks

self-trade prevention

no instance fighting itself

Scale without nuking your own fills.

⏱️ Performance (real-world)
Metric	Value
Tick → Decision	~2–5 ms
Orderbook Updates	Sub-ms
UI Refresh	50–100 Hz
Runtime	Linux
Architecture	Single-process, in-memory
🔒 About the source code

This repo does not expose:

strategy math

thresholds

execution heuristics

tuning parameters

exchange credentials

That stuff stays private.

What this repo shows is:

how the system is built, not how the alpha works.

📂 Repo layout
cmd/engine      → engine bootstrap
internal/       → core system interfaces
docs/           → architecture & design
examples/       → safe mock demos
notes/          → design thoughts

🗺️ Architecture Flow
<p align="center"> <img src="docs/images/architecture-flow.png" width="900"/> </p>

High-level control flow of the engine — including fallback paths,
safety checks, and shutdown handling.

The diagram explains how data moves, not how decisions are made.

🖥 Runtime Dashboard
<p align="center"> <img src="docs/images/runtime-dashboard.png" width="900"/> </p>

Live snapshot from the terminal UI.

Shows:

real orderbook movement

execution state transitions

engine health

loop timing

structured logs

All balances and identifiers are blurred on purpose.

This UI exists for observability, not flexing PnL.

🧠 What this project proves

I can build real-time systems under latency pressure

I understand exchange data flow deeply

I can design clean multi-layer architectures

I know how execution actually works

I think in failure modes, not happy paths

This isn’t a “bot”.

It’s a trading engine.

<p align="center"> <b>⚡ fast code. zero vibes. only execution.</b> </p>
