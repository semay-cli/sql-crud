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

func GenerateServicesSQLC(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("repository-sqlc")
	tmplTest := stemplates.LoadTemplate("repository-sqlc-test")
	rlnTmpl := stemplates.LoadTemplate("relationRepository-sqlc")
	rlnTmplTest := stemplates.LoadTemplate("relationRepository-sqlc-test")
	tmplSvc := stemplates.LoadTemplate("serviceApp-sqlc")
	rlnTmplSvc := stemplates.LoadTemplate("relationAppService-sqlc")
	queryTmpl := stemplates.LoadTemplate("sqlcdb")
	schemaTmpl := stemplates.LoadTemplate("sqlcschema")
	// create folder if not exists
	if err := os.MkdirAll("sqlc", os.ModePerm); err != nil {
		panic("could not create directory")
	}

	// stemplates.WriteTemplateToFile("sqlc.yaml", tmpl, data)
	stemplates.WriteTemplateToFile("sqlc/db.go", queryTmpl, data)
	stemplates.WriteTemplateToFile("schema.sql", schemaTmpl, data)

	// create folder if not exists
	if err := os.MkdirAll("repository", os.ModePerm); err != nil {
		panic("could not create directory")
	}
	if err := os.MkdirAll("services", os.ModePerm); err != nil {
		panic("could not create directory")
	}

	for _, model := range data.Models {
		filePath := fmt.Sprintf("repository/%s_crud_repo.go", strings.ToLower(model.Name))
		filePathTest := fmt.Sprintf("repository/%s_crud_repo_test.go", strings.ToLower(model.Name))
		filePathSvc := fmt.Sprintf("services/%s_crud_svc.go", strings.ToLower(model.Name))

		stemplates.WriteTemplateToFileModel(filePath, tmpl, model)
		stemplates.WriteTemplateToFileModel(filePathTest, tmplTest, model)
		stemplates.WriteTemplateToFileModel(filePathSvc, tmplSvc, model)
	}

	for _, model := range data.Models {
		for _, relation := range model.Relations {
			var response_model stemplates.Model
			for _, nested_model := range data.Models {
				if relation.FieldName == nested_model.Name {
					response_model = nested_model
				}
			}

			filePath := fmt.Sprintf("repository/%s_%s_repo.go", strings.ToLower(relation.ParentName), strings.ToLower(relation.FieldName))
			filePathTest := fmt.Sprintf("repository/%s_%s_repo_test.go", strings.ToLower(relation.ParentName), strings.ToLower(relation.FieldName))
			filePathSvc := fmt.Sprintf("services/%s_%s_svc.go", strings.ToLower(relation.ParentName), strings.ToLower(relation.FieldName))
			relation.AuthAppName = model.AuthAppName
			relation.ResponseModel = response_model
			// fmt.Println(response_model)
			stemplates.WriteTemplateToFileRelation(filePath, rlnTmpl, relation)
			stemplates.WriteTemplateToFileRelation(filePathTest, rlnTmplTest, relation)
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
