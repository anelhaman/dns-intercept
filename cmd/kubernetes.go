package cmd

import (
	"fmt"
	"strings"
)

func addK8sRule(rules []string) error {
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

	currentConfig := configMap.Data["Corefile"]
	lines := strings.Split(currentConfig, "\n")

	// Find the correct position to insert the rules
	insertPos := -1
	for i, line := range lines {
		if strings.Contains(line, "{") {
			insertPos = i + 1
			break
		}
	}

	if insertPos == -1 {
		return fmt.Errorf("could not find appropriate position to insert rules")
	}

	// Insert new rules
	newLines := make([]string, 0, len(lines)+len(rules))
	newLines = append(newLines, lines[:insertPos]...)

	skipCount := 0
	for _, rule := range rules {
		if !strings.Contains(currentConfig, rule) {
			newLines = append(newLines, "    rewrite name "+rule)
			fmt.Printf("Added record: %s\n", rule)
		} else {
			fmt.Printf("Rule already exists, skipping: %s\n", rule)
			skipCount++
		}
	}

	if skipCount == len(rules) {
		fmt.Println("All rules already exist in the configuration")
		return nil
	}

	newLines = append(newLines, lines[insertPos:]...)

	// Update the ConfigMap
	configMap.Data["Corefile"] = strings.Join(newLines, "\n")
	if err := updateCoreDNSConfigMap(clientset, configMap); err != nil {
		return fmt.Errorf("failed to update CoreDNS configmap: %v", err)
	}

	// Restart CoreDNS pods after adding rules
	if err := restartCoreDNS(clientset); err != nil {
		return fmt.Errorf("failed to restart CoreDNS: %v", err)
	}

	return nil
}

func removeK8sRule(domain string) error {
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

	currentConfig := configMap.Data["Corefile"]
	lines := strings.Split(currentConfig, "\n")

	// Find and remove the rule containing the domain
	newLines := make([]string, 0, len(lines))
	found := false

	for _, line := range lines {
		if strings.Contains(line, domain) {
			found = true
			fmt.Printf("Removed domain %s\n", domain)
			continue
		}
		newLines = append(newLines, line)
	}

	if !found {
		return fmt.Errorf("no rule found containing domain: %s", domain)
	}

	// Update the ConfigMap
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
