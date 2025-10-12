package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	if err := writeConfigFile(*c); err != nil {
		return err
	}
	return nil
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	configFile, err := os.Open(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("error opening config file: %w", err)
	}
	defer configFile.Close()
	var config Config
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		return Config{}, fmt.Errorf("error decoding config file: %w", err)
	}
	return config, nil
}

func getConfigFilePath() (fileName string, err error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error finding current user home dir: %w", err)
	}
	return homePath + "/" + configFileName, nil
}

func writeConfigFile(config Config) error {
	configFileName, err := getConfigFilePath()
	if err != nil {
		return err
	}
	configFile, err := os.Create(configFileName)
	if err != nil {
		return fmt.Errorf("error creating config file: %w", err)
	}
	defer configFile.Close()
	if err = json.NewEncoder(configFile).Encode(&config); err != nil {
		return fmt.Errorf("error encoding config file: %w", err)
	}
	return nil
}
