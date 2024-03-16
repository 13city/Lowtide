package main

import (
    "flag"
    // Remove the log import if it's not used directly in this file
)

func main() {
    useConfig := flag.Bool("useConfig", false, "Load settings from config.json")
    flag.Parse()

    // SetupLogger is called here, but the log package is used within logging.go,
    // not directly in main.go
    logger := SetupLogger()

    if !*useConfig {
        logger.Fatal("Configuration required. Use -useConfig to load from config.json.")
    }

    cfg, err := LoadConfig("configuration/config.json")
    if err != nil {
        logger.Fatal("Failed to load config: ", err)
    }

    ips, err := ResolveIPs(cfg)
    if err != nil {
        logger.Fatal(err)
    }

    ports, err := ParsePorts(cfg.Ports) // Ensure ParsePorts function is accessible
    if err != nil {
        logger.Fatal(err)
    }

    ScanNetwork(ips, ports, cfg.Timeout, logger)
}
