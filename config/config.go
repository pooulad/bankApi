package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type PostgresConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Name     string `json:"name"`
	SSL      string `json:"ssl"`
}

func ReadPostgresConfig(path string) (*PostgresConfig, error) {
	if path == "" {
		return nil, fmt.Errorf("no path provided")
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg PostgresConfig
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
