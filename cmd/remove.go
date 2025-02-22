package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [domain]",
	Short: "Remove DNS interception rules from CoreDNS ConfigMap",
	Long: `Remove DNS interception rules for specified domain from CoreDNS ConfigMap.
Example: dns-interceptor remove a.domain.local`,
	RunE: runRemove,
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func runRemove(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("exactly one domain argument is required")
	}

	if err := removeK8sRule(args[0]); err != nil {
		return fmt.Errorf("failed to remove rule from kubernetes: %w", err)
	}
	
	return nil
}