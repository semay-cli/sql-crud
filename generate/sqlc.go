package generate

import (
	"os"

	"github.com/semay-cli/sql-crud/stemplates"
)

func GenerateSQLCModels(data stemplates.Data) {
	// tmpl := stemplates.LoadTemplate("sqlcyaml")
	queryTmpl := stemplates.LoadTemplate("sqlcdb")
	schemaTmpl := stemplates.LoadTemplate("sqlcschema")
	// create folder if not exists
	if err := os.MkdirAll("sqlc", os.ModePerm); err != nil {
		panic("could not create directory")
	}

	// stemplates.WriteTemplateToFile("sqlc.yaml", tmpl, data)
	stemplates.WriteTemplateToFile("sqlc/db.go", queryTmpl, data)
	stemplates.WriteTemplateToFile("schema.sql", schemaTmpl, data)
}
