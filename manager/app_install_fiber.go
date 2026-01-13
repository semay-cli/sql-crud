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
	installauthclifiber = &cobra.Command{
		Use:   "fiber-install",
		Short: "Install Authentication Managment App With the UI",
		Long:  `Install Authentication Managment App With the UI`,
		Run: func(cmd *cobra.Command, args []string) {
			authAppName, _ := cmd.Flags().GetString("app")
			projectName, _ := cmd.Flags().GetString("project")
			userName, _ := cmd.Flags().GetString("user")
			InstallSSOhAppFiber(userName, projectName, authAppName)

			runSwagInitForApps()
			stemplates.CommonCMDInit()
			stemplates.CommonCMD()
		},
	}
)

func InstallSSOhAppFiber(userName, projectName, authAppName string) {
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
	generate.GenerateLogsFiber(stemplates.RenderData)
	generate.GenerateCommon(stemplates.RenderData)
	generate.GenerateDBConn(stemplates.ProjectSettings)
	generate.GenerateCacheService(stemplates.ProjectSettings)
	generate.GenerateTracerFiberSetup(stemplates.RenderData)

	currentDir, _ := os.Getwd()
	handleAppInitialization(authAppName, currentDir, authAppName)
	os.Chdir(currentDir)
	_ = handleAppDirectoryAndLoadConfig(authAppName)

	stemplates.RenderData.AuthAppName = authAppName
	stemplates.RenderData.AppName = authAppName
	stemplates.RenderData.AppNames = stemplates.ProjectSettings.AppNames
	stemplates.RenderData.CurrentAppName = authAppName
	stemplates.ProjectSettings.CurrentAppName = authAppName
	generate.GenerateSQLCModels(stemplates.RenderData)
	generate.GenerateTasks(stemplates.RenderData)
	// generate.GenerateConfigTestEnv(stemplates.RenderData)
	// generate.GenerateEchoCoverage(stemplates.RenderData)
	generate.GenerateFiberSetup(stemplates.RenderData)

	generate.GenerateModelsSQLc(stemplates.RenderData)
	generate.GenerateUtilsApp(stemplates.ProjectSettings)
	generate.GenerateServicesInit(stemplates.RenderData)
	generate.GenerateServicesSQLC(stemplates.RenderData)
	generate.GenerateControllersFiber(stemplates.RenderData)
	generate.GenerateControllerInit(stemplates.RenderData)
	generate.GenerateSSOLoginFiber(stemplates.ProjectSettings)

	stemplates.ProjectSettings.CurrentAppName = authAppName
	stemplates.ProjectSettings.AuthAppName = authAppName
	generate.GenerateFiberAppMiddleware(stemplates.RenderData)
	generate.GenerateFiberSetup(stemplates.RenderData)

	// Go back to root directory
	os.Chdir(currentDir)
	generate.GenerateAppDatabaseMigration(stemplates.RenderData)
	generate.GenerateGlobalFiberAppMiddleware(stemplates.RenderData)
	generate.GenerateAppFiberGlobal(stemplates.RenderData)

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
	installauthclifiber.Flags().StringP("project", "p", "", "Specify the project name using the app flag")
	installauthclifiber.Flags().StringP("app", "a", "", "Specify the app name using the app flag")
	installauthclifiber.Flags().StringP("user", "u", "", "Specify the githubrepo user name using the user flag")
	goFrame.AddCommand(installauthclifiber)
}
