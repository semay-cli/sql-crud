package generate

import (
	"os"

	"github.com/semay-cli/sql-crud/stemplates"
)

func GenerateTracerEchoSetup(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("echoObserve")
	promTmpl := stemplates.LoadTemplate("prometheus")
	tmplMetric := stemplates.LoadTemplate("promyml")
	err := os.MkdirAll("observe", os.ModePerm)
	if err != nil {
		panic(err)
	}
	stemplates.WriteTemplateToFile("observe/tracer.go", tmpl, data)
	stemplates.WriteTemplateToFile("observe/prometheus.go", promTmpl, data)
	stemplates.WriteTemplateToFile("prometheus.yml", tmplMetric, data)
}
