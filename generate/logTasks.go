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
	tmplPools := stemplates.LoadTemplate("pools")
	tmplZap := stemplates.LoadTemplate("zaplogger")
	tmplZapH := stemplates.LoadTemplate("zaphelper")
	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll("pools", os.ModePerm)
	if err != nil {
		panic(err)
	}
	stemplates.WriteTemplateToFile("logs/logfile.go", tmpl, data)
	stemplates.WriteTemplateToFile("logs/zaplogger.go", tmplZap, data)
	stemplates.WriteTemplateToFile("logs/helper.go", tmplZapH, data)
	stemplates.WriteTemplateToFile("pools/pools.go", tmplPools, data)
}

func GenerateLogsFiber(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("logs")
	tmplPools := stemplates.LoadTemplate("pools")
	tmplZap := stemplates.LoadTemplate("zaploggerFiber")
	tmplZapH := stemplates.LoadTemplate("zaphelper")
	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll("pools", os.ModePerm)
	if err != nil {
		panic(err)
	}
	stemplates.WriteTemplateToFile("logs/logfile.go", tmpl, data)
	stemplates.WriteTemplateToFile("logs/zaplogger.go", tmplZap, data)
	stemplates.WriteTemplateToFile("logs/helper.go", tmplZapH, data)
	stemplates.WriteTemplateToFile("pools/pools.go", tmplPools, data)
}
