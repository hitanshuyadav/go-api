#!/bin/bash

set -euo pipefail

echo "Checking Go dependencies..."

go mod verify

echo "Installing vulnerability scanner..."

go install golang.org/x/vuln/cmd/govulncheck@latest

echo "Scanning application..."

govulncheck ./...

echo "Security checks completed."
