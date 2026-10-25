# System Design

A practical collection of system design notes, Go implementations, and engineering case studies. The goal is to connect design fundamentals with working code and understand the trade-offs behind software architecture.

## What's inside

- **[Low-level design](LLD/)** — OOP fundamentals, class relationships, design principles, and design patterns.
- **[Design problems](Problems/)** — Applied exercises including LRU Cache, Bloom Filter, Parking Lot, Tic-Tac-Toe, and Stack Overflow.
- **[Concurrency patterns](Concurrency%20Patterns/)** — Go examples of generators, pipelines, fan-in, fan-out, and worker pools.
- **[High-level design](HLD/)** — Architecture notes and engineering case studies on payments, databases, and notification systems.

## Run the examples

Use the Go version specified in the example's `go.mod`. Each Go project has its own module; run it from that directory:

```bash
cd LLD/OOP
go run .
```

Concurrency examples run individually. From the repository root:

```bash
go run "Concurrency Patterns/worker-pool.go"
```

Start with the fundamentals, apply them to a design problem, and explore the architecture notes for a broader perspective.

This repository combines personal implementations with learnings from online engineering resources and grows through continued practice.
