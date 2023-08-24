package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"sync"
)

const path = "resources/config.json"

var initOnce sync.Once

// Config is the exported global constant for server configurations
var Config = &TodoConfig{}

type TodoConfig struct {
	DatabaseConfig DatabaseConfig
	ServerConfig   ServerConfig
	CacheConfig    CacheConfig
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

type CacheConfig struct {
	Limit int    `json:"limit"`
	Type  string `json:"type"`
}

func LoadConfig() error {
	pat, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Assuming that you are 2 directories deep into the project root
	// Use filepath.Dir as many times as required to get to the root
	rootPath := filepath.Dir(filepath.Dir(pat))
	err = os.Chdir(rootPath)
	if err != nil {
		log.Fatal(err)
	}
	pat, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(Config)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	initOnce.Do(func() {
		err := LoadConfig()
		if err != nil {
			log.Printf("Error loading config: %s\n", err.Error())
			panic(err)
		}
	})
}
