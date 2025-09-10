package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadMCPConfig_JSON(t *testing.T) {
	// Create temporary JSON file
	jsonContent := `{
		"mcpServers": {
			"test-server": {
				"command": "python",
				"args": ["-m", "test"],
				"env": {
					"TEST_VAR": "test_value"
				}
			}
		}
	}`

	tmpDir := t.TempDir()
	jsonFile := filepath.Join(tmpDir, "test.json")
	
	err := os.WriteFile(jsonFile, []byte(jsonContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test JSON file: %v", err)
	}

	config, err := LoadMCPConfig(jsonFile)
	if err != nil {
		t.Fatalf("Failed to load JSON config: %v", err)
	}

	if len(config.McpServers) != 1 {
		t.Errorf("Expected 1 server, got %d", len(config.McpServers))
	}

	server, exists := config.McpServers["test-server"]
	if !exists {
		t.Fatal("Expected test-server to exist")
	}

	if server.Command != "python" {
		t.Errorf("Expected command 'python', got '%s'", server.Command)
	}
}

func TestLoadMCPConfig_YAML(t *testing.T) {
	// Create temporary YAML file
	yamlContent := `mcpServers:
  test-server:
    command: python
    args:
      - "-m"
      - "test"
    env:
      TEST_VAR: test_value
`

	tmpDir := t.TempDir()
	yamlFile := filepath.Join(tmpDir, "test.yaml")
	
	err := os.WriteFile(yamlFile, []byte(yamlContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test YAML file: %v", err)
	}

	config, err := LoadMCPConfig(yamlFile)
	if err != nil {
		t.Fatalf("Failed to load YAML config: %v", err)
	}

	if len(config.McpServers) != 1 {
		t.Errorf("Expected 1 server, got %d", len(config.McpServers))
	}

	server, exists := config.McpServers["test-server"]
	if !exists {
		t.Fatal("Expected test-server to exist")
	}

	if server.Command != "python" {
		t.Errorf("Expected command 'python', got '%s'", server.Command)
	}
}

func TestLoadMCPConfig_FileNotFound(t *testing.T) {
	_, err := LoadMCPConfig("nonexistent.json")
	if err == nil {
		t.Error("Expected error for nonexistent file, got nil")
	}
}

func TestLoadMCPConfig_InvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	invalidFile := filepath.Join(tmpDir, "invalid.json")
	
	err := os.WriteFile(invalidFile, []byte("{invalid json"), 0644)
	if err != nil {
		t.Fatalf("Failed to create invalid JSON file: %v", err)
	}

	_, err = LoadMCPConfig(invalidFile)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}