package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isError  bool
}{
	{"valid-data", 100.0, 20.0, 5.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"fraction-result", 8.0, 5, 1.6, false},
}

func TestDivide(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isError && err == nil {
			t.Error("Expected error but didn't get error")
		} else if !tt.isError && err != nil {
			t.Error("Got unexpected error")
		}
		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}

}
