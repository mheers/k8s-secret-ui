package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// logLevelFlag describes the verbosity of logs
	logLevelFlag string

	// outputFormatFlag can be json, yaml or table
	outputFormatFlag string

	rootCmd = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(versionCmd)
}
