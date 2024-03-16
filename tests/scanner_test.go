package main

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "reflect"
    "testing"
)

// mockServiceRequests creates a temporary JSON file to mock service requests data.
func mockServiceRequests(data map[string][]string) (string, error) {
    file, err := ioutil.TempFile("", "service_requests_*.json")
    if err != nil {
        return "", err
    }
    defer file.Close()

    jsonData, err := json.Marshal(data)
    if err != nil {
        return "", err
    }

    if _, err := file.Write(jsonData); err != nil {
        return "", err
    }

    return file.Name(), nil
}

// TestLoadServiceRequests tests the LoadServiceRequests function.
func TestLoadServiceRequests(t *testing.T) {
    expected := map[string][]string{
        "80": {"GET / HTTP/1.1\r\nHost: localhost\r\n\r\n"},
    }
    filePath, err := mockServiceRequests(expected)
    if err != nil {
        t.Fatalf("Failed to create mock service requests file: %v", err)
    }
    defer os.Remove(filePath)

    err = LoadServiceRequests(filePath)
    if err != nil {
        t.Fatalf("LoadServiceRequests() error = %v", err)
    }

    if !reflect.DeepEqual(ServiceRequests, expected) {
        t.Errorf("LoadServiceRequests() got = %v, want %v", ServiceRequests, expected)
    }
}
