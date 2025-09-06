package manager

import (
	"fmt"

	"github.com/semay-cli/sql-crud/generate"
	"github.com/semay-cli/sql-crud/stemplates"
	"github.com/spf13/cobra"
)

var (

	// Controllers Cli for generating data models
	controllercli = &cobra.Command{
		Use:   "controller",
		Short: "Generate data models Controllers for the model",
		Long:  `This command generates data models crud Controllers(generates get,getone,create,update,delete and relations if available).`,
		Run:   runControllerCommand,
	}
)

// runModelsCommand handles the execution of the 'models' command
func runControllerCommand(cmd *cobra.Command, args []string) {

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

	// Generate models and migrations
	// if framework == "init" {
	// generate.GenerateModels(stemplates.RenderData)

	// } else {
	generate.GenerateControllers(stemplates.RenderData)
	generate.GenerateControllerInit(stemplates.RenderData)
	stemplates.CommonCMD()
	// }
}

func init() {
	// Register flags for CRUD and Models commands
	controllercli.Flags().StringVarP(&config_file, "config", "c", "config.json", "Specify the data file to load")
	// controllercli.Flags().StringP("frame", "f", "", "frame work to use for building the")
	controllercli.Flags().StringP("app", "a", "", "Set app name, e.g., \"blue-auth\"")

	// Register commands to the root (goFrame)

	goFrame.AddCommand(controllercli)
}
