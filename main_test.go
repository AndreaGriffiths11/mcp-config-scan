package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestMainFunction(t *testing.T) {
	// Test that main function doesn't panic
	// This is a basic smoke test
	if os.Getenv("BE_MAIN") == "1" {
		main()
		return
	}

	// Run the test in a subprocess
	cmd := exec.Command(os.Args[0], "-test.run=TestMainFunction")
	cmd.Env = append(os.Environ(), "BE_MAIN=1")
	err := cmd.Run()
	
	// We expect main to exit with code 1 (no configs found)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			// Exit code 1 is expected when no configs are found
			if exitError.ExitCode() == 1 {
				return // This is expected
			}
		}
		t.Errorf("Unexpected error running main: %v", err)
	}
}

func TestVersion(t *testing.T) {
	// Simple test to ensure the package compiles
	// More comprehensive CLI tests would go here
	t.Log("Main package compiles successfully")
}