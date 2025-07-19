package generate

import (
	"os"

	"github.com/bushubdegefu/sql-crud/stemplates"
)

func GenerateUtilsApp(data stemplates.ProjectSetting) {
	tmpl := stemplates.LoadTemplate("utilsApp")
	err := os.MkdirAll("utils", os.ModePerm)
	if err != nil {
		panic(err)
	}
	stemplates.WriteTemplateToFileSetting("utils/jwt_utils.go", tmpl, data)
}
