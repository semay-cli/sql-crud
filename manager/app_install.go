package manager

import (
	"fmt"
	"os"

	"github.com/semay-cli/sql-crud/dist"
	"github.com/semay-cli/sql-crud/generate"
	"github.com/semay-cli/sql-crud/stemplates"
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
		Use:   "install",
		Short: "Install Authentication Managment App With the UI",
		Long:  `Install Authentication Managment App With the UI`,
		Run: func(cmd *cobra.Command, args []string) {
			authAppName, _ := cmd.Flags().GetString("app")
			projectName, _ := cmd.Flags().GetString("project")
			userName, _ := cmd.Flags().GetString("user")
			InstallSSOhApp(userName, projectName, authAppName)

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
	stemplates.ProjectSettings.AuthAppType = "sso"
	stemplates.ProjectSettings.AuthAppName = authAppName
	stemplates.ProjectSettings.CurrentAppName = authAppName

	stemplates.RenderData.ProjectName = moduleName
	generate.GenerateMainAndManager(stemplates.RenderData)
	generate.GenerateLogs(stemplates.RenderData)
	generate.GenerateCommon(stemplates.RenderData)
	generate.GenerateDBConn(stemplates.ProjectSettings)
	generate.GenerateCacheService(stemplates.ProjectSettings)
	generate.GenerateTracerEchoSetup(stemplates.RenderData)

	currentDir, _ := os.Getwd()
	handleAppInitialization(authAppName, currentDir, authAppName)
	os.Chdir(currentDir)
	_ = handleAppDirectoryAndLoadConfig(authAppName)

	stemplates.RenderData.AuthAppName = authAppName
	stemplates.RenderData.AppName = authAppName
	stemplates.RenderData.AppNames = stemplates.ProjectSettings.AppNames
	stemplates.RenderData.CurrentAppName = authAppName
	stemplates.ProjectSettings.CurrentAppName = authAppName
	// generate.GenerateSQLCModels(stemplates.RenderData)
	generate.GenerateModelsSQLc(stemplates.RenderData)
	generate.GenerateTasks(stemplates.RenderData)
	// generate.GenerateConfigTestEnv(stemplates.RenderData)
	// generate.GenerateEchoCoverage(stemplates.RenderData)
	generate.GenerateEchoSetup(stemplates.RenderData)

	// generate.GenerateModels(stemplates.RenderData)
	generate.GenerateUtilsApp(stemplates.ProjectSettings)
	generate.GenerateServicesInit(stemplates.RenderData)
	generate.GenerateServicesSQLC(stemplates.RenderData)
	generate.GenerateControllers(stemplates.RenderData)
	generate.GenerateControllerInit(stemplates.RenderData)
	generate.GenerateSSOLogin(stemplates.ProjectSettings)

	stemplates.ProjectSettings.CurrentAppName = authAppName
	stemplates.ProjectSettings.AuthAppName = authAppName
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
	generate.GenerateCrons(stemplates.RenderData)
	generate.GenerateConfigEnv(stemplates.RenderData)
	generate.GenerateConfigAppEnv(stemplates.RenderData)

	dist.SSOAdminUI()
	// After all commands are successfully executed

	fmt.Println("App Installed successfully.")
}

func init() {
	installauthcli.Flags().StringP("project", "p", "", "Specify the project name using the app flag")
	installauthcli.Flags().StringP("app", "a", "", "Specify the app name using the app flag")
	installauthcli.Flags().StringP("user", "u", "", "Specify the githubrepo user name using the user flag")
	goFrame.AddCommand(appinstallcli)
	goFrame.AddCommand(installauthcli)
}
