package main

import (
	"net/http"
	"time"
)

func main() {
	var (
		userEVK  = "GHISSUES_USER"
		tokenEVK = "GHISSUES_TOKEN"
		expiry   = time.Minute * 3
	)

	user, err := lookupEnv(userEVK)
	trip(err)
	token, err := lookupEnv(tokenEVK)
	trip(err)

	gh := newGitHub(user, token, expiry)

	m := http.NewServeMux()
	m.HandleFunc("/favicon.ico", http.NotFound)
	segOffset := 1
	m.Handle("/github.com/", wildcardRoutes{
		segOffset: segOffset,
		routes: map[string]http.HandlerFunc{
			"":           redirectAddSegFunc("issues"),
			"issues":     handleIssuesFunc(segOffset, gh),
			"milestones": handleMilestonesFunc(segOffset, gh),
			"users":      handleUsersFunc(segOffset, gh),
		},
	})

	trip(http.ListenAndServe(":8042", m))
}

func trip(err error) {
	if err != nil {
		panic(err)
	}
}
