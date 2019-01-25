package main

import "fmt"

func dedup(s []string) []string {
	c := 1

	for i := 1; i < len(s); i++ {
		if s[i] != s[c-1] {
			s[c] = s[i]
			c++
		}
	}

	return s[:c:c]
}

func main() {
	s := []string{"test", "test", "this", "out", "out", "now"}
	fmt.Println(dedup(s))
	s = []string{"test", "test", "test", "out", "out"}
	fmt.Println(dedup(s))
}
