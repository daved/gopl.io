package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(comma("123456789"))
	fmt.Println(comma("12"))
	fmt.Println(comma("98765432123456789"))
	fmt.Println(comma("-123456789.123123"))
	fmt.Println(comma("+12"))
	fmt.Println(comma("+98765432123456789.00"))
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

	if len(s) > 0 && s[0:1] == "+" || s[0:1] == "-" {
		_, _ = b.WriteString(s[0:1]) //nolint
	}

	ss := strings.SplitN(s[b.Len():], ".", 2)
	ct := -(utf8.RuneCountInString(ss[0]) % 3)

	for i, r := range ss[0] {
		if ct%3 == 0 && i > 0 {
			_ = b.WriteByte(',') //nolint
		}
		_, _ = b.WriteRune(r) //nolint
		ct++
	}

	if len(ss) > 1 {
		_ = b.WriteByte('.') //nolint
		for _, r := range ss[1] {
			_, _ = b.WriteRune(r) //nolint
		}
	}

	return b.String()
}
