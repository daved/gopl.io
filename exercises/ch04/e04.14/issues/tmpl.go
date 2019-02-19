package main

import (
	"html/template"
	"time"
)

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
	return template.Must(template.New("tmpl").Funcs(
		template.FuncMap{
			"preview": preview,
			"age":     age,
		},
	).Parse(`{{ define "menu" }}
<a href="./issues">Issues</a>
<a href="./milestones">Milestones</a>
<a href="./users">Users</a>
{{ end }}`))
}
