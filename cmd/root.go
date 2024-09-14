package main

import (
	"io"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
func newRootCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "mintless-cli",
		Version:       "0.0.1",
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}
	cmd.AddCommand(newDumpCmd(out))
	return cmd
}
