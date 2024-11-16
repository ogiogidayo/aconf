package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Profiles map[string]string `yaml:"profiles"`
}

const configFileName = ".aconf.yaml"

func LoadConfig() (*Config, error) {
	path := filepath.Join(os.Getenv("HOME"), configFileName)
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{Profiles: make(map[string]string)}, nil
		}
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func SaveConfig(cfg *Config) error {
	path := filepath.Join(os.Getenv("HOME"), configFileName)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	return encoder.Encode(cfg)
}

func AddProfile(cfg *Config, profile, arn string) error {
	cfg.Profiles[profile] = arn
	return SaveConfig(cfg)
}
