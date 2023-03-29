package main

import (
	"html/template"
	// "text/template"
	"tgb/ch4/github"
	"os"
)

var htmlReport = template.Must(template.New("issuesList").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style = 'text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
<td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>`))

func main() {
	args := os.Args[1:]

	issues, err := github.SearchIssues(args)
	if err != nil {
		panic(err)
	}

	if err := htmlReport.Execute(os.Stdout, issues); err != nil {
		panic(err)
	}
}
