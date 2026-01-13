package generate

import (
	"os"

	"github.com/semay-cli/sql-crud/stemplates"
)

func GenerateCommon(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("common")
	tmplItem := stemplates.LoadTemplate("commonitem")
	tmplSlice := stemplates.LoadTemplate("commonslice")
	tmplJSONPool := stemplates.LoadTemplate("jsonFiber")
	tmplQueryBuilder := stemplates.LoadTemplate("querybuilder")
	err := os.MkdirAll("common", os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("querybuilder", os.ModePerm)
	if err != nil {
		panic(err)
	}
	data.SetBackTick()
	stemplates.WriteTemplateToFile("common/common.go", tmpl, data)
	stemplates.WriteTemplateToFile("common/common_item.go", tmplItem, data)
	stemplates.WriteTemplateToFile("common/common_slice.go", tmplSlice, data)
	stemplates.WriteTemplateToFile("common/json.go", tmplJSONPool, data)
	stemplates.WriteTemplateToFile("querybuilder/query_builder.go", tmplQueryBuilder, data)
}
