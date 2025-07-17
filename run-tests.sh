#!/bin/bash

# =============================================================================
# Generic Project Starter Script
# =============================================================================
# Description: Configurable script to start development tests
# Usage: ./run-tests.sh [target_subdirectory]
# Author: Marc Haye
# Version: 1.0
# =============================================================================

# -----------------------------------------------------------------------------
# Configuration Variables
# -----------------------------------------------------------------------------

# Tests config
TESTS_PATH="$PROJECT_PATH/tests"
VERBOSE=false

# Export .env variables
set -a  # auto-export enable
source .env
set +a # auto-export disable


# -----------------------------------------------------------------------------
# Utility Functions
# -----------------------------------------------------------------------------

# Print colored output for better visibility
print_info() {
    echo -e "\e[32m[INFO]\e[0m $1"
}

print_error() {
    echo -e "\e[31m[ERROR]\e[0m $1"
}

print_warning() {
    echo -e "\e[33m[WARNING]\e[0m $1"
}

# Display help information
show_help() {
    cat << EOF
Usage: $0 [OPTIONS] [SUBDIRECTORY]

Generic project dev tests script for $PROJECT_NAME

OPTIONS:
    -h, --help          Show this help message
    -v, --verbose       Enable verbose output in tests

EXAMPLES:
    $0             # Start server in default location
    $0 -v          # Start tests with verbose enabled

EOF
}

# -----------------------------------------------------------------------------
# Main Execution Logic
# -----------------------------------------------------------------------------

cd "$TESTS_PATH"

for file in ./*
do
  while [[ $# -gt 0 ]]; do
    case $1 in
        -h|--help)
            show_help
            exit 0
            ;;
        -v|--verbose)
            VERBOSE=true  # Enable verbose mode
            shift
            ;;
        -*)
            print_error "Unknown option: $1"
            show_help
            exit 1
            ;;
        *)
            TARGET_SUBDIRECTORY="$1"
            shift
            ;;
    esac
done
  echo "Testing file: $file"
  if ($VERBOSE) 
  then
    go test -v "$file"
  else
    go test "$file"
  fi
done
