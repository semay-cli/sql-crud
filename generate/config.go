package generate

import (
	"os"

	"github.com/bushubdegefu/sql-crud/stemplates"
)

func GenerateConfig(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("config")
	err := os.MkdirAll("configs", os.ModePerm)
	if err != nil {
		panic(err)
	}
	stemplates.WriteTemplateToFile("configs/configs.go", tmpl, data)
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
