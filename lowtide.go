package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
	
	"gopkg.in/natefinch/lumberjack.v2" // Log rotation package
)

type Config struct {
	Timeout int    `json:"timeout"`
	Ports   string `json:"ports"`
	StartIP string `json:"startIP,omitempty"`
	EndIP   string `json:"endIP,omitempty"`
	Subnet  string `json:"subnet,omitempty"`
}

var (
	config Config
	logger *log.Logger
	sem    chan struct{} // Semaphore for goroutine throttling
)

func main() {
	var wg sync.WaitGroup
	useConfig := flag.Bool("useConfig", false, "Whether to load settings from config.json")
	flag.Parse()

	logger = setupLogger()
	if *useConfig {
		loadConfig()
		if config.Subnet != "" {
			ips, err := generateIPsFromSubnet(config.Subnet)
			if err != nil {
				logger.Fatalf("Error generating IPs from subnet: %v", err)
			}
			config.StartIP = ips[0]
			config.EndIP = ips[len(ips)-1]
		}
	} else {
		fmt.Println("Configuration must be provided through config.json when using -useConfig")
		os.Exit(1)
	}

	sem = make(chan struct{}, 20) // Limit concurrency to 20 goroutines
	ips := generateSequentialIPs(config.StartIP, config.EndIP)
	ports := parsePorts(config.Ports)

	for _, ip := range ips {
		for _, port := range ports {
			sem <- struct{}{}
			wg.Add(1)
			go func(ip, port string) {
				defer wg.Done()
				scanPortAndGrabBanner(ip, port, time.Millisecond*time.Duration(config.Timeout))
				<-sem
			}(ip, port)
		}
	}
	wg.Wait()
}

func loadConfig() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		logger.Fatalf("Error reading config file: %v", err)
	}
	if err := json.Unmarshal(file, &config); err != nil {
		logger.Fatalf("Error parsing config file: %v", err)
	}
}

func generateIPsFromSubnet(subnetStr string) ([]string, error) {
	var ips []string
	_, subnet, err := net.ParseCIDR(subnetStr)
	if err != nil {
		return nil, err
	}
	for ip := subnet.IP.Mask(subnet.Mask); subnet.Contains(ip); incrementIP(ip) {
		ips = append(ips, ip.String())
	}
	if len(ips) > 2 {
		ips = ips[1 : len(ips)-1] // Optionally remove network and broadcast addresses
	}
	return ips, nil
}

func generateSequentialIPs(startIP, endIP string) []string {
	var ips []string
	start := net.ParseIP(startIP)
	end := net.ParseIP(endIP)

	for ip := start; !ip.Equal(end); incrementIP(ip) {
		ips = append(ips, ip.String())
	}
	ips = append(ips, end.String()) // Include end IP

	return ips
}

func incrementIP(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] != 0 {
			break
		}
	}
}

func setupLogger() *log.Logger {
	logDir := "./logging"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create logging directory: %v", err)
	}
	filePath := filepath.Join(logDir, "lowtide.log")
	logger := log.New(&lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	}, "", log.LstdFlags)
	return logger
}

func scanPortAndGrabBanner(targetIP, port string, timeout time.Duration) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(targetIP, port), timeout)
	if err != nil {
		// Log only if detailed info is required for closed/filtered ports
		return
	}
	defer conn.Close()
	banner, err := grabBanner(conn)
	if err != nil {
		// Handle error or log if necessary
		return
	}

	// Log or process the banner information if it meets certain conditions
	logger.Printf("Vulnerable server detected: IP %s, Port %s, Banner: %s\n", targetIP, port, banner)
}

func grabBanner(conn net.Conn) (string, error) {
	fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: %s\r\n\r\n", conn.RemoteAddr().String())
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	response := make([]byte, 4096)
	n, err := conn.Read(response)
	if err != nil {
		return "", err
	}
	return string(response[:n]), nil
}

func parsePorts(portsConfig string) []string {
	var ports []string
	if strings.Contains(portsConfig, ",") {
		ports = strings.Split(portsConfig, ",")
	} else if strings.Contains(portsConfig, "-") {
		parts := strings.Split(portsConfig, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		for p := start; p <= end; p++ {
			ports = append(ports, strconv.Itoa(p))
		}
	} else {
		ports = append(ports, portsConfig)
	}
	return ports
}
