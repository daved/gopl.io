package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
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
	m.Handle("/", wildcardRoutes{
		"":           redirectAddSegFunc("issues"),
		"issues":     handleIssuesFunc(user, token),
		"milestones": handleMilestonesFunc(user, token),
		"users":      handleUsersFunc(user, token),
	})

	trip(http.ListenAndServe(":8042", m))
}

func lookupEnv(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("cannot get envvar %s", key)
	}
	return val, nil
}

func trip(err error) {
	if err != nil {
		panic(err)
	}
}

type wildcardRoutes map[string]http.HandlerFunc

func (m wildcardRoutes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	owner, repo, route := firstThreeSegments(r.URL.Path)

	fn, ok := m[route]
	if owner == "" || repo == "" || !ok {
		http.NotFound(w, r)
		return
	}

	fn(w, r)
}

func firstThreeSegments(s string) (string, string, string) {
	segs := strings.Split(s, "/")

	var a, b, c string

	if len(segs) > 1 {
		a = segs[1]
	}
	if len(segs) > 2 {
		b = segs[2]
	}
	if len(segs) > 3 {
		c = segs[3]
	}

	return a, b, c
}

func redirectAddSegFunc(seg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, path.Join(r.URL.Path, seg), http.StatusPermanentRedirect)
	}
}

func handleIssuesFunc(user, token string) http.HandlerFunc {
	var tmpl = template.Must(newTemplate().Parse(`
<h1>{{ len . }} issues</h1>
<table>
  <tr style='text-align: left'>
    <th>#</th>
	<th>State</th>
	<th>Age</th>
	<th>User</th>
	<th>Title</th>
	<th>Milestone</th>
  </tr>
  {{ range . }}<tr>
    <td><a href='{{ .HTMLURL }}'>{{ .Number }}</a></td>
	<td>{{ .State }}</td>
	<td>{{ .CreatedAt | age }}</td>
	<td><a href='{{ .User.HTMLURL }}'>{{ .User.Login }}</a></td>
	<td><a href='{{ .HTMLURL }}'>{{ .Title | preview 48 }}</a></td>
	<td>
	  {{ if .Milestone }}
	  <a href='{{ .Milestone.HTMLURL }}'>{{ .Milestone.Title | preview 24 }}</a>
	  {{ end }}
	</td>
  </tr>{{ end }}
</table>
`))

	return func(w http.ResponseWriter, r *http.Request) {
		owner, repo, _ := firstThreeSegments(r.URL.Path)

		m, err := github.NewAccessMgmt(user, token, owner, repo)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		iss, err := m.ReadIssues()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			stts := http.StatusInternalServerError
			http.Error(w, http.StatusText(stts), stts)
			return
		}

		if err := tmpl.Execute(w, iss); err != nil {
			fmt.Fprintln(os.Stderr, err)
			stts := http.StatusInternalServerError
			http.Error(w, http.StatusText(stts), stts)
			return
		}
	}
}

func handleMilestonesFunc(user, token string) http.HandlerFunc {
	var tmpl = template.Must(newTemplate().Parse(`
<h1>{{ len . }} issues</h1>
<table>
  <tr style='text-align: left'>
    <th>#</th>
	<th>State</th>
	<th>Age</th>
	<th>User</th>
	<th>Title</th>
	<th>Milestone</th>
  </tr>
  {{ range . }}<tr>
    <td><a href='{{ .HTMLURL }}'>{{ .Number }}</a></td>
	<td>{{ .State }}</td>
	<td>{{ .CreatedAt | age }}</td>
	<td><a href='{{ .User.HTMLURL }}'>{{ .User.Login }}</a></td>
	<td><a href='{{ .HTMLURL }}'>{{ .Title | preview 48 }}</a></td>
	<td>
	  {{ if .Milestone }}
	  <a href='{{ .Milestone.HTMLURL }}'>{{ .Milestone.Title | preview 24 }}</a>
	  {{ end }}
	</td>
  </tr>{{ end }}
</table>
`))

	return func(w http.ResponseWriter, r *http.Request) {
		owner, repo, _ := firstThreeSegments(r.URL.Path)

		m, err := github.NewAccessMgmt(user, token, owner, repo)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		iss, err := m.ReadIssues()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			stts := http.StatusInternalServerError
			http.Error(w, http.StatusText(stts), stts)
			return
		}

		if err := tmpl.Execute(w, iss); err != nil {
			fmt.Fprintln(os.Stderr, err)
			stts := http.StatusInternalServerError
			http.Error(w, http.StatusText(stts), stts)
			return
		}
	}
}

func handleUsersFunc(user, token string) http.HandlerFunc {
	var tmpl = template.Must(newTemplate().Parse(`
<h1>{{ len . }} issues</h1>
<table>
  <tr style='text-align: left'>
    <th>#</th>
	<th>State</th>
	<th>Age</th>
	<th>User</th>
	<th>Title</th>
	<th>Milestone</th>
  </tr>
  {{ range . }}<tr>
    <td><a href='{{ .HTMLURL }}'>{{ .Number }}</a></td>
	<td>{{ .State }}</td>
	<td>{{ .CreatedAt | age }}</td>
	<td><a href='{{ .User.HTMLURL }}'>{{ .User.Login }}</a></td>
	<td><a href='{{ .HTMLURL }}'>{{ .Title | preview 48 }}</a></td>
	<td>
	  {{ if .Milestone }}
	  <a href='{{ .Milestone.HTMLURL }}'>{{ .Milestone.Title | preview 24 }}</a>
	  {{ end }}
	</td>
  </tr>{{ end }}
</table>
`))

	return func(w http.ResponseWriter, r *http.Request) {
		owner, repo, _ := firstThreeSegments(r.URL.Path)

		m, err := github.NewAccessMgmt(user, token, owner, repo)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		iss, err := m.ReadIssues()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			stts := http.StatusInternalServerError
			http.Error(w, http.StatusText(stts), stts)
			return
		}

		if err := tmpl.Execute(w, iss); err != nil {
			fmt.Fprintln(os.Stderr, err)
			stts := http.StatusInternalServerError
			http.Error(w, http.StatusText(stts), stts)
			return
		}
	}
}

func preview(l int, s string) string {
	if len(s) < l {
		return s
	}
	return s[:l] + "..."
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

func newTemplate() *template.Template {
	return template.New("tmpl").Funcs(
		template.FuncMap{
			"preview": preview,
			"age":     age,
		},
	)
}
