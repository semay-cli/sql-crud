package generate

import (
	"github.com/bushubdegefu/sql-crud/stemplates"
)

func GitFrame(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("gitignore")

	stemplates.WriteTemplateToFile(".gitignore", tmpl, data)
}

func DockerFrame(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("docker")

	stemplates.WriteTemplateToFile("Dockerfile", tmpl, data)
}

func HaproxyFrame(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("haproxy")
	stemplates.WriteTemplateToFile("haproxy.cfg", tmpl, data)
}

func AppServiceFrame(data stemplates.Data) {
	tmpl := stemplates.LoadTemplate("service.tmpl")
	stemplates.WriteTemplateToFile("app.service", tmpl, data)
}
