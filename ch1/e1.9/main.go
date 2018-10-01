// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	nonProto = "http://"
	secProto = "https://"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err) //nolint
		os.Exit(1)
	}
}

func run() error {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, nonProto) && !strings.HasPrefix(url, secProto) {
			url = secProto + url
		}

		resp, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("fetch: %s", err)
		}
		defer resp.Body.Close() //nolint

		if n, err := io.Copy(os.Stdout, resp.Body); err != nil {
			if n > 0 {
				fmt.Println()
			}

			return fmt.Errorf("fetch: reading %s: %s", url, err)
		}
		fmt.Println()

		fmt.Printf("status: %s\n", resp.Status)
	}

	return nil
}
