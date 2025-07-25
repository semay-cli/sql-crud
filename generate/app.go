package generate

import "github.com/semay-cli/sql-crud/stemplates"

func GenerateEchoSetup(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("echoSetup")

	stemplates.WriteTemplateToFile("setup.go", tmpl, data)
}

func GenerateEchoAppMiddleware(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("echoAppMiddleware")

	stemplates.WriteTemplateToFile("middleware.go", tmpl, data)
}

func GenerateGlobalEchoAppMiddleware(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("globalEchoMiddleware")

	stemplates.WriteTemplateToFile("manager/middleware.go", tmpl, data)
}

func GenerateAppEchoGlobal(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("globalEchoApp")

	data.SetBackTick()
	stemplates.WriteTemplateToFile("manager/app.go", tmpl, data)
}
