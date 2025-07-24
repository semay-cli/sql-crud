package generate

import (
	"github.com/bushubdegefu/sql-crud/stemplates"
)

func GenerateUtilsApp(data stemplates.ProjectSetting) {
	tmpl := stemplates.LoadTemplate("utilsApp")
	tmplHelp := stemplates.LoadTemplate("jwtUtils")

	stemplates.WriteTemplateToFileSetting("services/jwt_utils.go", tmpl, data)
	stemplates.WriteTemplateToFileSetting("services/utils.go", tmplHelp, data)
}
