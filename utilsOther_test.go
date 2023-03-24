package main

import (
	"runtime"
	"testing"
)

func TestDetectOS(t *testing.T) {
	// Call the function being tested
	actualOS := detectOS()

	// Compare the result with the expected value
	if actualOS != runtime.GOOS {
		t.Errorf("Unexpected OS value, got %s but expected %s", actualOS, runtime.GOOS)
	}
}

func TestIsHighEntropy(t *testing.T) {
	// Test a low entropy string
	if isHighEntropy("aaaaa") {
		t.Errorf("Expected low entropy for 'aaaaa'")
	}

	// Test a high entropy string
	if !isHighEntropy("aAbBcC123!@#") {
		t.Errorf("Expected high entropy for 'aAbBcC123!@#'")
	}
}
