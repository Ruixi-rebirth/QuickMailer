package main

import (
	"log"
	"my-email-tool/pkg/config"
	"my-email-tool/pkg/email"
	"my-email-tool/pkg/utils"
)

func main() {
	cfg, err := config.LoadConfig(utils.Current_Dir_Config_Path())
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	realMailer := email.NewRealMailer("smtp.qq.com", 587, cfg.SenderEmail, cfg.SenderPassword)
	if err := email.SendEmails(cfg, realMailer); err != nil {
		log.Fatalf("Failed to send emails: %v", err)
	}
}
