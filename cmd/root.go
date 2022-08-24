package cmd

import (
	"github.com/spf13/cobra"
)

var (
	version = "v0.0.1"
	rootCmd = &cobra.Command{
		Use:     "task",
		Version: version,
		Short:   "\nTask is a CLI to help you manage tasks",
	}
)

/*func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}*/
