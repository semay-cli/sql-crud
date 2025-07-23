package manager

import (
	"fmt"

	"github.com/bushubdegefu/sql-crud/generate"
	"github.com/bushubdegefu/sql-crud/stemplates"
	"github.com/spf13/cobra"
)

var (
	testscli = &cobra.Command{
		Use:   "test",
		Short: "Generate basic coverage test code for Echo Endpoint services",
		Long:  `Generate basic coverage test code for Echo for the generated crud endpoints.`,
		Run: func(cmd *cobra.Command, args []string) {

			appName, _ := cmd.Flags().GetString("app")

			if appName == "" {
				fmt.Println("Error: --app flag is required.")
				return
			}
			handleAppDirectoryAndLoadConfig(appName)
			generate.GenerateConfigTestEnv(stemplates.RenderData)
			generateTests()

		},
	}
)

func generateTests() {
	// stemplates.TestFrameEcho()
	generate.GenerateEchoCoverage(stemplates.RenderData)
}

func init() {
	// // Register flags for the fiber command
	testscli.Flags().StringP("frame", "f", "", "Specify the framework to use (echo or fiber) for the tests")
	testscli.Flags().StringP("app", "a", "", "Specify the application name")
	goFrame.AddCommand(testscli)
}
