package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashWhitespace(bs []byte) []byte {
	var c int
	var curr, prev rune
	var size int

	for i := 0; i < len(bs); i += size {
		prev = curr
		curr, size = utf8.DecodeRune(bs[i:])

		if unicode.IsSpace(curr) && unicode.IsSpace(prev) {
			continue
		}

		copy(bs[c:c+size], bs[i:i+size])
		c += size
	}

	return bs[:c:c]
}

func main() {
	bs := []byte("th界is is a  test   .")
	fmt.Println(string(squashWhitespace(bs)))
}
