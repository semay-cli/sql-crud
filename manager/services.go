package manager

import (
	"fmt"

	"github.com/semay-cli/sql-crud/generate"
	"github.com/semay-cli/sql-crud/stemplates"
	"github.com/spf13/cobra"
)

var (

	// Services Cli for generating data models
	servicecli = &cobra.Command{
		Use:   "service",
		Short: "Generate data models services for the model",
		Long:  `This command generates data models crud services(generates get,getone,create,update,delete and relations if available).`,
		Run:   runServiceCommand,
	}
)

// runModelsCommand handles the execution of the 'models' command
func runServiceCommand(cmd *cobra.Command, args []string) {
	stemplates.InitProjectJSON()
	// framework, _ := cmd.Flags().GetString("frame")
	appName, _ := cmd.Flags().GetString("app")

	if appName == "" {
		fmt.Println("Error: --app flag is required.")
		return
	}
	// Change to the app's directory and load the config data
	if err := handleAppDirectoryAndLoadConfig(appName); err != nil {
		fmt.Println(err)
		return
	}

	stemplates.RenderData.AppName = appName
	stemplates.RenderData.AuthAppName = stemplates.ProjectSettings.AuthAppName
	stemplates.RenderData.AppName = appName
	stemplates.ProjectSettings.CurrentAppName = appName
	//stemplates.ProjectSettings.AuthAppName = appName

	generate.GenerateUtilsApp(stemplates.ProjectSettings)
	generate.GenerateServicesInit(stemplates.RenderData)
	generate.GenerateServices(stemplates.RenderData)
	stemplates.CommonCMD()

}

func init() {
	// Register flags for CRUD and Models commands
	servicecli.Flags().StringVarP(&config_file, "config", "c", "config.json", "Specify the data file to load")
	// servicecli.Flags().StringP("frame", "f", "", "frame work to use for building the")
	servicecli.Flags().StringP("app", "a", "", "Set app name, e.g., \"blue-auth\"")

	// Register commands to the root (goFrame)

	goFrame.AddCommand(servicecli)
}
