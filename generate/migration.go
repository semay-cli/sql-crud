package generate

import (
	"os"

	"github.com/semay-cli/sql-crud/stemplates"
)

func GenerateAppDatabaseMigration(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("migrationApp")
	err := os.MkdirAll("manager", os.ModePerm)
	if err != nil {
		panic(err)
	}
	stemplates.WriteTemplateToFile("manager/migration.go", tmpl, data)
}
