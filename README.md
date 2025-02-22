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

# Add multiple rules from a file
dns-intercept add -f records.txt
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
The `completion` command generates an autocompletion script for various shells.

#### Usage:
```sh
dns-intercept completion [command]
```

#### Available Commands:
- `bash`        Generate the autocompletion script for bash
- `fish`        Generate the autocompletion script for fish
- `powershell`  Generate the autocompletion script for powershell
- `zsh`         Generate the autocompletion script for zsh

#### Example:
```sh
# Generate autocompletion script for bash
dns-intercept completion bash
```

## Requirements
- Kubernetes cluster with CoreDNS
- `kubectl` installed and configured
- Go (for building from source)

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

## License
This project is licensed under the MIT License. See `LICENSE` for details.

