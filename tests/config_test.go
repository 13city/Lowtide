package main

import (
    "testing"
    "reflect"
)

func TestLoadConfig(t *testing.T) {
    // Mock configuration file content
    mockConfigContent := `{
        "timeout": 1000,
        "ports": "80,443",
        "subnet": "192.168.1.0/24"
    }`
    expectedConfig := &Config{
        Timeout: 1000,
        Ports:   "80,443",
        Subnet:  "192.168.1.0/24",
    }

    // Use a temporary file or mock the ioutil.ReadFile function
    // For demonstration, assuming LoadConfig directly returns the expectedConfig
    got, err := LoadConfig("path/to/mock/config.json")
    if err != nil {
        t.Fatalf("LoadConfig() error = %v", err)
    }
    if !reflect.DeepEqual(got, expectedConfig) {
        t.Errorf("LoadConfig() got = %v, expected %v", got, expectedConfig)
    }
}
