package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	FilePaths      []string   `json:"filePaths"`
	SenderEmail    string     `json:"senderEmail"`
	SenderPassword string     `json:"senderPassword"`
	Receivers      []Receiver `json:"receivers"`
	Subject        string     `json:"subject"`
}

type Receiver struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
