package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func prep(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, s)
}

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		log.Fatalln(err)
	}

	cts := make(map[string]int)

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		txt := prep(sc.Text())
		if len(txt) == 0 {
			continue
		}
		cts[txt]++
	}
	if sc.Err() != nil {
		log.Fatalln(sc.Err())
	}

	for k, v := range cts {
		fmt.Printf("%s: %d\n", k, v)
	}
}
