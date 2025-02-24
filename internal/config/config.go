package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DatabaseFunctions  map[string]string `json:"database_functions"`
	DatabaseProcedures map[string]string `json:"database_procedures"`
	Schemas            map[string]string `json:"schemas"`
}

var AppConfig Config

func LoadConfig(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		return err
	}

	log.Println("Configuration loaded successfully.")
	return nil
}
