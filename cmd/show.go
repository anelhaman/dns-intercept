package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show DNS interception rules",
	Long: `Show DNS interception rules from CoreDNS configmap.
Example: dns-interceptor show`,
	RunE: runShow,
}

func init() {
	rootCmd.AddCommand(showCmd)
}

func runShow(cmd *cobra.Command, args []string) error {
	return showK8sRules()
}

func showK8sRules() error {
	// Get kubernetes client
	clientset, err := getK8sClient()
	if err != nil {
		return fmt.Errorf("failed to create kubernetes client: %v", err)
	}

	// Get current CoreDNS config
	configMap, err := getCoreDNSConfigMap(clientset)
	if err != nil {
		return fmt.Errorf("failed to get CoreDNS configmap: %v", err)
	}

	// in case configmap is empty return print nil
	if !(strings.Contains(configMap.Data["Corefile"], "rewrite name")) {
		fmt.Println("<nil>")
		return nil
	}

	lines := strings.Split(configMap.Data["Corefile"], "\n")

	for _, line := range lines {
		if strings.Contains(line, "rewrite name") {
			// print line without rewrite name
			if strings.Contains(line, "rewrite name") {
				line = strings.Replace(line, "rewrite name", "", 1)
				fmt.Println(strings.TrimSpace(line))
			}
		}
	}
	return nil
}
