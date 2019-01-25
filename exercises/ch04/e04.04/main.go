package main

import "fmt"

func rotatec(s []int, n int) {
	n = n % len(s)
	c := make([]int, n)
	copy(c, s[:n])
	copy(s, s[n:])
	copy(s[len(s)-n:], c)
}

// ~15% slower than using copy despite "reducing" iterations.
func rotate(s []int, n int) {
	n = n % len(s)
	c := make([]int, n)
	for i := range s {
		if i < n {
			c[i] = s[i]
		}
		if i < len(s)-n {
			s[i] = s[i+n]
		} else {
			s[i] = c[n-(len(s)-i)]
		}
	}
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate(s, 5)
	fmt.Println(s)
}
