package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Conor-Fleming/TaskManagerCLI/database"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove all the tasks from the list",
	Run: func(cmd *cobra.Command, args []string) {
		done := strings.Join(args, " ")
		taskID, _ := strconv.Atoi(done)
		task := database.ViewTask(taskID)
		_, err := database.RemoveTask(int(taskID))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Remove command called, task \"%v\" will be removed.\n", strings.TrimSuffix(task, " - DONE"))
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
