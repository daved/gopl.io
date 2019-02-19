package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/daved/gopl.io/exercises/ch04/e04.14/github"
)

func handleIssuesFunc(segOffset int, gh *gitHub) http.HandlerFunc {
	var tmpl = template.Must(newTemplate().Parse(`
<div align="right">{{ template "menu" . }}</div>
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
		owner, repo, _ := firstThreeSegments(r.URL.Path, segOffset)

		is, err := gh.issues(owner, repo)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			stts := http.StatusInternalServerError
			http.Error(w, http.StatusText(stts), stts)
			return
		}

		if err := tmpl.Execute(w, is.data); err != nil {
			fmt.Fprintln(os.Stderr, err)
			stts := http.StatusInternalServerError
			http.Error(w, http.StatusText(stts), stts)
			return
		}
	}
}

func handleMilestonesFunc(segOffset int, gh *gitHub) http.HandlerFunc {
	var tmpl = template.Must(newTemplate().Parse(`
<div align="right">{{ template "menu" . }}</div>
<h1>{{ len . }} milestones</h1>
<table>
  <tr style='text-align: left'>
    <th>#</th>
	<th>State</th>
	<th>Title</th>
	<th>Description</th>
  </tr>
  {{ range . }}<tr>
    <td><a href='{{ .HTMLURL }}'>{{ .Number }}</a></td>
	<td>{{ .State }}</td>
	<td><a href='{{ .HTMLURL }}'>{{ .Title | preview 16 }}</a></td>
	<td>{{ .Desc | preview 48 }}</td>
  </tr>{{ end }}
</table>
`))

	return func(w http.ResponseWriter, r *http.Request) {
		owner, repo, _ := firstThreeSegments(r.URL.Path, segOffset)

		iss, err := gh.issues(owner, repo)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			stts := http.StatusInternalServerError
			http.Error(w, http.StatusText(stts), stts)
			return
		}

		data := make(map[int]*github.Milestone)

		for _, is := range iss.data {
			if is.Milestone != nil {
				data[is.Milestone.Number] = is.Milestone
			}
		}

		if err := tmpl.Execute(w, data); err != nil {
			fmt.Fprintln(os.Stderr, err)
			stts := http.StatusInternalServerError
			http.Error(w, http.StatusText(stts), stts)
			return
		}
	}
}

func handleUsersFunc(segOffset int, gh *gitHub) http.HandlerFunc {
	var tmpl = template.Must(newTemplate().Parse(`
<div align="right">{{ template "menu" . }}</div>
<h1>{{ len . }} users</h1>
<table>
  <tr style='text-align: left'>
	<th>Login</th>
  </tr>
  {{ range . }}<tr>
	<td><a href='{{ .HTMLURL }}'>{{ .Login }}</a></td>
  </tr>{{ end }}
</table>
`))

	return func(w http.ResponseWriter, r *http.Request) {
		owner, repo, _ := firstThreeSegments(r.URL.Path, segOffset)

		iss, err := gh.issues(owner, repo)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			stts := http.StatusInternalServerError
			http.Error(w, http.StatusText(stts), stts)
			return
		}

		data := make(map[string]*github.User)

		for _, is := range iss.data {
			if is.User != nil {
				data[is.User.Login] = is.User
			}
		}

		if err := tmpl.Execute(w, data); err != nil {
			fmt.Fprintln(os.Stderr, err)
			stts := http.StatusInternalServerError
			http.Error(w, http.StatusText(stts), stts)
			return
		}
	}
}
