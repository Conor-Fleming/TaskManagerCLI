package task

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI to help you manage tasks",
	Long: `A helpful CLI taks manager built with you in mind.
				task will help with adding, completing, and remembering your tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		//do stuff
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
