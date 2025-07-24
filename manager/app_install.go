package manager

import (
	"fmt"
	"os"

	"github.com/bushubdegefu/sql-crud/dist"
	"github.com/bushubdegefu/sql-crud/generate"
	"github.com/bushubdegefu/sql-crud/stemplates"
	"github.com/spf13/cobra"
)

var (
	appinstallcli = &cobra.Command{
		Use:   "ui",
		Short: "Create React SSO Admin UI dist files",
		Long:  `Create React SSO Admin UI dist files`,
		Run: func(cmd *cobra.Command, args []string) {
			stemplates.InitProjectJSON()
			if stemplates.ProjectSettings.AuthAppType == "sso" {
				dist.SSOAdminUI()
			}
		},
	}

	installauthcli = &cobra.Command{
		Use:   "auth",
		Short: "Install Authentication Managment App With the UI",
		Long:  `Install Authentication Managment App With the UI`,
		Run: func(cmd *cobra.Command, args []string) {
			authAppName, _ := cmd.Flags().GetString("app")
			projectName, _ := cmd.Flags().GetString("project")
			userName, _ := cmd.Flags().GetString("user")
			authType, _ := cmd.Flags().GetString("sso")
			if authType == "sso" {
				InstallSSOhApp(userName, projectName, authAppName)
			}

			runSwagInitForApps()
			stemplates.CommonCMDInit()
			stemplates.CommonCMD()
		},
	}
)

func InstallSSOhApp(userName, projectName, authAppName string) {
	if userName == "" {
		fmt.Println("Should Provide Github Username")
		return
	}
	if authAppName == "" {
		authAppName = "blue-admin"
	}

	moduleName := fmt.Sprintf("github.com/%v/%v", userName, projectName)
	stemplates.CommonProjectName(moduleName, projectName, "sso")
	stemplates.CommonModInit(moduleName)
	// Get current working directory
	stemplates.InitProjectJSON()

	stemplates.RenderData.ProjectName = moduleName
	generate.GenerateMainAndManager(stemplates.RenderData)
	generate.GenerateLogs(stemplates.RenderData)
	generate.GenerateCommon(stemplates.RenderData)
	generate.GenerateDBConn(stemplates.ProjectSettings)
	generate.GenerateTracerEchoSetup(stemplates.RenderData)

	currentDir, _ := os.Getwd()
	handleAppInitialization(authAppName, currentDir, authAppName)
	os.Chdir(currentDir)
	_ = handleAppDirectoryAndLoadConfig(authAppName)

	stemplates.RenderData.AuthAppName = stemplates.ProjectSettings.AuthAppName
	stemplates.RenderData.AppName = stemplates.ProjectSettings.AuthAppName
	stemplates.RenderData.AppNames = stemplates.ProjectSettings.AppNames
	stemplates.ProjectSettings.CurrentAppName = authAppName

	generate.GenerateTasks(stemplates.RenderData)
	generate.GenerateConfigTestEnv(stemplates.RenderData)
	generate.GenerateEchoCoverage(stemplates.RenderData)

	generate.GenerateUtilsApp(stemplates.ProjectSettings)
	generate.GenerateModels(stemplates.RenderData)

	generate.GenerateEchoAppMiddleware(stemplates.RenderData)
	generate.GenerateEchoSetup(stemplates.RenderData)

	// Go back to root directory
	os.Chdir(currentDir)
	generate.GenerateAppDatabaseMigration(stemplates.RenderData)
	generate.GenerateGlobalEchoAppMiddleware(stemplates.RenderData)
	generate.GenerateAppEchoGlobal(stemplates.RenderData)

	stemplates.RenderData.ProjectName = stemplates.ProjectSettings.ProjectName
	stemplates.RenderData.AppNames = stemplates.ProjectSettings.AppNames
	generate.GenerateConfig(stemplates.RenderData)
	generate.GenerateConfigEnv(stemplates.RenderData)
	generate.GenerateConfigAppEnv(stemplates.RenderData)

	dist.SSOAdminUI()
	// After all commands are successfully executed

	fmt.Println("App Installed successfully.")
}

func init() {
	installauthcli.Flags().StringP("frame", "f", "", "Specify the framework for the template (echo or fiber)")
	installauthcli.Flags().StringP("project", "p", "", "Specify the project name using the app flag")
	installauthcli.Flags().StringP("app", "a", "", "Specify the app name using the app flag")
	installauthcli.Flags().StringP("user", "u", "", "Specify the githubrepo user name using the user flag")
	installauthcli.Flags().StringP("sso", "s", "", "Specify the authentication app type sso or standalone")
	goFrame.AddCommand(appinstallcli)
	goFrame.AddCommand(installauthcli)
}
