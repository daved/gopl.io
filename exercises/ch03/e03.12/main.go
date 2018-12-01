package main

import "unicode/utf8"

func main() {}

func isAnagram(a, b string) bool {
	if len(a) != len(b) || a == "" && b == "" {
		return false
	}

	bbs := []byte(b)
	for _, ra := range a {
		rb, s := utf8.DecodeLastRune(bbs)
		if ra != rb {
			return false
		}
		bbs = bbs[:len(bbs)-s]
	}

	return true
}
