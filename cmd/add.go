package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [domain] [target] or add -f [filename]",
	Short: "Add DNS interception rules to CoreDNS ConfigMap",
	Long: `Add DNS interception rules to CoreDNS ConfigMap.
Examples:
  dns-interceptor add a.domain.local domain.com
  dns-interceptor add -f records.txt`,
	RunE: runAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("file", "f", "", "Specify the record file containing rules to add")
}

func runAdd(cmd *cobra.Command, args []string) error {
	filename, _ := cmd.Flags().GetString("file")
	var rewriteLines []string
	var err error

	if filename != "" {
		// File-based mode
		rewriteLines, err = readInterceptRecords(filename)
		if err != nil {
			return fmt.Errorf("failed to read intercept records: %w", err)
		}
	} else {
		// Inline mode
		if len(args) != 2 {
			return fmt.Errorf("inline mode requires exactly 2 arguments: domain and target")
		}
		rewriteLine := fmt.Sprintf("%s %s", args[0], args[1])
		rewriteLines = []string{rewriteLine}
	}

	if err := addK8sRule(rewriteLines); err != nil {
		return fmt.Errorf("failed to add rule to kubernetes: %w", err)
	}

	return nil
}

func readInterceptRecords(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rewriteLines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		rewriteLines = append(rewriteLines, strings.TrimSpace(line))
	}

	return rewriteLines, scanner.Err()
}
