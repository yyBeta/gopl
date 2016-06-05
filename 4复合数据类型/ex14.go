package main

import (
	"fmt"
	"gopl/4复合数据类型/github"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func getIssues(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path[1:], " ")
	result, err := github.SearchIssues(params)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	}
	if err := issueList.Execute(w, result); err != nil {
		fmt.Fprintf(w, "%v", err)
	}
}

func main() {
	http.HandleFunc("/", getIssues) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// run URL/args
