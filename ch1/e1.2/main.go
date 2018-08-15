package main

import (
	"fmt"
	"os"
)

func main() {
	fi := 1
	for k, v := range os.Args[fi:] {
		fmt.Printf("%d: %s\n", k+fi, v)
	}
}
