package main

import (
	"gopl/code/ch4/github"
	"log"
	"os"
	"text/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}--------------------------------------------
Number: {{.Number}}
User: 	{{.User.Login}}
Title:	{{.Title | printf "%.64s"}}
Age:	{{.CreatedAt | daysAgo}} days
{{end}}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

// test run args: repo:golang/go is:open json decoder
