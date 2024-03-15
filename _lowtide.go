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

	"gopkg.in/natefinch/lumberjack.v2"
)

// Config stores settings loaded from a JSON configuration file.
type Config struct {
	Timeout int    `json:"timeout"`
	Ports   string `json:"ports"`
	StartIP string `json:"startIP,omitempty"`
	EndIP   string `json:"endIP,omitempty"`
	Subnet  string `json:"subnet,omitempty"`
}

var (
	config Config           // Holds the application configuration.
	logger *log.Logger      // Application logger for outputting information.
	sem    chan struct{}    // Semaphore to limit concurrency.
)

func main() {
	var wg sync.WaitGroup
	useConfig := flag.Bool("useConfig", false, "Load settings from config.json")
	flag.Parse()

	logger = setupLogger() // Initialize the logger.
	logger.Println("Starting LowTide...")

	if *useConfig {
		loadConfig() // Load the configuration from the JSON file.

		// If a subnet is specified, generate IP addresses from it.
		if config.Subnet != "" {
			ips, err := generateIPsFromSubnet(config.Subnet)
			if err != nil {
				logger.Fatalf("Error generating IPs from subnet: %v", err)
			}
			logger.Printf("Scanning the subnet %s, generated %d IPs.", config.Subnet, len(ips))
			config.StartIP = ips[0]
			config.EndIP = ips[len(ips)-1]
		}
	} else {
		logger.Fatal("Configuration required. Use -useConfig to load from config.json.")
	}

	// Limit the number of concurrent goroutines.
	sem = make(chan struct{}, 20)
	ips := generateSequentialIPs(config.StartIP, config.EndIP)
	ports := parsePorts(config.Ports)

	// Scan each IP and port combination.
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
	logger.Println("Scanning completed.")
}

// loadConfig reads and parses the configuration from a JSON file.
func loadConfig() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		logger.Fatalf("Failed to read config.json: %v", err)
	}
	if err := json.Unmarshal(file, &config); err != nil {
		logger.Fatalf("Failed to parse config.json: %v", err)
	}
	logger.Println("Config loaded successfully.")
}

func generateIPsFromSubnet(subnetStr string) ([]string, error) {
	var ips []string
	_, subnet, err := net.ParseCIDR(subnetStr)
	if err != nil {
		return nil, fmt.Errorf("invalid subnet format: %v", err)
	}
	for ip := subnet.IP.Mask(subnet.Mask); subnet.Contains(ip); incrementIP(ip) {
		ips = append(ips, ip.String())
	}
	if len(ips) > 2 { // Optionally remove network and broadcast addresses.
		ips = ips[1 : len(ips)-1]
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
	ips = append(ips, end.String()) // Include end IP in the range.
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
	return log.New(&lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    10, // Max log file size in megabytes.
		MaxBackups: 3,  // Max number of old log files to keep.
		MaxAge:     28, // Max age in days to retain log files.
		Compress:   true, // Compress/zip old log files.
	}, "", log.LstdFlags)
}

func scanPortAndGrabBanner(targetIP, port string, timeout time.Duration) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(targetIP, port), timeout)
	if err != nil {
		logger.Printf("Port %s on %s is closed or filtered: %v", port, targetIP, err)
		return
	}
	defer conn.Close()
	banner, err := grabBanner(conn)
	if err != nil {
		logger.Printf("Failed to grab banner from port %s on %s: %v", port, targetIP, err)
		return
	}
	logger.Printf("Vulnerable server detected: IP %s, Port %s, Banner: %s", targetIP, port, banner)
}

func grabBanner(conn net.Conn) (string, error) {
	fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: %s\r\n\r\n", conn.RemoteAddr().String())
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	response := make([]byte, 4096)
	n, err := conn.Read(response)
	if err != nil {
		return "", fmt.Errorf("failed to read banner: %v", err)
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
