package email

// Mailer 接口，定义邮件发送功能
type Mailer interface {
	SendEmail(from, to, subject, body string, filePaths []string) error
}
