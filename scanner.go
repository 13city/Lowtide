package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net"
    "sync"
    "time"
)

// ServiceRequests maps port numbers to their associated request strings.
var ServiceRequests map[string][]string

// LoadServiceRequests loads service request mappings from a JSON file.
func LoadServiceRequests(filePath string) error {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        return fmt.Errorf("failed to read %s: %v", filePath, err)
    }
    err = json.Unmarshal(data, &ServiceRequests)
    if err != nil {
        return fmt.Errorf("failed to unmarshal service requests: %v", err)
    }
    return nil
}

func ScanNetwork(ips, ports []string, timeout int, logger *log.Logger) {
    // Ensure ServiceRequests is loaded before scanning
    if err := LoadServiceRequests("service_requests.json"); err != nil {
        logger.Fatal(err)
    }

    var wg sync.WaitGroup
    sem := make(chan struct{}, 20)

    for _, ip := range ips {
        for _, port := range ports {
            wg.Add(1)
            sem <- struct{}{}
            go func(ip, port string) {
                defer wg.Done()
                scanPortAndLogResult(ip, port, timeout, logger)
                <-sem
            }(ip, port)
        }
    }
    wg.Wait()
}

func scanPortAndLogResult(ip, port string, timeout int, logger *log.Logger) {
    conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, port), time.Millisecond*time.Duration(timeout))
    if err != nil {
        logger.Printf("Port %s on %s is closed or filtered: %v", port, ip, err)
        return
    }
    defer conn.Close()

    banner := grabBanner(conn, port)
    logger.Printf("Open port detected: IP %s, Port %s, Banner: %q", ip, port, banner)
}

func grabBanner(conn net.Conn, port string) string {
    requests, exists := ServiceRequests[port]
    if !exists {
        requests = []string{"\n"} // Fallback to a newline if no specific requests are defined for the port
    }

    conn.SetReadDeadline(time.Now().Add(5 * time.Second))
    var response string
    for _, req := range requests {
        _, err := conn.Write([]byte(req))
        if err != nil {
            continue
        }

        buffer := make([]byte, 4096)
        n, err := conn.Read(buffer)
        if err == nil && n > 0 {
            response = string(buffer[:n])
            break
        }
    }

    return response
}
