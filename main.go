package main

import (
    "flag"
)

func main() {
    useConfig := flag.Bool("useConfig", false, "Load settings from config.json")
    flag.Parse()

    logger := SetupLogger()

    if !*useConfig {
        logger.Fatal("Configuration required. Use -useConfig to load from config.json.")
    }

    cfg, err := LoadConfig("configuration/config.json") // Adjust the path to your config.json if necessary
    if err != nil {
        logger.Fatal("Failed to load config: ", err)
    }

    ips, err := ResolveIPs(cfg)
    if err != nil {
        logger.Fatal(err)
    }

    ports, err := parsePorts(cfg.Ports)
    if err != nil {
        logger.Fatal(err)
    }

    ScanNetwork(ips, ports, cfg.Timeout, logger)
}