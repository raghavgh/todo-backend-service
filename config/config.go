package config

import (
	"encoding/json"
	"os"
)

const path = "resources/config.json"

type TodoConfig struct {
	DatabaseConfig DatabaseConfig
	ServerConfig   ServerConfig
}

type DatabaseConfig struct {
	Port string `json:"port"`
	Url  string `json:"url"`
}

type ServerConfig struct {
	Port         string `json:"port"`
	ReadTimeout  int64  `json:"readTimeout"`
	WriteTimeout int64  `json:"writeTimeout"`
}

func LoadConfig() (*TodoConfig, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	todoConfig := &TodoConfig{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(todoConfig)
	if err != nil {
		return nil, err
	}
	return todoConfig, nil
}
