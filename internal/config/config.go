package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config holds app configuration
type Config struct {
	StaticDir string         `json:"static_dir"`
	Server struct {
		PORT string `json:"port"`
		HOST string `json:"host"`
	} `json:"server"`
	DB struct {
		DSN string `json:"dsn"`
	} `json:"db"`
	Session struct {
		Key string `json:"key"`
	} `json:"session"`
}

// Read config.json and populate Config
func Load(path string) (*Config, error){
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("config: open %q: %w", path, err)
	}
	defer f.Close()

	var cfg Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("config: decode %q: %w", path, err)
	}
	return &cfg, nil
}