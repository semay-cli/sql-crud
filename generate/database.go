package generate

import (
	"os"

	"github.com/semay-cli/sql-crud/stemplates"
)

func GenerateDBConn(data stemplates.ProjectSetting) {
	tmpl := stemplates.LoadTemplate("database")
	err := os.MkdirAll("database", os.ModePerm)
	if err != nil {
		panic(err)
	}
	stemplates.WriteTemplateToFileSetting("database/database.go", tmpl, data)
}

func GenerateCacheService(data stemplates.ProjectSetting) {
	tmpl := stemplates.LoadTemplate("cache")
	redisTmpl := stemplates.LoadTemplate("cacheRedis")
	err := os.MkdirAll("cache", os.ModePerm)
	if err != nil {
		panic(err)
	}
	stemplates.WriteTemplateToFileSetting("cache/cache.go", tmpl, data)
	stemplates.WriteTemplateToFileSetting("cache/redis_cache.go", redisTmpl, data)
}
