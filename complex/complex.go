package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

const (
	numElements = 1000000
	numWorkers  = 10
)

var sampleKeysInt []int
var sampleKeysStr []string

func init() {
	rand.Seed(time.Now().UnixNano())
	sampleKeysInt = make([]int, numElements)
	sampleKeysStr = make([]string, numElements)

	for i := 0; i < numElements; i++ {
		sampleKeysInt[i] = rand.Intn(numElements * 10)
		sampleKeysStr[i] = fmt.Sprintf("key-%d", rand.Intn(numElements*10))
	}
}

func benchmarkMapIntInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		for _, key := range sampleKeysInt {
			m[key] = key
		}
	}
}

func benchmarkMapStrInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[string]int)
		for _, key := range sampleKeysStr {
			m[key] = 1
		}
	}
}

func benchmarkMapIntLookup(b *testing.B) {
	m := make(map[int]int)
	for _, key := range sampleKeysInt {
		m[key] = key
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[sampleKeysInt[i%numElements]]
	}
}

func benchmarkMapStrLookup(b *testing.B) {
	m := make(map[string]int)
	for _, key := range sampleKeysStr {
		m[key] = 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[sampleKeysStr[i%numElements]]
	}
}

func benchmarkConcurrentMapAccess(b *testing.B) {
	var wg sync.WaitGroup
	m := make(map[int]int)
	var mu sync.Mutex

	for i := 0; i < numElements; i++ {
		m[sampleKeysInt[i]] = sampleKeysInt[i]
	}

	b.ResetTimer()
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numElements/numWorkers; j++ {
				mu.Lock()
				_ = m[sampleKeysInt[j]]
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
}

func measureMemoryUsage() {
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory Usage: Alloc = %v KB, Sys = %v KB, NumGC = %v\n", m.Alloc/1024, m.Sys/1024, m.NumGC)
}

func main() {
	fmt.Println("Running Go 1.24 Map Benchmarks")

	t := testing.Benchmark(benchmarkMapIntInsert)
	fmt.Printf("Int Map Insert: %v\n", t)
	t = testing.Benchmark(benchmarkMapStrInsert)
	fmt.Printf("String Map Insert: %v\n", t)

	t = testing.Benchmark(benchmarkMapIntLookup)
	fmt.Printf("Int Map Lookup: %v\n", t)
	t = testing.Benchmark(benchmarkMapStrLookup)
	fmt.Printf("String Map Lookup: %v\n", t)

	t = testing.Benchmark(benchmarkConcurrentMapAccess)
	fmt.Printf("Concurrent Map Access: %v\n", t)

	measureMemoryUsage()
}
