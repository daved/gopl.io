package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

func lookupEnv(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("cannot get envvar %s", key)
	}
	return val, nil
}

type wildcardRoutes struct {
	segOffset int
	routes    map[string]http.HandlerFunc
}

func (m wildcardRoutes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	owner, repo, route := firstThreeSegments(r.URL.Path, m.segOffset)

	fn, ok := m.routes[route]
	if owner == "" || repo == "" || !ok {
		http.NotFound(w, r)
		return
	}

	dropTrailingSlash(
		fn,
	).ServeHTTP(w, r)
}

func firstThreeSegments(s string, at int) (string, string, string) {
	if s[0] == '/' {
		s = s[1:]
	}
	segs := strings.Split(s, "/")

	var a, b, c string

	i := at
	if len(segs) > i {
		a = segs[i]
	}

	i++
	if len(segs) > i {
		b = segs[i]
	}

	i++
	if len(segs) > i {
		c = segs[i]
	}

	return a, b, c
}

func redirectAddSegFunc(seg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, path.Join(r.URL.Path, seg), http.StatusPermanentRedirect)
	}
}

func dropTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := len(r.URL.Path)
		if r.URL.Path[l-1] == '/' {
			http.Redirect(w, r, r.URL.Path[:l-1], http.StatusPermanentRedirect)
			return
		}

		next.ServeHTTP(w, r)
	})
}
