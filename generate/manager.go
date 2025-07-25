package generate

import (
	"os"

	"github.com/semay-cli/sql-crud/stemplates"
)

func GenerateMainAndManager(data stemplates.Data) {
	tmplMain := stemplates.LoadTemplate("main")
	tmplManager := stemplates.LoadTemplate("manager")
	err := os.MkdirAll("manager", os.ModePerm)
	if err != nil {
		panic(err)
	}
	stemplates.WriteTemplateToFile("main.go", tmplMain, data)
	stemplates.WriteTemplateToFile("manager/manager.go", tmplManager, data)
}
