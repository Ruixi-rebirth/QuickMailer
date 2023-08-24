package email

import (
	"fmt"
	"github.com/Ruixi-rebirth/QuickMailer/pkg/config"
)

func SendEmails(cfg *config.Config, mailer Mailer) error {
	for _, receiver := range cfg.Receivers {
		if err := mailer.SendEmail(cfg.SenderEmail, receiver.Email, cfg.Subject,cfg.Body ,cfg.FilePaths); err != nil {
			return err
		}
		fmt.Printf("Email sent to %s (%s)\n", receiver.Name, receiver.Email)
	}
	return nil
}
