package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/bushubdegefu/sql-crud/stemplates"
)

func GenerateModels(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("model")
	// migrationTmpl := stemplates.LoadTemplate("migration")
	helperTmpl := stemplates.LoadTemplate("helperModel")

	_ = os.MkdirAll("models", os.ModePerm)

	for _, model := range data.Models {
		filePath := fmt.Sprintf("models/%s.go", strings.ToLower(model.Name))
		stemplates.WriteTemplateToFileModel(filePath, tmpl, model)
	}

	// stemplates.WriteTemplateToFile("models/init.go", migrationTmpl, data)
	stemplates.WriteTemplateToFile("models/helper.go", helperTmpl, data)
}
