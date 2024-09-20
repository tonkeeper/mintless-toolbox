package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newHashCmd(out io.Writer) *cobra.Command {
	options := hashOptions{}
	cmd := &cobra.Command{
		Use:  "hash <airdrop-file>",
		Args: cobra.ExactArgs(1),
		Long: "hash extracts data from airdrop file and prints its merkle root hash",
		RunE: func(cmd *cobra.Command, args []string) error {
			options.airdropFile = args[0]
			return hash(cmd.Context(), options, out)
		},
	}
	return cmd
}

type hashOptions struct {
	airdropFile string
}

func hash(ctx context.Context, options hashOptions, out io.Writer) error {
	content, err := os.ReadFile(options.airdropFile)
	if err != nil {
		return fmt.Errorf("failed to read airdrop file: %w", err)
	}
	airdropCells, err := deserialize(content)
	if err != nil {
		return fmt.Errorf("failed to deserialize airdrop file: %w", err)
	}
	if len(airdropCells) != 1 {
		return fmt.Errorf("invalid airdrop file: expected 1 cell, got %d", len(airdropCells))
	}
	hash, err := airdropCells[0].Hash()
	if err != nil {
		return fmt.Errorf("failed to get merkle root: %w", err)
	}
	_, err = fmt.Fprintf(out, "%x\n", hash)
	return nil
}
