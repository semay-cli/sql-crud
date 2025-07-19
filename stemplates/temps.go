package stemplates

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"text/template"
	"time"
)

//go:embed temps/*.tmpl
var TemplateFS embed.FS

func WriteTemplateToFile(filePath string, tmpl *template.Template, data Data) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(fmt.Errorf("failed to create file %s: %w", filePath, err))
	}
	defer f.Close()

	if err := tmpl.Execute(f, data); err != nil {
		panic(fmt.Errorf("failed to execute template on %s: %w", filePath, err))
	}
}

func WriteTemplateToFileSetting(filePath string, tmpl *template.Template, data ProjectSetting) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(fmt.Errorf("failed to create file %s: %w", filePath, err))
	}
	defer f.Close()

	if err := tmpl.Execute(f, data); err != nil {
		panic(fmt.Errorf("failed to execute template on %s: %w", filePath, err))
	}
}

func WriteTemplateToFileModel(filePath string, tmpl *template.Template, data Model) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(fmt.Errorf("failed to create file %s: %w", filePath, err))
	}
	defer f.Close()

	if err := tmpl.Execute(f, data); err != nil {
		panic(fmt.Errorf("failed to execute template on %s: %w", filePath, err))
	}
}

func LoadTemplate(name string) *template.Template {
	tmplContent, err := TemplateFS.ReadFile("temps/" + name + ".tmpl")
	if err != nil {
		panic(fmt.Errorf("failed to read embedded template: %w", err))
	}
	tmpl, err := template.New(name).Funcs(FuncMap).Parse(string(tmplContent))
	if err != nil {
		panic(fmt.Errorf("failed to parse template: %w", err))
	}
	return tmpl
}

func CommonCMDInit() {
	time.Sleep(2 * time.Second)
	// running go mod tidy finally
	if err := exec.Command("go", "get", "-u", ".").Run(); err != nil {
		fmt.Printf("error go get: %v \n", err)
	}
}

func CommonCMD() {

	time.Sleep(2 * time.Second)
	// running go mod tidy finally
	if err := exec.Command("go", "mod", "tidy").Run(); err != nil {
		fmt.Printf("error tidy: %v \n", err)
	}
	if err := exec.Command("go", "fmt", "./...").Run(); err != nil {
		fmt.Printf("error formating codes: %v \n", err)
	}
}

func CommonModInit(project_module string) {
	// running go mod tidy finally
	if err := exec.Command("go", "mod", "init", project_module).Run(); err != nil {
		fmt.Printf("error: %v \n", err)
	}
}
