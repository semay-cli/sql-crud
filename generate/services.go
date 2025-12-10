package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/semay-cli/sql-crud/stemplates"
)

func GenerateServices(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("repository")
	rlnTmpl := stemplates.LoadTemplate("relationRepository")
	tmplSvc := stemplates.LoadTemplate("serviceApp")
	rlnTmplSvc := stemplates.LoadTemplate("relationAppService")

	// create folder if not exists
	if err := os.MkdirAll("repository", os.ModePerm); err != nil {
		panic("could not create directory")
	}
	if err := os.MkdirAll("services", os.ModePerm); err != nil {
		panic("could not create directory")
	}

	for _, model := range data.Models {
		filePath := fmt.Sprintf("repository/%s_crud_repo.go", strings.ToLower(model.Name))
		filePathSvc := fmt.Sprintf("services/%s_crud_svc.go", strings.ToLower(model.Name))

		stemplates.WriteTemplateToFileModel(filePath, tmpl, model)
		stemplates.WriteTemplateToFileModel(filePathSvc, tmplSvc, model)
	}

	for _, model := range data.Models {
		for _, relation := range model.Relations {
			filePath := fmt.Sprintf("repository/%s_%s_repo.go", strings.ToLower(relation.ParentName), strings.ToLower(relation.FieldName))
			filePathSvc := fmt.Sprintf("services/%s_%s_svc.go", strings.ToLower(relation.ParentName), strings.ToLower(relation.FieldName))
			relation.AuthAppName = model.AuthAppName
			stemplates.WriteTemplateToFileRelation(filePath, rlnTmpl, relation)
			stemplates.WriteTemplateToFileRelation(filePathSvc, rlnTmplSvc, relation)
		}
	}

}

func GenerateServicesInit(data stemplates.Data) {
	initServicetmpl := stemplates.LoadTemplate("initService")
	initRepo := stemplates.LoadTemplate("initRepository")
	if err := os.MkdirAll("repository", os.ModePerm); err != nil {
		panic("could not create directory")
	}
	if err := os.MkdirAll("services", os.ModePerm); err != nil {
		panic("could not create directory")
	}

	stemplates.WriteTemplateToFile("repository/init.go", initRepo, data)
	stemplates.WriteTemplateToFile("services/init.go", initServicetmpl, data)

}
