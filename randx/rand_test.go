package randx

import (
	"testing"
)

func BenchmarkNumber(b *testing.B) {
	for i := 1; i <= 100000; i++ {
		Number(10)
	}
}

func BenchmarkAny(b *testing.B) {
	for i := 1; i <= 100000; i++ {
		Any(10)
	}
}
