package generate

import (
	"os"

	"github.com/semay-cli/sql-crud/stemplates"
)

func GenerateUtilsApp(data stemplates.ProjectSetting) {
	tmpl := stemplates.LoadTemplate("utilsApp")
	tmplHelp := stemplates.LoadTemplate("jwtUtils")
	if err := os.MkdirAll("services", os.ModePerm); err != nil {
		panic("could not create directory")
	}

	stemplates.WriteTemplateToFileSetting("services/jwt_utils.go", tmpl, data)
	stemplates.WriteTemplateToFileSetting("services/utils.go", tmplHelp, data)
}
