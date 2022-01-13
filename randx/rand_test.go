package randx

import (
	"testing"
)

func BenchmarkNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Number(10)
	}
}

func BenchmarkAny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Any(10)
	}
}
