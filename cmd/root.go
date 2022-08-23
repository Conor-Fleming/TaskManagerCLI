package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "v0.0.1"

var rootCmd = &cobra.Command{
	Use:     "task",
	Version: version,
	Short:   "Task is a CLI to help you manage tasks",
	Long: `A helpful CLI taks manager built with you in mind.

Task will help with adding, completing, and remembering your tasks.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
