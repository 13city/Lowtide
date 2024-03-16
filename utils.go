package main

import (
    "fmt"
    "strconv"
    "strings"
)

// parsePorts converts a ports string into a slice of port strings.
// It supports single ports, comma-separated lists, and ranges.
func ParsePorts(portsConfig string) ([]string, error) {
    var ports []string
    portsConfig = strings.ReplaceAll(portsConfig, " ", "")
    if strings.Contains(portsConfig, ",") {
        ports = strings.Split(portsConfig, ",")
    } else if strings.Contains(portsConfig, "-") {
        rangeParts := strings.Split(portsConfig, "-")
        start, err := strconv.Atoi(rangeParts[0])
        if err != nil {
            return nil, fmt.Errorf("invalid start port: %v", err)
        }
        end, err := strconv.Atoi(rangeParts[1])
        if err != nil {
            return nil, fmt.Errorf("invalid end port: %v", err)
        }
        for p := start; p <= end; p++ {
            ports = append(ports, strconv.Itoa(p))
        }
    } else {
        ports = append(ports, portsConfig)
    }
    return ports, nil
}
