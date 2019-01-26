package main

import "fmt"

func size(b byte) int {
	var c int

	for i := 0; i < 4; i++ {
		if int(b)&(1<<uint(7-i)) == 0 {
			break
		}
		c++
	}

	return c
}

func reverse(bs []byte) {
	rev := func(s []byte) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	rev(bs)

	for i := len(bs) - 1; i >= 0; i-- {
		sz := size(bs[i])
		if sz > 0 {
			rev(bs[i-sz+1 : i+1])
			i -= (sz - 1)
		}
	}
}

func main() {
	bs := []byte("th界is is a test")

	reverse(bs)
	fmt.Println(string(bs))
}
