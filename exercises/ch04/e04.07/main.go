package main

import (
	"fmt"
	"unicode/utf8"
)

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

func reverseu(bs []byte) {
	rev := func(s []byte) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	var size int
	for i := 0; i < len(bs); i += size {
		_, size = utf8.DecodeRune(bs[i:])
		if size > 1 {
			rev(bs[i : i+size])
		}
	}

	rev(bs)
}

func main() {
	s := "th界is is a test"

	bs := []byte(s)
	reverse(bs) // faster
	fmt.Println(string(bs))

	bs = []byte(s)
	reverse(bs) // more readable
	fmt.Println(string(bs))
}
