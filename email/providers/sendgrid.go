package providers

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendGridProvider struct
type SendGridProvider struct {
	APIKey string
}

// Send method untuk mengirim email dengan SendGrid
func (s *SendGridProvider) Send(to []string, from *mail.Email, subject, body, html string) error {
	client := sendgrid.NewSendClient(s.APIKey)

	// Konversi ke format SendGrid
	recipients := []*mail.Email{}
	for _, recipient := range to {
		recipients = append(recipients, mail.NewEmail("", recipient))
	}

	message := mail.NewSingleEmail(from, subject, recipients[0], body, html)
	response, err := client.Send(message)
	if err != nil {
		return err
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("SendGrid error: %d - %s", response.StatusCode, response.Body)
	}

	return nil
}
