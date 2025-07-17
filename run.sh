#!/bin/bash

# =============================================================================
# Generic Project Starter Script
# =============================================================================
# Description: Configurable script to start development servers
# Usage: ./run.sh [target_subdirectory]
# Author: Marc Haye
# Version: 1.0
# =============================================================================

# -----------------------------------------------------------------------------
# Configuration Variables
# -----------------------------------------------------------------------------

# Main project configuration
PROJECT_NAME="falloutdle"
PROJECT_PATH="$HOME/dev/$PROJECT_NAME"
MAIN_PATH="cmd/server"                    # Default server path
MAIN_FILE="main.go"                       # Main executable file
FULL_PATH="$PROJECT_PATH/$MAIN_PATH"

# Runtime configuration
BUILD_COMMAND="go run"                    # Command to run the application
LOG_LEVEL="info"                          # Default log level
PORT="8080"                               # Default port

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

# Navigate to project subdirectory
# Usage: goto_project [subdirectory]
goto_project() {
    local target_path="$1"
    
    if [[ -n "$target_path" ]]; then
        local full_target="$PROJECT_PATH/$target_path"
        if [[ -d "$full_target" ]]; then
            cd "$full_target"
            print_info "Navigated to: $(pwd)"
        else
            print_error "Directory $full_target doesn't exist"
            return 1
        fi
    else
        cd "$PROJECT_PATH"
        print_info "Navigated to project root: $(pwd)"
    fi
}

# Check if required dependencies are installed
check_dependencies() {
    print_info "Checking dependencies..."
    
    # Check for Go installation
    if ! command -v go &> /dev/null; then
        print_error "Go is not installed or not in PATH"
        exit 1
    fi
    
    # Add more dependency checks here as needed
    # Example: docker, npm, python, etc.
    
    print_info "All dependencies are available"
}

# Validate project structure
validate_project_structure() {
    print_info "Validating project structure..."
    
    # Check if project directory exists
    if [[ ! -d "$PROJECT_PATH" ]]; then
        print_error "Project directory $PROJECT_PATH doesn't exist"
        exit 1
    fi
    
    # Check if main path exists
    if [[ ! -d "$FULL_PATH" ]]; then
        print_error "Main path $FULL_PATH doesn't exist"
        exit 1
    fi
    
    # Check if main file exists
    if [[ ! -f "$FULL_PATH/$MAIN_FILE" ]]; then
        print_error "Main file $FULL_PATH/$MAIN_FILE doesn't exist"
        exit 1
    fi
    
    print_info "Project structure is valid"
}

# Setup environment variables
setup_environment() {
    print_info "Setting up environment..."
    
    # Set common environment variables
    export APP_ENV="${APP_ENV:-development}"
    export LOG_LEVEL="$LOG_LEVEL"
    export PORT="$PORT"
    
    # Add more environment setup here
    # Example: database URLs, API keys, etc.
    
    print_info "Environment configured for $APP_ENV mode"
}

# Clean up function (called on script exit)
cleanup() {
    print_info "Cleaning up..."
    # Add cleanup tasks here
    # Example: stop background processes, remove temp files, etc.
}

# Display help information
show_help() {
    cat << EOF
Usage: $0 [OPTIONS] [SUBDIRECTORY]

Generic project starter script for $PROJECT_NAME

OPTIONS:
    -h, --help          Show this help message
    -p, --port PORT     Set custom port (default: $PORT)
    -v, --verbose       Enable verbose output
    --dry-run          Show what would be executed without running

EXAMPLES:
    $0                  # Start server in default location
    $0 api/v1          # Start server in api/v1 subdirectory
    $0 --port 3000     # Start server on port 3000

EOF
}

# -----------------------------------------------------------------------------
# Main Execution Logic
# -----------------------------------------------------------------------------

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        -h|--help)
            show_help
            exit 0
            ;;
        -p|--port)
            PORT="$2"
            shift 2
            ;;
        -v|--verbose)
            set -x  # Enable verbose mode
            shift
            ;;
        -dr|--dry-run)
            DRY_RUN=true
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

# Set up signal handlers for graceful shutdown
trap cleanup EXIT
trap 'print_warning "Interrupted by user"; exit 130' INT

# Pre-flight checks
print_info "Starting $PROJECT_NAME server..."
check_dependencies
validate_project_structure
setup_environment

# Navigate to target directory if specified
if [[ -n "$TARGET_SUBDIRECTORY" ]]; then
    if ! goto_project "$TARGET_SUBDIRECTORY"; then
        exit 1
    fi
    EXECUTION_PATH="$PROJECT_PATH/$TARGET_SUBDIRECTORY"
else
    EXECUTION_PATH="$FULL_PATH"
fi

# Prepare the command to execute
COMMAND="$BUILD_COMMAND $EXECUTION_PATH/$MAIN_FILE"

# Display execution information
print_info "Execution path: $EXECUTION_PATH"
print_info "Command: $COMMAND"
print_info "Port: $PORT"
print_info "Environment: $APP_ENV"

# Execute the command (or show what would be executed in dry-run mode)
if [[ "$DRY_RUN" == "true" ]]; then
    print_warning "DRY RUN - Would execute: $COMMAND"
else
    cd "$EXECUTION_PATH"
    print_info "Starting server... (Press Ctrl+C to stop)"
    exec $COMMAND
fi