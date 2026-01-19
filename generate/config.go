package generate

import (
	"os"

	"github.com/semay-cli/sql-crud/stemplates"
)

func GenerateConfig(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("config")
	err := os.MkdirAll("configs", os.ModePerm)
	if err != nil {
		panic(err)
	}
	stemplates.WriteTemplateToFile("configs/configs.go", tmpl, data)

	tmplKeycloak := stemplates.LoadTemplate("utilsKeycloak")
	tmplbase64 := stemplates.LoadTemplate("utilsBase64")
	tmplEmail := stemplates.LoadTemplate("utilsEmail")
	tmplClient := stemplates.LoadTemplate("utilsClient")
	tmplCsv := stemplates.LoadTemplate("utilsCsv")

	err = os.MkdirAll("utils", os.ModePerm)
	if err != nil {
		panic(err)
	}
	stemplates.WriteTemplateToFile("utils/keycloak.go", tmplKeycloak, data)
	stemplates.WriteTemplateToFile("utils/base64.go", tmplbase64, data)
	stemplates.WriteTemplateToFile("utils/email.go", tmplEmail, data)
	stemplates.WriteTemplateToFile("utils/apiClient.go", tmplClient, data)
	stemplates.WriteTemplateToFile("utils/csv.go", tmplCsv, data)
}

func GenerateCrons(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("chron")
	err := os.MkdirAll("crons", os.ModePerm)
	if err != nil {
		panic(err)
	}

	stemplates.WriteTemplateToFile("crons/cron.go", tmpl, data)
}

func GenerateConfigEnv(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("env")

	stemplates.WriteTemplateToFile("configs/.env", tmpl, data)
}

func GenerateConfigAppEnv(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("projectEnv")

	stemplates.WriteTemplateToFile("configs/.dev.env", tmpl, data)
	stemplates.WriteTemplateToFile("configs/.prod.env", tmpl, data)
}

func GenerateConfigTestEnv(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("testEnv")
	err := os.MkdirAll("tests", os.ModePerm)
	if err != nil {
		panic(err)
	}
	stemplates.WriteTemplateToFile("tests/.test.env", tmpl, data)
}
