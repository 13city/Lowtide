package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net"
)

type Config struct {
    Timeout int    `json:"timeout"`
    Ports   string `json:"ports"`
    StartIP string `json:"startIP,omitempty"`
    EndIP   string `json:"endIP,omitempty"`
    Subnet  string `json:"subnet,omitempty"`
}

func LoadConfig(filePath string) (*Config, error) {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        return nil, fmt.Errorf("failed to read %s: %v", filePath, err)
    }
    var cfg Config
    if err = json.Unmarshal(data, &cfg); err != nil {
        return nil, fmt.Errorf("failed to unmarshal config: %v", err)
    }
    return &cfg, nil
}

func ResolveIPs(cfg *Config) ([]string, error) {
    if cfg.Subnet != "" {
        return generateIPsFromSubnet(cfg.Subnet)
    }
    return generateSequentialIPs(cfg.StartIP, cfg.EndIP)
}

func generateIPsFromSubnet(subnetStr string) ([]string, error) {
    _, ipv4Net, err := net.ParseCIDR(subnetStr)
    if err != nil {
        return nil, fmt.Errorf("parsing CIDR failed: %v", err)
    }
    
    var ips []string
    for ip := ipv4Net.IP.Mask(ipv4Net.Mask); ipv4Net.Contains(ip); incrementIP(ip) {
        ips = append(ips, ip.String())
    }
    return ips, nil
}

func generateSequentialIPs(startIP, endIP string) ([]string, error) {
    start := net.ParseIP(startIP)
    end := net.ParseIP(endIP)
    if start == nil || end == nil {
        return nil, fmt.Errorf("invalid IP address")
    }
    
    var ips []string
    for ip := start; !ip.Equal(end); incrementIP(ip) {
        ips = append(ips, ip.String())
    }
    ips = append(ips, end.String())
    return ips, nil
}

func incrementIP(ip net.IP) {
    for j := len(ip) - 1; j >= 0; j-- {
        ip[j]++
        if ip[j] > 0 {
            break
        }
    }
}