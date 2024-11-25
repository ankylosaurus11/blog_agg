package gatorconfig

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, configFileName), nil
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		log.Fatalf("Error fetching file path: %v", err)
	}
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	var gatorconfig Config

	err = json.Unmarshal(fileContent, &gatorconfig)
	if err != nil {
		fmt.Println(err)
	}

	return gatorconfig, err
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
