package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	configJSON := `{"log_level": "debug"}`
	tempFile, err := os.CreateTemp("", "config.json")
	if err != nil {
		t.Fatalf("Unable to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write([]byte(configJSON)); err != nil {
		t.Fatalf("Unable to write to temp file: %v", err)
	}
	tempFile.Close()

	originalFile := configFilePath
	configFilePath = tempFile.Name()
	defer func() { configFilePath = originalFile }()

	cfg := LoadConfig()
	if cfg.LogLevel != "debug" {
		t.Errorf("Expected log_level to be 'debug', got '%s'", cfg.LogLevel)
	}
}
