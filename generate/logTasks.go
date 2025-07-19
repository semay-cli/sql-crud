package generate

import (
	"os"

	"github.com/bushubdegefu/sql-crud/stemplates"
)

func GenerateTasks(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("tasks")
	err := os.MkdirAll("bluetasks", os.ModePerm)
	if err != nil {
		panic(err)
	}
	stemplates.WriteTemplateToFile("bluetasks/tasks.go", tmpl, data)
}

func GenerateLogs(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("logs")
	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		panic(err)
	}
	stemplates.WriteTemplateToFile("logs/logfile.go", tmpl, data)
}
