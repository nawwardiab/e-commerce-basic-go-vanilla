package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config holds app configuration
type Config struct {
	StaticDir string         `yaml:"static_dir"`
	Server struct {
		PORT string `yaml:"port"`
		HOST string `yaml:"host"`
	} `yaml:"server"`
	DB struct {
		HOST string `yaml:"host"`
		PORT string `yaml:"port"`
		USER string `yaml:"user"`
		PWD string `yaml:"pwd"`
		DBNAME string `yaml:"dbname"`
	} `yaml:"db"`
	Session struct {
		Key string `yaml:"key"`
	} `yaml:"session"`
}

// Read config.json and populate Config
func Load(path string) (*Config, error){
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("config: open %q: %w", path, err)
	}
	defer f.Close()

	var cfg Config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("config: decode %q: %w", path, err)
	}
	return &cfg, nil
}