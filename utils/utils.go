package utils

import (
	"encoding/json"
	"io"
	"os"

	"github.com/muesli/go-app-paths"
)

var (
	DefaultServerUrl = "http://localhost:6756"
	configPath = getConfigPath()
)

func GetConfig() map[string]string {
	config := make(map[string]string)

	if configPath == "" {
		return config
	}

	file, err := os.Open(configPath)
	if err != nil {
		return config
	}
	defer file.Close()

	fileContents, err := io.ReadAll(file)
	if err != nil {
		return config
	}

	_ = json.Unmarshal(fileContents, &config)
	return config
}

func getConfigPath() string {
	configPath, err := gap.NewScope(gap.User, "timeruler").ConfigPath("config.json")
	if err != nil {
		return ""
	}

	return configPath
}
