package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/typomedia/gitti/app"
	"github.com/typomedia/gitti/app/cfg"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show the current config",
	Long:  `Show the current config in the console and the path to the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(app.Logo())
		file, config := cfg.Show()
		fmt.Println(config)
		fmt.Printf("\nFile: %s\n", file)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().BoolP("show", "s", false, "Show the current config")
}
