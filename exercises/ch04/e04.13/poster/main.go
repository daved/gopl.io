package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
	if err := run(); err != nil {
		cmd := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "%s: %s\n", cmd, err)
		os.Exit(1)
	}
}

func run() error {
	var (
		envkKey    = "POSTER_APIKEY"
		srchURLFmt = "https://www.omdbapi.com/?apikey=%s&type=movie&s=%s&page=%d"
		storageDir = "./downloads"
		srchPrompt = "movie search"
		dwnlPrompt = "select download"
		pace       = time.Millisecond * 50
	)

	apikey, err := lookupEnv(envkKey)
	if err != nil {
		return err
	}

	for {
		term, err := userQueryTerm(srchPrompt)
		if err != nil {
			return err
		}

		res, err := search(pace, srchURLFmt, apikey, term)
		if err != nil {
			return err
		}

		if err = userQueryDownload(dwnlPrompt, storageDir, res); err != nil {
			return err
		}
	}
}

func lookupEnv(key string) (string, error) {
	v, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("cannot find envvar %q", key)
	}
	if key == "" {
		return "", fmt.Errorf("%q must not be empty", key)
	}
	return v, nil
}

func userQueryTerm(prompt string) (string, error) {
	fmt.Print(prompt + ": ")

	r := bufio.NewReader(os.Stdin)
	txt, err := r.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("user query: %s", err)
	}

	return txt, nil
}

type movieResponse struct {
	Title  string
	Year   string
	ImdbID string `json:"imdbID"`
	Type   string
	Poster string
}

type searchResponse struct {
	Search       []*movieResponse
	TotalResults string `json:"totalResults"`
	Response     string
}

func additionalPages(s *searchResponse) (int, int) {
	// magic# 2 is the first "additional page"
	a := 2
	// magic# 10 is the max objects per page
	inc := 10

	ct, err := strconv.Atoi(s.TotalResults)
	if err != nil {
		return a, 1
	}

	z := ct / inc
	if ct%inc > 0 {
		z++
	}

	return a, z
}

func search(pace time.Duration, urlFmt, apikey, term string) (*searchResponse, error) {
	efmt := "search: %s"

	escTerm := url.PathEscape(term)
	url := fmt.Sprintf(urlFmt, apikey, escTerm, 1)
	var v searchResponse

	res, err := http.Get(url) //nolint
	if err != nil {
		return nil, fmt.Errorf(efmt, err)
	}
	defer res.Body.Close() //nolint

	if !statusSuccess(res.StatusCode) {
		return nil, fmt.Errorf(efmt, "http request: "+res.Status)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return nil, fmt.Errorf(efmt, err)
	}

	start, end := additionalPages(&v)
	for i := start; i <= end; i++ {
		time.Sleep(pace)

		url := fmt.Sprintf(urlFmt, apikey, escTerm, i)
		var vx searchResponse

		res, err := http.Get(url) //nolint
		if err != nil {
			return nil, fmt.Errorf(efmt, err)
		}
		defer res.Body.Close() //nolint

		if !statusSuccess(res.StatusCode) {
			return nil, fmt.Errorf(efmt, "http request: "+res.Status)
		}

		if err = json.NewDecoder(res.Body).Decode(&vx); err != nil {
			return nil, fmt.Errorf(efmt, err)
		}

		v.Search = append(v.Search, vx.Search...)
	}

	return &v, nil
}

func userQueryDownload(prompt, dir string, sres *searchResponse) error {
	efmt := "user download: %s"

	if len(sres.Search) == 0 {
		fmt.Println("no results")
		return nil
	}

	for i, m := range sres.Search {
		fmt.Printf("%3d) %s (%s)\n", i+1, m.Title, m.Year)
	}

	fmt.Print(prompt + ": ")

	var txt string
	_, err := fmt.Scanln(&txt)
	if err != nil {
		return fmt.Errorf(efmt, err)
	}

	i, err := strconv.Atoi(txt)
	if err != nil {
		return fmt.Errorf(efmt, err)
	}
	i--

	m := sres.Search[i]

	if len(m.Poster) < len("http://") {
		fmt.Println("not available")
		return nil
	}

	title := strings.Map(filenameFilter, m.Title)
	filename := path.Join(dir, title) + "_" + m.Year + path.Ext(m.Poster)

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf(efmt, err)
	}
	defer f.Close() //nolint

	res, err := http.Get(m.Poster)
	if err != nil {
		return fmt.Errorf(efmt, err)
	}
	defer res.Body.Close()

	if _, err = io.Copy(f, res.Body); err != nil {
		return fmt.Errorf(efmt, err)
	}

	if err = f.Sync(); err != nil {
		return fmt.Errorf(efmt, err)
	}

	return nil
}

func filenameFilter(r rune) rune {
	switch {
	case unicode.IsSpace(r):
		return '_'
	case unicode.IsLetter(r) || unicode.IsDigit(r):
		return unicode.ToLower(r)
	}

	return -1
}

func statusSuccess(code int) bool {
	return code > 199 && code < 300
}
