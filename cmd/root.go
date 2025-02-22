package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dns-intercept",
	Short: "DNS interception management tool",
	Long: `A CLI tool for managing DNS interceptions in CoreDNS.
This tool can add or remove DNS rewrites in the CoreDNS configuration.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
