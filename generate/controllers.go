package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/bushubdegefu/sql-crud/stemplates"
)

func GenerateControllers(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("controllers")

	_ = os.MkdirAll("controllers", os.ModePerm)

	for _, model := range data.Models {
		filePath := fmt.Sprintf("controllers/%s_controllers.go", strings.ToLower(model.Name))
		stemplates.WriteTemplateToFileModel(filePath, tmpl, model)
	}

	// stemplates.WriteTemplateToFile("models/init.go", migrationTmpl, data)

}
