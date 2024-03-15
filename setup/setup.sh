#!/bin/bash

# Define Go version and project specific variables
GO_VERSION="1.18.1"
PROJECT_DIR="Lowtide"
REPO_URL="https://github.com/13city/Lowtide.git"
MODULE_PATH="github.com/13city/Lowtide" # This should match the module's path used in go mod init

# Function to install Go for Debian/Ubuntu
install_go_debian() {
    echo "Installing Go for Debian/Ubuntu..."
    wget "https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz" -O go${GO_VERSION}.linux-amd64.tar.gz
    sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
    echo "export PATH=\$PATH:/usr/local/go/bin" >> "$HOME/.profile"
    source "$HOME/.profile"
}

# Function to install Go for Arch/Manjaro
install_go_arch() {
    echo "Installing Go for Arch/Manjaro..."
    sudo pacman -S go --noconfirm
}

# Detect the package manager and install dependencies
if command -v apt &> /dev/null; then
    echo "Debian/Ubuntu based system detected."
    echo "Updating and upgrading Debian/Ubuntu packages..."
    sudo apt-get update && sudo apt-get -y upgrade
    if ! command -v go &> /dev/null; then
        install_go_debian
    else
        echo "Go is already installed."
    fi
elif command -v pacman &> /dev/null; then
    echo "Arch/Manjaro based system detected."
    echo "Updating and upgrading Arch/Manjaro packages..."
    sudo pacman -Syu --noconfirm
    if ! command -v go &> /dev/null; then
        install_go_arch
    else
        echo "Go is already installed."
    fi
else
    echo "Unsupported Linux distribution."
    exit 1
fi

# Navigate to the project directory
cd "$(dirname "$0")/.."

# Initialize and tidy Go module if needed
if [ ! -f "go.mod" ]; then
    echo "Initializing the Go module..."
    go mod init "$MODULE_PATH"
    echo "Tidying the Go module..."
    go mod tidy
fi

# Build the project
echo "Building the Lowtide project..."
go build -o Lowtide

if [ $? -eq 0 ]; then
    echo "Lowtide has been successfully built."
else
    echo "Build failed. Please check the Go installation and the project repository."
    exit 1
fi

echo "Setup completed. Run ./Lowtide to start the application."