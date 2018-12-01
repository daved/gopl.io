package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		a, b string
		want bool
	}{
		{"simple valid", "test", "tset", true},
		{"simple invalid", "test", "test", false},
		{"empty junk", "", "junk", false},
		{"junk empty", "junk", "", false},
		{"empty empty", "", "", false},
		{"complex valid", "test-this_out and whatnot", "tontahw dna tuo_siht-tset", true},
		{"complex invalid", "test-this_out and whatnot", "test", false},
	}

	for _, tt := range tests {
		got := isAnagram(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("%s, got %v, want %v", tt.name, got, tt.want)
		}
	}
}
