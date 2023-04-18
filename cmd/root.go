package cmd

import (
	"fmt"
	"github.com/typomedia/gitti/app"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   app.App.Name,
	Short: app.App.Description,
	Long:  app.App.Explanation,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(app.Logo())
		fmt.Println(app.App.Explanation)

		if cmd.Flags().Changed("version") {
			os.Exit(0)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "V", false, "Show version")
	rootCmd.Flags().BoolP("help", "h", false, "Show help")
}
