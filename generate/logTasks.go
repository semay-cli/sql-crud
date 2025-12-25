package generate

import (
	"os"

	"github.com/semay-cli/sql-crud/stemplates"
)

func GenerateTasks(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("tasks")
	err := os.MkdirAll("scheduler", os.ModePerm)
	if err != nil {
		panic(err)
	}
	stemplates.WriteTemplateToFile("scheduler/tasks.go", tmpl, data)
}

func GenerateLogs(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("logs")
	tmplZap := stemplates.LoadTemplate("zaplogger")
	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		panic(err)
	}
	stemplates.WriteTemplateToFile("logs/logfile.go", tmpl, data)
	stemplates.WriteTemplateToFile("logs/zaplogger.go", tmplZap, data)
}
