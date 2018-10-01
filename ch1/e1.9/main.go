// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	protoPfx = "https://"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, protoPfx[:4]) {
			url = protoPfx + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err) //nolint
			os.Exit(1)
		}
		defer resp.Body.Close() //nolint

		if n, err := io.Copy(os.Stdout, resp.Body); err != nil {
			pfx := ""
			if n > 0 {
				pfx = "\n"
			}

			fmt.Fprintf(os.Stderr, "%sfetch: reading %s: %v\n", pfx, url, err) //nolint
			os.Exit(1)
		}
		fmt.Println()

		fmt.Printf("status: %s\n", resp.Status)
	}
}
