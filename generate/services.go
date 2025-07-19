package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/bushubdegefu/sql-crud/stemplates"
)

func GenerateServices(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("services")

	_ = os.MkdirAll("services", os.ModePerm)

	for _, model := range data.Models {
		filePath := fmt.Sprintf("services/%s_service.go", strings.ToLower(model.Name))
		stemplates.WriteTemplateToFileModel(filePath, tmpl, model)
	}

}
func GenerateServicesInit(data stemplates.Data) {
	inittmpl := stemplates.LoadTemplate("initService")

	stemplates.WriteTemplateToFile("services/init.go", inittmpl, data)

}
