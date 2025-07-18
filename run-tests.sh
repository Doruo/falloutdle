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

# Export .env variables
set -a  # auto-export enable
source .env
set +a # auto-export disable

# Project config
PROJECT_NAME="falloutdle"
PROJECT_PATH="$HOME/dev/$PROJECT_NAME"

# Test config
TESTS_PATH="$PROJECT_PATH/tests"
VERBOSE=false

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

print_info "Starting $PROJECT_NAME tests..."

# Pre-flight checks
# Navigate to target directory if specified
if [[ -n "$TARGET_SUBDIRECTORY" ]]; then
    EXECUTION_PATH="$TARGET_SUBDIRECTORY"
else
    EXECUTION_PATH="$TESTS_PATH"
fi

print_info "Execution path: $EXECUTION_PATH"
cd "$EXECUTION_PATH"

if ($VERBOSE) 
then
    print_warning "Verbose enabled !"
    EXECUTION_COMMAND="go test -v"
else
    EXECUTION_COMMAND="go test"
fi

print_info "Execution command: $EXECUTION_COMMAND [file].go"

for file in *
do
  print_info "Testing file: $file"
  $EXECUTION_COMMAND $file
done
