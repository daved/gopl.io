package main

import (
	"reflect"
	"testing"
)

var (
	ns    = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	shift = 9
)

func TestRotate(t *testing.T) {
	tests := []struct {
		s     []int
		shift int
		want  []int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6}, 5, []int{5, 6, 0, 1, 2, 3, 4}},
		{[]int{0, 1, 2, 3, 4}, 7, []int{2, 3, 4, 0, 1}},
		{[]int{0, 1, 2, 3}, 4, []int{0, 1, 2, 3}},
		{[]int{0, 1, 2, 3}, 3, []int{3, 0, 1, 2}},
	}

	for _, tt := range tests {
		s := make([]int, len(tt.s))
		copy(s, tt.s)

		rotate(s, tt.shift)
		got := s
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("(rotate) shift %v: got %v, want %v", tt.shift, got, tt.want)
		}

		rotatec(tt.s, tt.shift)
		got = tt.s
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("(rotatec) shift %v: got %v, want %v", tt.shift, got, tt.want)
		}
	}
}

func BenchmarkRotatec(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rotatec(ns, shift)
	}
}

func BenchmarkRotate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rotate(ns, shift)
	}
}
