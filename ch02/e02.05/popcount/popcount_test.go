package popcount

import (
	"testing"
)

var (
	x interface{}
	v = uint64(1 + (1 << 1) + (1 << 2) + (1 << 3) + (1 << 4))
)

func BenchmarkClear(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = Clear(v)
	}
	x = n
}

func BenchmarkShift(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = Shift(v)
	}
	x = n
}

func BenchmarkLoop(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = Loop(v)
	}
	x = n
}

func BenchmarkLong(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = Long(v)
	}
	x = n
}
