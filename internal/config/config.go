package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/simplyzetax/oracle/pkg/types"
)

// GetConfigDir returns the oracle config directory path
func GetConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(homeDir, ".oracle")

	// Create config directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}

	return configDir, nil
}

// GetFirstRunFilePath returns the path to the first run marker file
func GetFirstRunFilePath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, ".first_run_complete"), nil
}

// IsFirstRun checks if this is the first time running oracle
func IsFirstRun() bool {
	firstRunFile, err := GetFirstRunFilePath()
	if err != nil {
		return true // Assume first run if we can't check
	}

	_, err = os.Stat(firstRunFile)
	return os.IsNotExist(err)
}

// MarkFirstRunComplete creates the first run marker file
func MarkFirstRunComplete() error {
	firstRunFile, err := GetFirstRunFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(firstRunFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("First run completed at: %s\n",
		os.Getenv("USER")))
	return err
}

// GetConfigFilePath returns the path to the config file
func GetConfigFilePath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "config.json"), nil
}

// LoadConfig loads the configuration from the config file
func LoadConfig() (*types.Config, error) {
	configFile, err := GetConfigFilePath()
	if err != nil {
		return nil, err
	}

	// If config file doesn't exist, return empty config
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return &types.Config{}, nil
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config types.Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

// SaveConfig saves the configuration to the config file
func SaveConfig(config *types.Config) error {
	configFile, err := GetConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configFile, data, 0600); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// GetAPIKey retrieves the API key from config, environment, or parameter
func GetAPIKey(flagAPIKey string) (string, error) {
	// 1. Check flag parameter first
	if flagAPIKey != "" {
		return flagAPIKey, nil
	}

	// 2. Check environment variable
	if envKey := os.Getenv("GOOGLE_AI_API_KEY"); envKey != "" {
		return envKey, nil
	}

	// 3. Check config file
	config, err := LoadConfig()
	if err != nil {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	if config.APIKey != "" {
		return config.APIKey, nil
	}

	return "", nil
}

// SetAPIKey saves the API key to the config file
func SetAPIKey(apiKey string) error {
	config, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load existing config: %w", err)
	}

	config.APIKey = apiKey

	if err := SaveConfig(config); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	return nil
}
