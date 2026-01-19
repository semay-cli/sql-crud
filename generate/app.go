package generate

import (
	"github.com/semay-cli/sql-crud/stemplates"
)

func GenerateEchoSetup(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("echoSetup")
	stemplates.WriteTemplateToFile("setup.go", tmpl, data)
}

func GenerateFiberSetup(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("fiberSetup")
	stemplates.WriteTemplateToFile("setup.go", tmpl, data)
}

func GenerateEchoAppMiddleware(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("echoAppMiddleware")

	stemplates.WriteTemplateToFile("middleware.go", tmpl, data)
}

func GenerateFiberAppMiddleware(data stemplates.Data) {
	// tmpl := stemplates.LoadTemplate("echoAppMiddleware")

	// stemplates.WriteTemplateToFile("middleware.go", tmpl, data)
}

func GenerateGlobalEchoAppMiddleware(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("globalEchoMiddleware")
	tmplZap := stemplates.LoadTemplate("globalZapMiddleware")
	tmplJSONPool := stemplates.LoadTemplate("jsonPool")

	stemplates.WriteTemplateToFile("manager/middleware.go", tmpl, data)
	stemplates.WriteTemplateToFile("manager/zapmiddleware.go", tmplZap, data)
	stemplates.WriteTemplateToFile("manager/json.go", tmplJSONPool, data)
}

func GenerateAppEchoGlobal(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("globalEchoDig")
	// tmpl := stemplates.LoadTemplate("globalEchoApp")

	data.SetBackTick()
	stemplates.WriteTemplateToFile("manager/app.go", tmpl, data)
}

func GenerateGlobalFiberAppMiddleware(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("globalFiberMiddleware")
	tmplZap := stemplates.LoadTemplate("globalFiberZapMiddleware")
	tmplAdaptiveLimiter := stemplates.LoadTemplate("adaptiveLimiterFiber")

	stemplates.WriteTemplateToFile("manager/adaptive.go", tmplAdaptiveLimiter, data)
	stemplates.WriteTemplateToFile("manager/middleware.go", tmpl, data)
	stemplates.WriteTemplateToFile("manager/zapmiddleware.go", tmplZap, data)

}

func GenerateAppFiberGlobal(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("globalFiberDig")
	// tmpl := stemplates.LoadTemplate("globalEchoApp")

	data.SetBackTick()
	stemplates.WriteTemplateToFile("manager/app.go", tmpl, data)
}
