package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(comma("123456789"))
	fmt.Println(comma("12"))
	fmt.Println(comma("98765432123456789"))
}

// goplComma inserts commas in a non-negative decimal integer string.
func goplComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma(s string) string {
	b := &bytes.Buffer{}

	ct := -(utf8.RuneCountInString(s) % 3)
	for _, r := range s {
		if ct%3 == 0 && b.Len() > 0 {
			_ = b.WriteByte(',') //nolint
		}
		_, _ = b.WriteRune(r) //nolint
		ct++
	}

	return b.String()
}
