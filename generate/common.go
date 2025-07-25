package generate

import (
	"os"

	"github.com/semay-cli/sql-crud/stemplates"
)

func GenerateCommon(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("common")
	err := os.MkdirAll("common", os.ModePerm)
	if err != nil {
		panic(err)
	}
	data.SetBackTick()
	stemplates.WriteTemplateToFile("common/common.go", tmpl, data)
}
