package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type Receiver struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}

type Config struct {
	FilePaths      []string   `toml:"filePaths"`
	SenderEmail    string     `toml:"senderEmail"`
	SenderPassword string     `toml:"senderPassword"`
	Receivers      []Receiver `toml:"receivers"`
	Subject        string     `toml:"subject"`
	Body           string     `toml:"body"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	if _, err := toml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
