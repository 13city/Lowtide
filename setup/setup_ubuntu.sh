#!/bin/bash
# Ubuntu

# Define Go version
GO_VERSION="1.18.1"

# Define project specific variables
PROJECT_DIR="lowtide"
REPO_URL="https://github.com/13city/Lowtide.git"
MODULE_PATH="github.com/13city/Lowtide" # This should match the module's path used in go mod init

# Update and Upgrade Ubuntu Packages
echo "Updating and upgrading Ubuntu packages..."
sudo apt-get update && sudo apt-get -y upgrade

# Install Go if not already installed
if ! command -v go &> /dev/null; then
    echo "Installing Go..."
    wget "https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz" -O go${GO_VERSION}.linux-amd64.tar.gz
    sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
    echo "export PATH=\$PATH:/usr/local/go/bin" >> "$HOME/.profile"
    source "$HOME/.profile"
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
