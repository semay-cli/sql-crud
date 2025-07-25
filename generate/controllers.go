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
