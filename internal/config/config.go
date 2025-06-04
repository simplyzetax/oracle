package config

import (
	"fmt"
	"os"
	"path/filepath"
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
