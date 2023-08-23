package email

import "my-email-tool/pkg/config"

func SendEmails(cfg *config.Config, mailer Mailer) error {
	for _, receiver := range cfg.Receivers {
		if err := mailer.SendEmail(cfg.SenderEmail, receiver.Email, cfg.Subject, cfg.FilePaths); err != nil {
			return err
		}
	}
	return nil
}
