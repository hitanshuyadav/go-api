#!/bin/bash

set -euo pipefail

echo "Checking Go formatting..."

FILES=$(gofmt -l .)

if [ -n "$FILES" ]; then
    echo "Formatting issues found:"
    echo "$FILES"
    exit 1
fi

echo "All files are properly formatted."
