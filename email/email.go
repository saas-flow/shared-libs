package email

import (
	"text/template"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailProvider interface {
	Send([]string, *mail.Email, string, string) error
}

func LoadTemplates() (*template.Template, error) {
	return template.ParseFiles("./email_template.tmpl")
}
