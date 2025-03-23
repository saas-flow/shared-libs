package providers

import (
	"fmt"
	"net/smtp"
)

type SMTProvider struct {
	Host     string
	Port     string
	Username string
	Password string
}

// Send
// Params
//   - to: []string
//   - subject: string
//   - body: string
func (s *SMTProvider) Send(to []string, subject, body string) error {
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)

	// Format pesan email
	msg := fmt.Sprintf("Subject: %s\n\n%s", subject, body)

	err := smtp.SendMail(s.Host+":"+s.Port, auth, s.Username, to, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}
