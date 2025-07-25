package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/semay-cli/sql-crud/stemplates"
)

func GenerateAPIClientJS(data stemplates.Data) {
	indexTmpl := stemplates.LoadTemplate("indexJs")
	authTmpl := stemplates.LoadTemplate("authJs")
	clientTmpl := stemplates.LoadTemplate("clientJs")
	serviceTmpl := stemplates.LoadTemplate("serviceJs")
	err := os.MkdirAll("api", os.ModePerm)
	if err != nil {
		panic(err)
	}

	for _, model := range data.Models {
		filePath := fmt.Sprintf("api/%sService.js", strings.ToLower(model.Name))
		stemplates.WriteTemplateToFileModel(filePath, serviceTmpl, model)
	}
	stemplates.WriteTemplateToFile("api/index.js", indexTmpl, data)
	stemplates.WriteTemplateToFile("api/authService.js", authTmpl, data)
	stemplates.WriteTemplateToFile("api/client.js", clientTmpl, data)
}
