package generate

import "github.com/semay-cli/sql-crud/stemplates"

func GenerateSSOLogin(data stemplates.ProjectSetting) {
	tmplService := stemplates.LoadTemplate("loginService")
	tmplController := stemplates.LoadTemplate("loginControl")

	stemplates.WriteTemplateToFileSetting("repository/login_repo.go", tmplService, data)
	stemplates.WriteTemplateToFileSetting("controllers/login_controller.go", tmplController, data)
}
