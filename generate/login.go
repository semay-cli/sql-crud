package generate

import "github.com/bushubdegefu/sql-crud/stemplates"

func GenerateSSOLogin(data stemplates.ProjectSetting) {
	tmplService := stemplates.LoadTemplate("loginService")
	tmplController := stemplates.LoadTemplate("loginControl")

	stemplates.WriteTemplateToFileSetting("services/login_service.go", tmplService, data)
	stemplates.WriteTemplateToFileSetting("controllers/login_controller.go", tmplController, data)
}
