package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/bushubdegefu/sql-crud/stemplates"
)

// For Echo coverage testing, we need to generate a test file for each model
func GenerateEchoCoverage(data stemplates.Data) {
	tmplSetting := stemplates.LoadTemplate("echoCoverSetting")
	tmplTests := stemplates.LoadTemplate("echoCover")
	tmplHelperTests := stemplates.LoadTemplate("testHelper")
	err := os.MkdirAll("testsetting", os.ModePerm)
	if err != nil {
		panic(err)
	}

	_ = os.MkdirAll("tests", os.ModePerm)

	for _, model := range data.Models {
		filePath := fmt.Sprintf("tests/%s_crud_controller_test.go", strings.ToLower(model.Name))
		stemplates.WriteTemplateToFileModel(filePath, tmplTests, model)
	}
	stemplates.WriteTemplateToFile("testsetting/settings.go", tmplSetting, data)
	stemplates.WriteTemplateToFile("tests/helper.go", tmplHelperTests, data)
}
