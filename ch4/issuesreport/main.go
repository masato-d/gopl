package main

import (
	"text/template"
	"time"
	"github.com/masato-d/gopl/ch4/github"
	"os"
	"log"
)

const templ = `{{.TotalCount}} issue:
{{range .Items}}---------------------------------------------
Number:	{{.Number}}
User:	{{.User.Login}}
Title:	{{.Title | printf "%.64s"}}
Age:	{{.CreatedAt | daysAgo}} days
{{end}}`

var report = template.Must(template.New("issuelist").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}