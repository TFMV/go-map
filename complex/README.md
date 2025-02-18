# go-map/complex

This one should output something like this:

```bash
☁  complex [main] ⚡  go run complex.go
Running Go 1.24 Map Benchmarks
Int Map Insert:       25          47548105 ns/op
String Map Insert:       10      100332617 ns/op
Int Map Lookup: 45514886                23.88 ns/op
String Map Lookup: 24928891             47.15 ns/op
Concurrent Map Access: 1000000000                0.1568 ns/op
Memory Usage: Alloc = 39311 KB, Sys = 186457 KB, NumGC = 130
```
