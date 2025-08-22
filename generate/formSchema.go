package generate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/semay-cli/sql-crud/stemplates"
)

func GenerateFormSchemaZOD(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("pageSchema")
	formTmpl := stemplates.LoadTemplate("pageForm")
	pageTmpl := stemplates.LoadTemplate("pageModel")

	_ = os.MkdirAll("schemas", os.ModePerm)

	for _, model := range data.Models {
		filePath := fmt.Sprintf("schemas/%sSchema.tsx", strings.ToLower(model.Name))
		stemplates.WriteTemplateToFileModel(filePath, tmpl, model)
	}

	originalDir, _ := os.Getwd()
	_ = os.MkdirAll("pages", os.ModePerm)
	newDir := filepath.Join(originalDir, "pages")
	os.Chdir(newDir)
	for _, model := range data.Models {
		_ = os.MkdirAll(fmt.Sprintf("%s", model.LowerName), os.ModePerm)
		filePathPage := fmt.Sprintf("%s/%s.tsx", model.LowerName, model.Name)
		filePath := fmt.Sprintf("%s/%sForm.tsx", model.LowerName, model.Name)
		stemplates.WriteTemplateToFileModel(filePathPage, pageTmpl, model)
		stemplates.WriteTemplateToFileModel(filePath, formTmpl, model)
	}

	// for _, model := range data.Models {
	// 	filePath := fmt.Sprintf("models/%s.go", strings.ToLower(model.Name))
	// 	stemplates.WriteTemplateToFileModel(filePath, tmpl, model)
	// }

}
