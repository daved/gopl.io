package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	rcs := make(map[rune]int)
	var ulens [utf8.UTFMax + 1]int
	var invalid int
	var ltrs, dgts, spcs, pncs int

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		rcs[r]++
		ulens[n]++

		switch {
		case unicode.IsLetter(r):
			ltrs++
		case unicode.IsSpace(r):
			spcs++
		case unicode.IsPunct(r):
			pncs++
		case unicode.IsDigit(r):
			dgts++
		}
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range rcs {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range ulens {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	if ltrs > 0 {
		fmt.Printf("\n%d letters\n", ltrs)
	}
	if dgts > 0 {
		fmt.Printf("\n%d digits\n", dgts)
	}
	if spcs > 0 {
		fmt.Printf("\n%d spaces\n", spcs)
	}
	if pncs > 0 {
		fmt.Printf("\n%d punctuations\n", pncs)
	}
}
