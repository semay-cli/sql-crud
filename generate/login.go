package generate

import "github.com/semay-cli/sql-crud/stemplates"

func GenerateSSOLogin(data stemplates.ProjectSetting) {
	tmplService := stemplates.LoadTemplate("loginService")
	tmplRepository := stemplates.LoadTemplate("loginRepository")
	tmplController := stemplates.LoadTemplate("loginControl")

	stemplates.WriteTemplateToFileSetting("repository/login_repo.go", tmplRepository, data)
	stemplates.WriteTemplateToFileSetting("services/login_svc.go", tmplService, data)
	stemplates.WriteTemplateToFileSetting("controllers/login_controller.go", tmplController, data)
}
