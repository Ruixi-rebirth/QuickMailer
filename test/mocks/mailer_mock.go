package mocks

import "github.com/Ruixi-rebirth/QuickMailer/pkg/email"

type MockMailer struct {
	SentMessages []Message
}

type Message struct {
	From    string
	To      string
	Subject string
	Body    string
	Files   []string
}

func (m *MockMailer) SendEmail(from, to, subject, body string, filePaths []string) error {
	message := Message{
		From:    from,
		To:      to,
		Subject: subject,
		Body:    body,
		Files:   filePaths,
	}
	m.SentMessages = append(m.SentMessages, message)
	return nil
}

var _ email.Mailer = &MockMailer{} // Ensure MockMailer implements the Mailer interface
