package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(join(os.Args[1:], " "))
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func join(ss []string, sep string) string {
	s := ss[0]

	for _, v := range ss[1:] {
		s += sep + v
	}

	return s
}
