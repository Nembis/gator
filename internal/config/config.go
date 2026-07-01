package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (*Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("error: failed to open config file: %v", err)
	}

	config := &Config{}

	if err = json.Unmarshal(file, config); err != nil {
		return nil, fmt.Errorf("error: failled to unmarshal to json: %v", err)
	}

	return config, nil
}

func (cfg *Config) SetUser(username string) error {
	previousUsername := cfg.CurrentUserName
	cfg.CurrentUserName = username
	if err := write(cfg); err != nil {
		cfg.CurrentUserName = previousUsername
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error: can't get home directory: %v", err)
	}
	configFilePath := path.Join(homeDir, configFileName)

	return configFilePath, nil
}

func write(cfg *Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error: failed to write config to config file: %v", err)
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("eror: faield to marsual config: %v", err)
	}
	if err = os.WriteFile(configFilePath, data, 0644); err != nil {
		return fmt.Errorf("error: failed to write file to system: %v", err)
	}

	return nil
}
