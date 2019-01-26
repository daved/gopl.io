package main

import "testing"

var (
	s = "th界is is a test"
)

func BenchmarkReverse(b *testing.B) {
	bs := []byte(s)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		reverse(bs)
	}
}

func BenchmarkReverseu(b *testing.B) {
	bs := []byte(s)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		reverseu(bs)
	}
}
