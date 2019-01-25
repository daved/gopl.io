package main

import "fmt"

func dedup(s []string) []string {
	var c int

	for i := 1; i < len(s); i++ {
		if s[i] == "" {
			continue
		}

		if s[i-1] == "" {
			s[i-c] = s[i]
			s[i] = ""
			i -= c + 1
			continue
		}

		if s[i] == s[i-1] {
			c++
			s[i] = ""
		}
	}

	lc := len(s) - c
	return s[:lc:lc]
}

func main() {
	s := []string{"test", "test", "this", "out", "out", "now"}
	fmt.Println(dedup(s))
	s = []string{"test", "test", "test", "out", "out"}
	fmt.Println(dedup(s))
}
