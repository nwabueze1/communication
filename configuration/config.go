package configuration

import (
	"encoding/json"
	"os"
)

type Config struct {
	DataSourceName string `json:"dataSourceName"`
}

func New() (*Config, error) {
	file, err := os.Open("configuration/config.json")
	config := Config{}
	if err != nil {
		return &config, err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&config)

	if err != nil {
		return &config, err
	}

	return &config, nil
}
