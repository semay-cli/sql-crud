package generate

import "github.com/bushubdegefu/sql-crud/stemplates"

func GenerateDjangoAuth(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("django")

	stemplates.WriteTemplateToFile("config.json", tmpl, data)
}

func GenerateSSOAuth(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("sso")

	stemplates.WriteTemplateToFile("config.json", tmpl, data)
}
