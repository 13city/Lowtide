#!/bin/bash

# Function to check if Lowtide binary exists
function check_binary() {
    if [[ ! -f "./Lowtide" ]]; then
        echo "Lowtide binary not found, attempting to build."
        go build -o Lowtide
        if [[ $? -ne 0 ]]; then
            echo "Failed to build Lowtide. Ensure Go is correctly set up and all dependencies are satisfied."
            exit 1
        fi
    fi
}

# Function to clean up after tests
function cleanup() {
    echo "Cleaning up..."
    rm -f config.json
}

# Function to run a test case
function run_test() {
    description=$1
    command=$2
    echo "--------------------------------"
    echo "Test: $description"
    echo "Executing: $command"
    eval "$command"
    if [[ $? -ne 0 ]]; then
        echo "Test failed: $description"
        echo "Please check the configuration or Lowtide setup."
        cleanup
        exit 1
    else
        echo "Test passed."
    fi
    echo "--------------------------------"
    echo ""
}

# Ensure Lowtide binary is ready
check_binary

# Test 1: Command-line arguments for scanning an internal network range
run_test "CLI arguments for internal network range" "./Lowtide --startIP '192.168.1.1' --endIP '192.168.1.10' --ports '80,443' --timeout 1000"

# Preparing config.json for the next test
cat << EOF > config.json
{
    "timeout": 2000,
    "ports": "22,80",
    "startIP": "10.0.0.1",
    "endIP": "10.0.0.5"
}
EOF

# Test 2: Using config.json for scanning a different internal network range
run_test "Using config.json for internal network range" "./Lowtide -useConfig"

# Test 3: Command-line arguments for scanning a small public IP range
# Warning: Only perform this test with permission from the IP range owners.
# run_test "CLI arguments for small public IP range" "./Lowtide --startIP '203.0.113.1' --endIP '203.0.113.5' --ports '80' --timeout 500"

# Cleanup after tests
cleanup

echo "All tests completed successfully."
