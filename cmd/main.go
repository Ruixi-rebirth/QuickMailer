package main

import (
	"github.com/Ruixi-rebirth/QuickMailer/pkg/config"
	"github.com/Ruixi-rebirth/QuickMailer/pkg/email"
	"github.com/Ruixi-rebirth/QuickMailer/pkg/utils"
	"log"
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
