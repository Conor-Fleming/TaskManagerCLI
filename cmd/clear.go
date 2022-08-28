package cmd

import (
	"github.com/Conor-Fleming/TaskManagerCLI/database"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Remove all the tasks from the list",
	Run: func(cmd *cobra.Command, args []string) {
		database.Clearlist()
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
