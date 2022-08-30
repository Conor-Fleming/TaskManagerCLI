package cmd

import (
	"fmt"

	"github.com/Conor-Fleming/TaskManagerCLI/database"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "View your list of tasks to be completed",
	Run: func(cmd *cobra.Command, args []string) {
		list, err := database.ViewList()
		if err != nil {
			fmt.Println(err)
		}
		if len(list) < 1 {
			fmt.Println("There are no tasks on the list")
		} else {
			for i, v := range list {
				fmt.Printf("%v)-id(%v)  %s \n", i+1, v.Key, v.Value)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
