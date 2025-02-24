package cmd

import (
	"fmt"
	"strings"
)

func removeAllK8sRules() error {
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

	// Get Corefile content
	corefile := configMap.Data["Corefile"]

	// Remove all custom rules
	lines := strings.Split(corefile, "\n")
	var newLines []string

	for _, line := range lines {
		if strings.Contains(line, "rewrite name") {
			continue
		}
		newLines = append(newLines, line)
	}

	// Update CoreDNS config
	configMap.Data["Corefile"] = strings.Join(newLines, "\n")

	if err := updateCoreDNSConfigMap(clientset, configMap); err != nil {
		return fmt.Errorf("failed to update CoreDNS configmap: %v", err)
	}

	// Restart CoreDNS pods after removing rules
	if err := restartCoreDNS(clientset); err != nil {
		return fmt.Errorf("failed to restart CoreDNS: %v", err)
	}

	return nil
}
