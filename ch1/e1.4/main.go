package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type data struct {
	ct    int
	files []string
}

func main() {
	counts := make(map[string]*data)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	}
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		defer func() { _ = f.Close() }()

		countLines(f, counts)
	}

	for line, d := range counts {
		if d.ct > 1 {
			fs := strings.Join(d.files, " ")
			fmt.Printf("%d\t%s\t%s\n", d.ct, line, fs)
		}
	}
}

func countLines(f *os.File, counts map[string]*data) {
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		t := sc.Text()

		if _, ok := counts[t]; !ok {
			counts[t] = &data{}
		}

		counts[t].ct++
		counts[t].files = appendMissing(counts[t].files, f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}

func appendMissing(ss []string, new string) []string {
	for _, s := range ss {
		if s == new {
			return ss
		}
	}

	return append(ss, new)
}
