package email

import (
	"gopkg.in/gomail.v2"
	"mime"
	"path/filepath"
)

type RealMailer struct {
	Dialer *gomail.Dialer
}

func NewRealMailer(host string, port int, username, password string) *RealMailer {
	return &RealMailer{
		Dialer: gomail.NewDialer(host, port, username, password),
	}
}

func (r *RealMailer) SendEmail(from, to, subject, body string, filePaths []string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	if body != "" {
		m.SetBody("text/plain", body)
	}

	for _, path := range filePaths {
		_, fileName := filepath.Split(path) // 提取文件名
		encodedFilename := mime.QEncoding.Encode("utf-8", fileName)
		m.Attach(path, gomail.Rename(encodedFilename))
	}
	return r.Dialer.DialAndSend(m)
}
