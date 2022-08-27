package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Conor-Fleming/TaskManagerCLI/database"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Use done to mark a task as complete and remove from the list",
	Run: func(cmd *cobra.Command, args []string) {
		done := strings.Join(args, " ")
		taskID, _ := strconv.Atoi(done)
		result, err := database.RemoveTask(int(taskID))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
		fmt.Printf("Done command has been called, task %v: %v will be removed from the list\n", result, placeholder value)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
