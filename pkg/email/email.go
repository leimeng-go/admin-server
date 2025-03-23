package email

import (
	"fmt"
	"net/smtp"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	FromName string
}

type Email struct {
	config Config
}

func NewEmail(config Config) *Email {
	return &Email{
		config: config,
	}
}

func (e *Email) SendVerificationCode(to string, code string) error {
	subject := "验证码"
	body := fmt.Sprintf(`
		<html>
		<body>
			<h3>您的验证码是：%s</h3>
			<p>验证码有效期为5分钟，请尽快使用。</p>
			<p>如果这不是您的操作，请忽略此邮件。</p>
		</body>
		</html>
	`, code)

	message := fmt.Sprintf("To: %s\r\n"+
		"From: %s <%s>\r\n"+
		"Subject: %s\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n"+
		"\r\n%s",
		to, e.config.FromName, e.config.From, subject, body)

	auth := smtp.PlainAuth("", e.config.Username, e.config.Password,
		e.config.Host)

	return smtp.SendMail(
		fmt.Sprintf("%s:%d", e.config.Host, e.config.Port),
		auth,
		e.config.From,
		[]string{to},
		[]byte(message),
	)
}
