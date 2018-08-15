package main

import (
	"strings"
	"testing"
)

var (
	top interface{}
	ss  = []string{"test", "this", "out", "now", "test", "this", "out", "again"}
	sep = " "
)

func BenchmarkJoin(b *testing.B) {
	var s string

	for n := 0; n < b.N; n++ {
		s = join(ss, sep)
	}

	top = s
}

func BenchmarkStringsJoin(b *testing.B) {
	var s string

	for n := 0; n < b.N; n++ {
		s = strings.Join(ss, sep)
	}

	top = s
}
