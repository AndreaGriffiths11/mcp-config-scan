package config

import (
	"encoding/json"
	"os"
)

type MCPConfig struct {
	McpServers map[string]ServerConfig `json:"mcpServers"`
	Defaults   *DefaultsConfig         `json:"defaults,omitempty"`
}

type ServerConfig struct {
	Command    string                 `json:"command,omitempty"`
	Args       []string               `json:"args,omitempty"`
	Env        map[string]string      `json:"env,omitempty"`
	Settings   map[string]interface{} `json:"settings,omitempty"`
	Timeout    int                    `json:"timeout,omitempty"`
	Disabled   bool                   `json:"disabled,omitempty"`
	WorkingDir string                 `json:"workingDir,omitempty"`
}

type DefaultsConfig struct {
	Timeout int                    `json:"timeout,omitempty"`
	Env     map[string]string      `json:"env,omitempty"`
	Args    []string               `json:"args,omitempty"`
	Settings map[string]interface{} `json:"settings,omitempty"`
}

func LoadMCPConfig(filePath string) (*MCPConfig, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config MCPConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}