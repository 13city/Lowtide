package main

import (
    "reflect"
    "testing"
)

func TestParsePorts(t *testing.T) {
    cases := []struct {
        name     string
        input    string
        expected []string
        wantErr  bool
    }{
        {"Single Port", "80", []string{"80"}, false},
        {"Range", "80-82", []string{"80", "81", "82"}, false},
        {"Invalid Range", "abc-def", nil, true},
    }

    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            got, err := ParsePorts(c.input)
            if (err != nil) != c.wantErr {
                t.Errorf("ParsePorts() error = %v, wantErr %v", err, c.wantErr)
                return
            }
            if !reflect.DeepEqual(got, c.expected) {
                t.Errorf("ParsePorts() = %v, expected %v", got, c.expected)
            }
        })
    }
}
