package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/semay-cli/sql-crud/stemplates"
)

func GenerateServices(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("services")
	rlnTmpl := stemplates.LoadTemplate("relationServices")

	// create folder if not exists
	if err := os.MkdirAll("services", os.ModePerm); err != nil {
		panic("could not create directory")
	}

	for _, model := range data.Models {
		filePath := fmt.Sprintf("services/%s_crud_service.go", strings.ToLower(model.Name))
		stemplates.WriteTemplateToFileModel(filePath, tmpl, model)
	}

	for _, model := range data.Models {
		for _, relation := range model.Relations {
			filePath := fmt.Sprintf("services/%s_%s_service.go", strings.ToLower(relation.ParentName), strings.ToLower(relation.FieldName))
			stemplates.WriteTemplateToFileRelation(filePath, rlnTmpl, relation)
		}
	}

}
func GenerateServicesInit(data stemplates.Data) {
	inittmpl := stemplates.LoadTemplate("initService")
	if err := os.MkdirAll("services", os.ModePerm); err != nil {
		panic("could not create directory")
	}

	stemplates.WriteTemplateToFile("services/init.go", inittmpl, data)

}
