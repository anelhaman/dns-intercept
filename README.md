# dns-intercept

## Overview
`dns-intercept` is a CLI tool for managing DNS interceptions in CoreDNS within a Kubernetes environment. This tool allows you to add, remove, and view DNS rewrite rules in the CoreDNS configuration.

## Installation
To install `dns-intercept`, clone the repository and build it using Go:

```sh
# Clone the repository
git clone https://github.com/anelhaman/dns-intercept.git
cd dns-intercept

# Build the binary for Linux
GOOS=linux GOARCH=amd64 go build -o dns-intercept-linux-amd64
GOOS=linux GOARCH=arm64 go build -o dns-intercept-linux-arm64

# Move to a directory in your PATH
mv dns-intercept-linux-amd64 /usr/local/bin/

# Build the binary for macOS (Intel and Apple Silicon)
GOOS=darwin GOARCH=amd64 go build -o dns-intercept-darwin-amd64
GOOS=darwin GOARCH=arm64 go build -o dns-intercept-darwin-arm64

# Build the binary for Windows
GOOS=windows GOARCH=amd64 go build -o dns-intercept-windows-amd64.exe
GOOS=windows GOARCH=arm64 go build -o dns-intercept-windows-arm64.exe
```

## Usage
The CLI supports three main operations: `add`, `remove`, and `show`.

### Add DNS Interception Rules
This command adds DNS interception rules to the CoreDNS ConfigMap.

#### Examples:
```sh
# Add a single interception rule
dns-intercept add a.domain.local domain.com

```

### Remove DNS Interception Rules
This command removes DNS interception rules for a specified domain from the CoreDNS ConfigMap.

#### Example:
```sh
# Remove interception for a domain
dns-intercept remove a.domain.local
```

### Show DNS Interception Rules
This command displays the current DNS interception rules configured in CoreDNS.

#### Example:
```sh
dns-intercept show
```

### Autocompletion
The `completion` command generates an autocompletion script for various shells. For example:

```sh
# Generate Bash completion script
dns-intercept completion bash > /usr/local/etc/bash_completion.d/dns-intercept

# Generate Zsh completion script
dns-intercept completion zsh > "${fpath[1]}/_dns-intercept"

# Generate Fish completion script
dns-intercept completion fish > ~/.config/fish/completions/dns-intercept.fish
```

## Prerequisites
For more information on how CoreDNS handles rewrites, refer to the official documentation: [CoreDNS Rewrite Plugin](https://coredns.io/plugins/rewrite/)

- Kubernetes cluster with CoreDNS deployed
- `kubectl` configured with appropriate cluster access
- Go 1.16 or later (for building from source)

## Configuration
By default, dns-intercept looks for the CoreDNS ConfigMap in the `kube-system` namespace.

## File Format for Bulk Import
When using the `-f` flag with the `add` command, the input file should contain one rule per line in the following format:
```
source.domain.local target.domain.com
another.source.local another.target.com
```

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
