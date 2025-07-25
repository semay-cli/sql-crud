package manager

import (
	"fmt"
	"os/exec"

	"github.com/semay-cli/sql-crud/generate"
	"github.com/semay-cli/sql-crud/stemplates"
	"github.com/spf13/cobra"
)

var (
	echocli = &cobra.Command{
		Use:   "echo",
		Short: "generate the basic structure file to start app using echo",
		Long:  `generate the basic structure file to start app using echo`,
		Run: func(cmd *cobra.Command, args []string) {

			// Initialize the project settings
			stemplates.InitProjectJSON()
			stemplates.RenderData.ProjectName = stemplates.ProjectSettings.ProjectName
			stemplates.RenderData.AppNames = stemplates.ProjectSettings.AppNames
			stemplates.RenderData.AuthAppType = stemplates.ProjectSettings.AuthAppType

			appName, _ := cmd.Flags().GetString("app")
			globalName, _ := cmd.Flags().GetBool("global")

			if appName != "" {
				handleAppDirectory(appName)
				if err := stemplates.LoadData(config_file); err != nil {
					fmt.Printf("Error loading data: %v\n", err)
					return
				}

				// generate.GenerateFiberAppMiddleware(stemplates.RenderData)
				stemplates.ProjectSettings.CurrentAppName = appName
				stemplates.RenderData.CurrentAppName = appName
				stemplates.RenderData.AuthAppName = stemplates.ProjectSettings.AuthAppName
				generate.GenerateEchoAppMiddleware(stemplates.RenderData)
				generate.GenerateEchoSetup(stemplates.RenderData)
				if stemplates.ProjectSettings.AuthAppName == appName {
					generate.GenerateUtilsApp(stemplates.ProjectSettings)
					generate.GenerateSSOLogin(stemplates.ProjectSettings)
				}

			} else if globalName {
				generate.GenerateGlobalEchoAppMiddleware(stemplates.RenderData)
				generate.GenerateAppEchoGlobal(stemplates.RenderData)
				runSwagInitForApps()
			} else {
				fmt.Println("No app name specified")
			}
			stemplates.CommonCMD()
		},
	}
)

func runSwagInitForApps() {
	stemplates.InitProjectJSON()
	// swag init --generalInfo setup.go --output  blue-auth/docs --dir=blue-auth,common
	for _, appName := range stemplates.ProjectSettings.AppNames {
		dirArg := fmt.Sprintf("%s,common", appName)
		outputDir := fmt.Sprintf("%s/docs", appName)

		// Prepare the swag init command
		cmd := exec.Command(
			"swag", "init",
			"--generalInfo", "setup.go",
			"--output", outputDir,
			"--dir", dirArg,
		)

		// Run the command and handle errors
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error generating swagger for app '%s': %v\n", appName, err)
		} else {
			fmt.Printf("Swagger generated for app '%s'\n", appName)
		}
	}
}

func init() {
	echocli.Flags().StringP("app", "a", "", "Specify the app name, so that echo app will be generated")
	echocli.Flags().BoolP("global", "g", false, "basic echo app with for global, creates app.go( in manager package) and middleware.go on the main module takes true or false")
	echocli.Flags().StringVarP(&config_file, "config", "c", "config.json", "Specify the data file to load")
	goFrame.AddCommand(echocli)
}
