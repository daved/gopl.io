package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/daved/gopl.io/exercises/ch04/e04.14/github"
)

func main() {
	var (
		userEVK  = "GHISSUES_USER"
		tokenEVK = "GHISSUES_TOKEN"
	)

	user, err := lookupEnv(userEVK)
	trip(err)
	token, err := lookupEnv(tokenEVK)
	trip(err)

	m := http.NewServeMux()
	m.HandleFunc("/", handleIssuesFunc(user, token))

	trip(http.ListenAndServe(":8042", m))
}

func handleIssuesFunc(user, token string) http.HandlerFunc {
	var tmpl = template.Must(template.New("tmpl").Parse(
		`<h1>{{.TotalCount}} issues</h1>
<table>
  <tr style='text-align: left'>
    <th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
  </tr>
{{range .Items}}  <tr>
    <td>
	  <a href='{{.HTMLURL}}'>{{.Number}}
    </td>
	<td>{{.State}}</td>
	<td>
	  <a href='{{.User.HTMLURL}}'>{{.User.Login}}</a>
	</td>
	<td>
	  <a href='{{.HTMLURL}}'>{{.Title}}</a>
	</td>
  </tr>{{end}}
</table>
`))

	return func(w http.ResponseWriter, r *http.Request) {
		segs := strings.Split(r.URL.Path, "/")
		if len(segs) != 3 {
			http.NotFound(w, r)
			return
		}

		owner := segs[1]
		repo := segs[2]

		if owner == "" || repo == "" {
			http.NotFound(w, r)
			return
		}

		m, err := github.NewAccessMgmt(user, token, owner, repo)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		iss, err := m.SearchIssues(os.Args[1:])
		if err != nil {
			stts := http.StatusInternalServerError
			http.Error(w, http.StatusText(stts), stts)
			return
		}

		if err := tmpl.Execute(w, iss); err != nil {
			stts := http.StatusInternalServerError
			http.Error(w, http.StatusText(stts), stts)
			return
		}
	}
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
