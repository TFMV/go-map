# go-map

This is a simple benchmark to compare the performance of Go's map implementation between Go 1.23 and Go 1.24. Specifically, 1.24 introduces the Swiss Table, which is a new hash table implementation that is faster for certain use cases.

## Go Map Performance Benchmarks

| Benchmark      | Go 1.23            | Go 1.24            | Improvement |
|---------------|--------------------|--------------------|-------------|
| **Insert**    | 70.60 ms/op (16 iters)  | 47.41 ms/op (25 iters)  | **32.8% faster** |
| **Lookup**    | 54.34 ns/op (25.5M iters) | 38.54 ns/op (39.2M iters) | **29.1% faster** |
| **Delete**    | 25.08 ns/op (40.5M iters) | 15.21 ns/op (77.8M iters) | **39.4% faster** |

✅ **Summary:** Go 1.24's **Swiss Table** implementation shows **significant speed improvements** across **insertions, lookups, and deletions.**
