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

var removeAll bool

func init() {
	removeCmd.Flags().BoolVarP(&removeAll, "all", "a", false, "Remove all CoreDNS rules")
	rootCmd.AddCommand(removeCmd)
}

func runRemove(cmd *cobra.Command, args []string) error {
	if removeAll {
		if len(args) != 0 {
			return fmt.Errorf("no arguments should be provided when using --all flag")
		}
		if err := removeAllK8sRules(); err != nil {
			return fmt.Errorf("failed to remove all rules from kubernetes: %w", err)
		}
	} else {
		if len(args) != 1 {
			return fmt.Errorf("exactly one domain argument is required")
		}
		if err := removeK8sRule(args[0]); err != nil {
			return fmt.Errorf("failed to remove rule from kubernetes: %w", err)
		}
	}

	return nil
}
