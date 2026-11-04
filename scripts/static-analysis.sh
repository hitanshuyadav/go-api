#!/bin/bash

set -euo pipefail

echo "Running Go static analysis..."

go vet ./...

echo "Static analysis completed successfully."
