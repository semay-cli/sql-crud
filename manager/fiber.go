package manager

import (
	"fmt"

	"github.com/semay-cli/sql-crud/generate"
	"github.com/semay-cli/sql-crud/stemplates"
	"github.com/spf13/cobra"
)

var (
	fiber = &cobra.Command{
		Use:   "fiber",
		Short: "generate the basic structure file to start app using fiber",
		Long:  `generate the basic structure file to start app using fiber`,
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
				generate.GenerateFiberAppMiddleware(stemplates.RenderData)
				generate.GenerateFiberSetup(stemplates.RenderData)
				if stemplates.ProjectSettings.AuthAppName == appName {
					generate.GenerateUtilsApp(stemplates.ProjectSettings)
					generate.GenerateSSOLogin(stemplates.ProjectSettings)
				}

			} else if globalName {
				generate.GenerateGlobalFiberAppMiddleware(stemplates.RenderData)
				generate.GenerateAppFiberGlobal(stemplates.RenderData)
				runSwagInitForApps()
			} else {
				fmt.Println("No app name specified")
			}
			stemplates.CommonCMD()
		},
	}
)

func init() {
	fiber.Flags().StringP("app", "a", "", "Specify the app name, so that echo app will be generated")
	fiber.Flags().BoolP("global", "g", false, "basic echo app with for global, creates app.go( in manager package) and middleware.go on the main module takes true or false")
	fiber.Flags().StringVarP(&config_file, "config", "c", "config.json", "Specify the data file to load")
	goFrame.AddCommand(fiber)
}
