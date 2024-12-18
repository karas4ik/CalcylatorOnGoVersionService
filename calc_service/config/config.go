package config

import (
	"encoding/json"
	"os"
)

var configFilePath = "config.json"

type Config struct {
	LogLevel string `json:"log_level"`
}

func LoadConfig() Config {
	file, err := os.Open(configFilePath)
	if err != nil {
		return Config{LogLevel: "info"}
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{LogLevel: "info"}
	}
	return cfg
}
