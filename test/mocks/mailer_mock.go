package mocks

import "github.com/Ruixi-rebirth/QuickMailer/pkg/email"

type MockMailer struct {
	SentMessages []Message
}

type Message struct {
	From    string
	To      string
	Subject string
	Files   []string
}

func (m *MockMailer) SendEmail(from, to, subject string, filePaths []string) error {
	message := Message{
		From:    from,
		To:      to,
		Subject: subject,
		Files:   filePaths,
	}
	m.SentMessages = append(m.SentMessages, message)
	return nil
}

var _ email.Mailer = &MockMailer{} // Ensure MockMailer implements the Mailer interface
