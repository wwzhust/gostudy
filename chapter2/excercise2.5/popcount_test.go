package popcount_test

import (
	"testing"

	"./popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount64Bits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount64Bits(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountRightmostAnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountRightmostAnd(0x1234567890ABCDEF)
	}
}
