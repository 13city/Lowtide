package main

import (
    "fmt"
    "log"
    "net"
    "strconv"
    "strings"
    "sync"
    "time"
)

func ScanNetwork(ips, ports []string, timeout int, logger *log.Logger) {
    var wg sync.WaitGroup
    sem := make(chan struct{}, 20) // Limit for concurrent goroutines.

    for _, ip := range ips {
        for _, port := range ports {
            wg.Add(1)
            sem <- struct{}{} // Acquire concurrency slot.
            go func(ip, port string) {
                defer wg.Done()
                scanPortAndLogResult(ip, port, timeout, logger)
                <-sem // Release slot once done.
            }(ip, port)
        }
    }
    wg.Wait()
}

func parsePorts(portsConfig string) ([]string, error) {
    portsConfig = strings.ReplaceAll(portsConfig, " ", "")
    if strings.Contains(portsConfig, ",") {
        return strings.Split(portsConfig, ","), nil
    } else if strings.Contains(portsConfig, "-") {
        parts := strings.Split(portsConfig, "-")
        startPort, err := strconv.Atoi(parts[0])
        if err != nil {
            return nil, fmt.Errorf("invalid port start: %v", err)
        }
        endPort, err := strconv.Atoi(parts[1])
        if err != nil {
            return nil, fmt.Errorf("invalid port end: %v", err)
        }
        var ports []string
        for port := startPort; port <= endPort; port++ {
            ports = append(ports, strconv.Itoa(port))
        }
        return ports, nil
    }
    return []string{portsConfig}, nil
}

func scanPortAndLogResult(ip, port string, timeout int, logger *log.Logger) {
    conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, port), time.Millisecond*time.Duration(timeout))
    if err != nil {
        logger.Printf("Port %s on %s is closed or filtered: %v", port, ip, err)
        return
    }
    defer conn.Close()
    // Implement additional logic if needed.
    logger.Printf("Open port detected: IP %s, Port %s", ip, port)
}