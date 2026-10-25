1. Atomic swap means switching the reference from the old cache to the new cache in one indivisible operation
2. Lock contention happens when multiple goroutines need the same lock and some must wait because others hold it. This waiting can slow requests down.

# Learnings: Uber’s Translations at Scale

Source: [Uber Engineering Blog](https://www.uber.com/in/en/blog/translations-at-scale/)

## 1. What “translation” means

A stored string for a particular language, such as `hello-es → hola`. The system retrieves already translated text.

## 2. Reduce allocations and GC work

Go’s garbage collector follows references to identify reachable heap objects. Unnecessary allocations and pointer-heavy structures can increase its work.

Uber replaced its older `sync.Map` storage with `map[string]string` plus `RWMutex`, reducing overhead for its workload. This comparison concerns Go versions before 1.24.

**Learning:** fewer allocations can reduce CPU usage without greatly reducing the actual data stored. Choose structures based on access patterns and profiling.

A pointer is not automatically a separate allocation; ordinary maps still use memory and can allocate.

## 3. Atomic pointers and cache refreshes

An **atomic pointer** lets goroutines safely load or replace a shared reference.

| Operation | Purpose |
| --- | --- |
| `Load()` | Read the current pointer. |
| `Store(next)` | Publish a pointer. |
| `Swap(next)` | Publish a pointer and return the previous one. |

During a refresh:

1. Continue serving the old cache.
2. Build the replacement separately.
3. Atomically publish it when complete.
4. Let existing readers finish using their old reference.

This avoids exposing an unfinished cache or locking readers throughout the rebuild.

### Go example

```go
package main

import (
    "fmt"
    "sync/atomic"
)

type Cache struct {
    Data map[string]string
}

func main() {
    var active atomic.Pointer[Cache]
    active.Store(&Cache{Data: map[string]string{"message": "Old"}})

    snapshot := active.Load()
    next := &Cache{Data: map[string]string{"message": "New"}}
    active.Swap(next)

    fmt.Println(snapshot.Data["message"])      // Old
    fmt.Println(active.Load().Data["message"]) // New
}
```

Keep published maps immutable in this example. Atomic operations protect the pointer, not later map mutations.

Load once when multiple lookups must use the same cache version. `Store` also publishes atomically; `Swap` additionally returns the old pointer.

## 4. Refreshes need extra memory

The current cache and its replacement temporarily coexist. Two 500 MB caches need roughly 1 GB, plus buffers and other application memory.

Swapping does not copy the cache or immediately free the old one. GC can reclaim it after no references remain.

**Learning:** plan for peak memory during refresh, not just normal operation.

## 5. Lock contention and sharding

**Lock contention** happens when goroutines compete for a lock and some must wait.

Sharding divides entries into groups with separate locks. Operations on different shards can proceed independently.

An `RWMutex` allows multiple readers or one exclusive writer. Readers may wait for writers.

**Learning:** sharding reduces contention, but popular keys can still overload one shard.

## 6. Keep hot data in memory

Uber’s hybrid design keeps frequently accessed values in RAM and other values on local disk, with lookup metadata in memory.

A lookup checks memory, retrieves a missing value from backing storage, and may cache it.

**Learning:** a small cached portion of the dataset can serve a large portion of requests when access concentrates on popular entries.

## 7. Cache bookkeeping also has a cost

**LRU — Least Recently Used** prefers evicting entries that have gone unused longest.

Tracking exact recency on every access adds shared writes and synchronization.

Uber used time-based recency updates, a shared approximate clock, and preloading of recently accessed entries during refresh.

**Learning:** approximate metadata can reduce bookkeeping; preloading hot entries can reduce misses after refresh.

## 8. Compact metadata saves memory

Hashes, numeric indices, and compact disk offsets reduce metadata size.

- Hash collisions require handling: different keys can share a hash.
- A disk offset identifies where a record begins.
- Aligned offsets can be encoded in larger units, extending their range but potentially adding padding.

**Learning:** smaller representations save memory but introduce correctness checks and complexity.

## 9. Stream large inputs

Process JSON incrementally instead of decoding the entire input into one large collection.

Streaming reduces temporary parsing memory. Retained values and metadata still occupy space.

**Learning:** validate the replacement before publishing it; retain the current cache if construction fails.

## 10. Verify improvements with measurements

Compare CPU, allocations, GC work, memory, lock waiting, cache-hit rate, and latency under normal traffic and during refresh.

**p99:** approximately 99% of observations fall at or below that latency. It is not the maximum.

## Quick Revision: Questions and Answers

**1. Why can fewer allocations improve CPU without greatly reducing stored data?**  
GC manages fewer allocated objects, while the stored contents may remain the same size.

**2. What does an atomic pointer protect?**  
Loading and replacing the shared reference, not mutations of the referenced object.

**3. Why complete a replacement before publishing it?**  
Readers should receive a complete cache rather than a partially built version.

**4. Why can refreshes exhaust memory?**  
The old cache, replacement, and construction buffers may coexist.

**5. When can the old cache be reclaimed?**  
After no references keep it reachable, including references held by existing readers.

**6. How does sharding reduce lock contention?**  
It distributes operations across separate locks instead of one shared lock.

**7. Why can LRU bookkeeping become expensive?**  
Updating recency on every hit adds writes, synchronization, and CPU work.

**8. Why handle hash collisions?**  
Different keys can share a hash; treating them as identical can return incorrect data.

**9. Does offset compression compress the stored values?**  
No. It encodes storage positions more compactly.

**10. Does streaming guarantee constant total memory?**  
No. It reduces temporary input memory; retained data and metadata can still grow.

**11. How do I verify an optimization?**  
Compare relevant profiles and metrics under representative traffic, including startup and refresh.

## References

- [Go GC Guide](https://go.dev/doc/gc-guide)
- [sync/atomic](https://pkg.go.dev/sync/atomic)
- [sync and RWMutex](https://pkg.go.dev/sync)
- [JSON Decoder](https://pkg.go.dev/encoding/json#Decoder)