#!/bin/sh
#
# This script is used to compile your program on CodeCrafters
#
# This runs before .codecrafters/run.sh
#
# Learn more: https://codecrafters.io/program-interface

set -e # Exit on failure

go build -o /tmp/codecrafters-build-grep-go cmd/mygrep/main.go cmd/mygrep/utils.go cmd/mygrep/matchers.go
