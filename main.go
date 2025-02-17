package main

import (
	"fmt"
	"math/rand"
	"testing"
)

// Number of elements to test
const numElements = 1_000_000

// Generate test data
func generateTestData(n int) map[int]int {
	data := make(map[int]int, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Intn(n)
	}
	return data
}

// Benchmark map insertions
func BenchmarkMapInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testMap := make(map[int]int)
		for j := 0; j < numElements; j++ {
			testMap[j] = j
		}
	}
}

// Benchmark map lookups
func BenchmarkMapLookup(b *testing.B) {
	testMap := generateTestData(numElements)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = testMap[rand.Intn(numElements)]
	}
}

// Benchmark map deletions
func BenchmarkMapDelete(b *testing.B) {
	testMap := generateTestData(numElements)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		delete(testMap, rand.Intn(numElements))
	}
}

// Run benchmarks and print results
func main() {
	fmt.Println("�� Running benchmarks for Go's map implementation...")

	// Run tests manually for precise output
	testCases := []struct {
		name string
		fn   func(*testing.B)
	}{
		{"Insert", BenchmarkMapInsert},
		{"Lookup", BenchmarkMapLookup},
		{"Delete", BenchmarkMapDelete},
	}

	for _, tc := range testCases {
		fmt.Printf("\n🏁 Benchmarking %s...\n", tc.name)
		result := testing.Benchmark(tc.fn)
		fmt.Println(result)
	}

	fmt.Println("\n✅ Done.")
}
