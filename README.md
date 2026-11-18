# Systems From Scratch

A low-level systems engineering laboratory focused on building core backend infrastructure components from the ground up using **Go's standard library**—with **zero external frameworks or third-party dependencies**.

## Engineering Philosophy

The objective of this repository is to strip away high-level abstractions and implement fundamental network protocols and storage engines directly over raw OS primitives (`net`, `bufio`, `os`, `sync`).

## Roadmap & Modules

| Module | Core Concepts & Primitives | Status |
| :--- | :--- | :--- |
| **01. Raw TCP & HTTP/1.1 Server** | `net.Listen`, `Accept` loop, goroutine-per-connection model, buffered byte reading, manual HTTP/1.1 request parsing, response framing | `In Progress` |
| **02. In-Memory KV Store (Mini-Redis)** | RESP protocol parser, concurrent hash map, sharded mutex locking, active/passive TTL key expiration | `Planned` |
| **03. Append-Only Log Storage Engine** | Durable disk persistence, sequential writes, in-memory byte-offset indexing, crash recovery | `Planned` |

## Benchmarks & Diagnostics

* **Race Condition Verification:** `go run -race main.go`
* **Memory & Throughput Profiling:** `go test -bench=. -benchmem ./...`