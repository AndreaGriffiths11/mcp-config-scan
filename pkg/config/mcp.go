package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type MCPConfig struct {
	McpServers map[string]ServerConfig `json:"mcpServers" yaml:"mcpServers"`
	Defaults   *DefaultsConfig         `json:"defaults,omitempty" yaml:"defaults,omitempty"`
}

type ServerConfig struct {
	Command    string                 `json:"command,omitempty" yaml:"command,omitempty"`
	Args       []string               `json:"args,omitempty" yaml:"args,omitempty"`
	Env        map[string]string      `json:"env,omitempty" yaml:"env,omitempty"`
	Settings   map[string]interface{} `json:"settings,omitempty" yaml:"settings,omitempty"`
	Timeout    int                    `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	Disabled   bool                   `json:"disabled,omitempty" yaml:"disabled,omitempty"`
	WorkingDir string                 `json:"workingDir,omitempty" yaml:"workingDir,omitempty"`
}

type DefaultsConfig struct {
	Timeout int                    `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	Env     map[string]string      `json:"env,omitempty" yaml:"env,omitempty"`
	Args    []string               `json:"args,omitempty" yaml:"args,omitempty"`
	Settings map[string]interface{} `json:"settings,omitempty" yaml:"settings,omitempty"`
}

func LoadMCPConfig(filePath string) (*MCPConfig, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config MCPConfig
	
	// Determine file format based on extension
	ext := strings.ToLower(filepath.Ext(filePath))
	
	switch ext {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(data, &config)
	case ".json":
		err = json.Unmarshal(data, &config)
	default:
		// Default to JSON for backward compatibility
		err = json.Unmarshal(data, &config)
	}
	
	if err != nil {
		return nil, err
	}

	return &config, nil
}