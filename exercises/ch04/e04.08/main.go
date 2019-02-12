package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type countDisplay struct {
	format string
	ct     int
}

func display(cds ...countDisplay) {
	for _, cd := range cds {
		if cd.ct > 0 {
			fmt.Printf(cd.format, cd.ct)
		}
	}
}

func isInvalid(r rune, byteCt int) bool {
	return r == unicode.ReplacementChar && byteCt == 1
}

func eofTrip(err error) bool {
	if err == io.EOF {
		return true
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		os.Exit(1)
	}

	return false
}

func main() {
	rcs := make(map[rune]int)
	var ulens [utf8.UTFMax + 1]int
	var invalid int
	var ltrs, dgts, spcs, pncs int

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if eofTrip(err) {
			break
		}
		if isInvalid(r, n) {
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

	cds := []countDisplay{
		{"\n%d invalid UTF-8 characters\n", invalid},
		{"\n%d letters\n", ltrs},
		{"\n%d digits\n", dgts},
		{"\n%d spaces\n", spcs},
		{"\n%d punctuations\n", pncs},
	}

	display(cds...)
}
