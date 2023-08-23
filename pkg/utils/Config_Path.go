package utils

import (
	"log"
	"os"
	"path/filepath"
)

func Current_Dir_Config_Path() (config_path string) {
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	config_path = filepath.Join(workDir, "config.json")
	_, err = os.Stat(config_path)
	if os.IsNotExist(err) {
		log.Fatalf("Can't find file：%v", config_path)
	}
	return config_path
}
