package main

import "fmt"

func rotate(s []int, n int) {
	n = n % len(s)
	c := make([]int, n)
	copy(c, s[:n])
	copy(s, s[n:])
	copy(s[len(s)-n:], c)
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate(s, 5)
	fmt.Println(s)
}
