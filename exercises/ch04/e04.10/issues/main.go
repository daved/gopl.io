// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/daved/gopl.io/exercises/ch04/e04.10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	t := time.Now()
	d := time.Hour * 24
	fs := ageFilters{
		{">  1 year", t.Add(d * -365).After},
		{"> 60 days", t.Add(d * -60).After},
		{"< 60 days", t.Add(d).After},
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf(
			"#%-5d %9.9s %-55.55s %10.10s\n",
			item.Number,
			item.User.Login,
			item.Title,
			fs.age(item.CreatedAt),
		)
	}
}

type ageFilter struct {
	s  string
	fn func(time.Time) bool
}

type ageFilters []ageFilter

func (fs ageFilters) age(t time.Time) string {
	for _, f := range fs {
		if f.fn(t) {
			return f.s
		}
	}
	return "unknown age"
}
