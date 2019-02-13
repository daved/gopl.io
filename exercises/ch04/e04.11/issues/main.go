// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/daved/gopl.io/exercises/ch04/e04.11/github"
)

func main() {
	var (
		userEVK  = "GHISSUES_USER"
		tokenEVK = "GHISSUES_TOKEN"
		ownerEVK = "GHISSUES_OWNER"
		repoEVK  = "GHISSUES_REPO"
	)

	user, err := lookupEnv(userEVK)
	trip(err)
	token, err := lookupEnv(tokenEVK)
	trip(err)
	owner, _ := lookupEnv(ownerEVK) //nolint
	repo, _ := lookupEnv(repoEVK)   //nolint

	m, err := github.NewAccessMgmt(user, token, owner, repo)
	trip(err)

	iss, err := m.SearchIssues(os.Args[1:])
	trip(err)
	printIssues(iss)

	i, err := m.CreateIssue(
		"Second programmatic test",
		"This is the second body. There are many like it, but this one is a second test.",
	)
	trip(err)

	i, err = m.UpdateIssue(i.Number, "", "", github.StateClosed)
	trip(err)

	iss, err = m.SearchIssues(os.Args[1:])
	trip(err)
	printIssues(iss)
}

func trip(err error) {
	if err != nil {
		panic(err)
	}
}

func lookupEnv(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("cannot get envvar %s", key)
	}
	return val, nil
}

func printIssues(iss *github.IssuesSearchResponse) {
	fmt.Printf("%d issues:\n", iss.TotalCount)

	for _, i := range iss.Items {
		fmt.Printf(
			"#%-5d %9.9s %-55.55s %10.10s\n",
			i.Number,
			i.User.Login,
			i.Title,
			age(i.CreatedAt),
		)
	}
}

func age(t time.Time) string {
	var s string
	now := time.Now()

	switch {
	case t.Before(now.Add(time.Hour * 24 * -365)):
		s = ">  1 year"
	case t.Before(now.Add(time.Hour * 24 * -60)):
		s = "> 60 days"
	default:
		s = "< 60 days"
	}

	return s
}
