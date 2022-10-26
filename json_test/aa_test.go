package main

import (
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int31()
	}
}

func BenchmarkRand2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}

// 61.31 ns/op
func BenchmarkRand3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Itoa(rand.Int())
	}
}

// 188.3 ns/op
func BenchmarkRand4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randStr(10)
	}
}
