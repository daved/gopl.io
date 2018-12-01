package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

const (
	b256 = (iota + 2) * 128
	b384
	b512
)

func main() {
	var (
		length = b256
	)

	lengthDesc := fmt.Sprintf("hash length (%d, %d, %d)", b256, b384, b512)
	flag.IntVar(&length, "length", length, lengthDesc)
	flag.Parse()

	input := []byte(flag.Arg(0))
	if len(input) == 0 {
		fmt.Fprintf(os.Stderr, "input is missing\n")
		os.Exit(1)
	}

	switch length {
	case b256:
		fmt.Println(sha256.Sum256(input))
	case b384:
		fmt.Println(sha512.Sum384(input))
	case b512:
		fmt.Println(sha512.Sum512(input))
	default:
		fmt.Fprintf(os.Stderr, "bad %s\n", lengthDesc)
		os.Exit(1)
	}
}
