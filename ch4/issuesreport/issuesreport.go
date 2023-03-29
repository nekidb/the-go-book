package main

import (
	"text/template"
	"os"
	"tgb/ch4/github"
	"time"
)

// var reportTemplate string = "TotalCount: {{.TotalCount}}"
const reportTemplate = `Total count: {{.TotalCount}}
{{range .Items}}------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.62s"}}
Age: {{.CreatedAt | daysAgo}}
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	args := os.Args[1:]

	issues, err := github.SearchIssues(args)
	if err != nil {
		panic(err)
	}
	issues.Items = issues.Items[:3]

	var report = template.Must(template.New("issues_report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(reportTemplate))
	report.Execute(os.Stdout, issues)
	
}
