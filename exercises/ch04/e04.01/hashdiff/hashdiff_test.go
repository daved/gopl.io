package hashdiff

import "testing"

func toArr(s string) [32]byte {
	if len(s) > 32 {
		panic("string must be 32 or less bytes")
	}

	var a [32]byte

	for i := 0; i < len(a) && i < len(s); i++ {
		a[i] = s[i]
	}

	return a
}

func TestDifferences(t *testing.T) {
	tests := []struct {
		a, b [32]byte
		want int
	}{
		{toArr("aaaaaaaaaa"), toArr("bbbbbbbbbb"), 20},
		{toArr("aaaaa"), toArr("bbbbb"), 10},
		{toArr("aa"), toArr("ff"), 6},
		{toArr(""), toArr(""), 0},
	}

	for _, tt := range tests {
		got := differences(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
