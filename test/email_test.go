package test

import (
	"my-email-tool/pkg/config"
	"my-email-tool/pkg/email"
	"my-email-tool/test/mocks"
	"testing"
)

func TestSendEmails(t *testing.T) {
	cfg := &config.Config{
		SenderEmail: "test@example.com",
		Subject:     "Test",
		FilePaths:   []string{"path1.txt"},
		Receivers:   []config.Receiver{{Name: "John", Email: "john@example.com"}},
	}

	mockMailer := &mocks.MockMailer{}
	err := email.SendEmails(cfg, mockMailer)

	if err != nil {
		t.Errorf("Failed to send emails: %v", err)
	}

	if len(mockMailer.SentMessages) != 1 {
		t.Errorf("Expected 1 email to be sent, but got %d", len(mockMailer.SentMessages))
	}
}
