#!/bin/bash
# Manjaro

# Define project specific variables
PROJECT_DIR="lowtide"
REPO_URL="https://github.com/13city/lowtide.git"
MODULE_PATH="github.com/13city/lowtide" # This should match the module's path used in go mod init

# Update and Upgrade Manjaro Packages
echo "Updating and upgrading Manjaro packages..."
sudo pacman -Syu --noconfirm

# Install Go if not already installed
if ! command -v go &> /dev/null; then
    echo "Installing Go..."
    sudo pacman -S go --noconfirm
else
    echo "Go is already installed."
fi

# Navigate to the project directory
cd "$(dirname "$0")/.."

# Check if it's already a Go module and if not, initialize and tidy it
if [ ! -f "go.mod" ]; then
    echo "Initializing the Go module..."
    go mod init "$MODULE_PATH"
    echo "Tidying the Go module..."
    go mod tidy
fi

# Build the project
echo "Building the LowTide project..."
go build -o lowtide

if [ $? -eq 0 ]; then
    echo "LowTide has been successfully built."
else
    echo "Build failed. Please check the Go installation and the project repository."
    exit 1
fi

echo "Setup completed. Run ./lowtide to start the application."
