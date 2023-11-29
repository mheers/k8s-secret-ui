package cmd

import (
	"github.com/mheers/k8s-secret-ui/server"
	"github.com/spf13/cobra"
)

var (
	namespaceRegexes []string
	configMapRegexes []string
	secretRegexes    []string

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := server.NewServer(namespaceRegexes, configMapRegexes, secretRegexes)
			if err != nil {
				return err
			}

			return s.Run()
		},
	}
)

func init() {
	serverCmd.PersistentFlags().StringSliceVarP(&namespaceRegexes, "namespaces", "n", []string{}, "Allowed namespaces")
	serverCmd.PersistentFlags().StringSliceVarP(&configMapRegexes, "configmaps", "c", []string{}, "Allowed configmaps")
	serverCmd.PersistentFlags().StringSliceVarP(&secretRegexes, "secrets", "s", []string{}, "Allowed secrets")
}
