package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/semay-cli/sql-crud/stemplates"
)

func GenerateControllers(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("controllers")
	rlnTmpl := stemplates.LoadTemplate("relationControllers")

	_ = os.MkdirAll("controllers", os.ModePerm)

	for _, model := range data.Models {
		filePath := fmt.Sprintf("controllers/%s_crud_controllers.go", strings.ToLower(model.Name))
		stemplates.WriteTemplateToFileModel(filePath, tmpl, model)
	}

	for _, model := range data.Models {
		for _, relation := range model.Relations {
			filePath := fmt.Sprintf("controllers/%s_%v_controllers.go", strings.ToLower(relation.ParentName), strings.ToLower(relation.FieldName))
			stemplates.WriteTemplateToFileRelation(filePath, rlnTmpl, relation)
		}
	}

}

func GenerateControllerInit(data stemplates.Data) {
	inittmpl := stemplates.LoadTemplate("initControllers")
	if err := os.MkdirAll("controllers", os.ModePerm); err != nil {
		panic("could not create directory")
	}
	tmplConcurent := stemplates.LoadTemplate("concurrency")
	err := os.MkdirAll("concurrency", os.ModePerm)
	if err != nil {
		panic(err)
	}
	stemplates.WriteTemplateToFile("concurrency/executor.go", tmplConcurent, data)
	stemplates.WriteTemplateToFile("controllers/init.go", inittmpl, data)

}
