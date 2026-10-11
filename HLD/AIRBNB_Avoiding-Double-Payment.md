# Learning: Avoiding Double Payments in a Distributed Payment System

**Reference:** [Airbnb Engineering – Avoiding Double Payments in a Distributed Payments System](https://medium.com/airbnb-engineering/avoiding-double-payments-in-a-distributed-payments-system-2981f6b070bb)

## Overview

This blog explains how to design a reliable distributed payment system that prevents duplicate payments while handling network failures, retries, concurrent requests, and database inconsistencies.

## Key Learnings

### 1. Idempotency in Distributed Systems

Idempotency ensures that executing the same logical operation multiple times has the same effect as executing it once.

**Request-Level Idempotency**
- Uses a unique idempotency key (typically a UUID) for each logical API request.
- Retrying the same request with the same key does not repeat the payment.
- Useful for handling network timeouts, client retries, and duplicate API submissions.

**Entity-Level Idempotency**
- Uses a deterministic key derived from the business entity and operation.
- Prevents different requests from executing the same logical business operation multiple times.
- Example: Preventing multiple refunds for the same refund operation.

**Key takeaway:** Request-level idempotency protects against duplicate requests, while entity-level idempotency protects against duplicate business operations.

### 2. Eventual Consistency in Distributed Systems

In a distributed system, multiple database replicas may temporarily contain different versions of data. Eventual consistency means replicas converge to a consistent state over time, assuming updates stop and synchronization succeeds.

**Consistency Repair Mechanisms:**

- **Read Repair:** Detects stale replicas during read operations and updates them with the correct data.
- **Write Repair:** Repairs inconsistent or failed writes to bring replicas into agreement.
- **Asynchronous Repair:** Background processes periodically detect and reconcile inconsistent replicas without waiting for application requests.

**Key takeaway:** Eventual consistency improves availability and scalability, but requires mechanisms to detect and repair inconsistencies. Payment correctness must not depend on stale reads.

### 3. Scaling the Master Database for Idempotency

A centralized database storing idempotency records can become a bottleneck as payment traffic increases.

Important considerations:

- Distribute idempotency records across partitions or shards.
- Use appropriate indexing and unique constraints for efficient key lookups.
- Understand the trade-offs between consistency, availability, and scalability.
- Ensure concurrent requests using the same idempotency key cannot independently execute the payment.
- Use reliable synchronization and recovery mechanisms when maintaining distributed copies of idempotency data.

**Key takeaway:** Scaling idempotency storage requires preserving correctness under concurrency, not merely distributing data across multiple databases.

### 4. Separating Database Transactions from Network Calls

External network calls should generally not execute inside long-running database transactions.

Holding a database transaction open during an RPC can cause:

- Long-held database locks.
- Increased transaction contention.
- Reduced throughput and scalability.
- Complicated failure handling when external services time out.

**Three-Phase Architecture:**

**Phase 1: Pre-RPC Transaction**
- Check or reserve the idempotency key.
- Persist the payment operation's initial state.
- Commit the database transaction.

**Phase 2: RPC**
- Call the external payment provider.
- Reuse the appropriate downstream idempotency key.
- Handle timeouts and uncertain outcomes without blindly repeating charges.

**Phase 3: Post-RPC Transaction**
- Persist the payment provider's response.
- Update the payment and idempotency states.
- Commit the final database changes.

**Key takeaway:** Keep database transactions short and separate from network calls while maintaining a recoverable payment state machine.

### 5. Atomic Transactions and Higher-Order Functions

Multiple related database operations should execute atomically: either all changes commit or none do.

The blog discusses using Java lambda expressions to group database statements into transactions.

**Go Alternative:** Anonymous functions and higher-order functions.

```go
func WithTransaction(
    ctx context.Context,
    db *sql.DB,
    fn func(*sql.Tx) error,
) error {
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }

    defer tx.Rollback()

    if err := fn(tx); err != nil {
        return err
    }

    return tx.Commit()
}
```

Usage:

```go
err := WithTransaction(ctx, db, func(tx *sql.Tx) error {
    // Execute multiple database operations atomically.
    return nil
})
```

**Key takeaway:** Go's anonymous functions provide an idiomatic way to encapsulate transactional business logic while keeping transaction management reusable and testable.

### 6. Failure Handling and Recovery

Distributed payment systems must handle failures at different stages:

- **Before RPC:** Database transaction fails before contacting the payment provider.
- **During RPC:** The payment provider may process the charge even if the response times out.
- **After RPC:** The payment succeeds externally, but persisting the result fails.

Important principles:

- Persist sufficient state to recover incomplete operations.
- Retry safely using stable idempotency keys.
- Reconcile uncertain outcomes with the payment provider.
- Avoid assuming that an RPC timeout means the payment failed.

**Key takeaway:** Exactly-once payment effects require coordinated idempotency, durable state, and recovery logic—not simply retrying failed requests.

## Final Takeaways

1. **Idempotency:** Prevent duplicate effects even when requests are retried.
2. **Eventual Consistency:** Understand replica inconsistencies and repair mechanisms.
3. **