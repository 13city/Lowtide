#!/bin/bash

# Run Go tests with verbose output and redirect the output to a file
go test -v ./... > test_report.txt 2>&1

# Check if the tests passed or failed by examining the exit status
if [ $? -eq 0 ]; then
    echo "Tests passed successfully. Check the test_report.txt for detailed output."
else
    echo "Tests failed. Check the test_report.txt for detailed output."
fi
