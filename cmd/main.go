package main

import (
	"log"
	"os"

	"github.com/spf13/pflag"
)

func main() {
	// Remove any flags that were added by libraries automatically.
	pflag.CommandLine = pflag.NewFlagSet("mintless-cli", pflag.ExitOnError)

	out := os.Stdout
	cmd := newRootCmd(out)
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
