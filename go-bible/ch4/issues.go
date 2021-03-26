// @Description: 模版
// @Author: Arvin
// @date: 2021/3/25 5:04 下午
package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

const tmpl = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	// time.Since 函数将时间转换为过去的时间长度(解释：t 相对于当前时间过去了多久)
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").Funcs(
	template.FuncMap{
		"daysAgo": daysAgo,
	}).Parse(tmpl))

func testTmpl() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
