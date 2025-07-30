package manager

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	goFrame = &cobra.Command{
		Use:           "sql-crud",
		Short:         "SQL – command-line tool to aid structure you ECHO backend projects for API with gorm",
		Long:          "ECHO – command-line tool to aid structure you ECHO backend projects for API with gorm",
		Version:       "0.1.6",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func Execute() {
	if err := goFrame.Execute(); err != nil {

		fmt.Println(err)
		os.Exit(1)
	}
}
