package cmd

import (
	"github.com/mheers/k8s-secret-ui/server"
	"github.com/spf13/cobra"
)

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return server.Run()
		},
	}
)

func init() {
	// serverCmd.PersistentFlags().StringVarP(&recipientFile, "recipient-file", "r", recipientFileDefault, "")
}
