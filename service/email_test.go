package service

import (
	"os"
	"testing"
)

func TestSendEmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	sender := NewEmailSender(
		os.Getenv("EMAIL_SENDER_NAME"),
		os.Getenv("EMAIL_SENDER_ADDRESS"),
		os.Getenv("EMAIL_SENDER_PASSWORD"),
		os.Getenv("SMTP_AUTH_ADDRESS"),
		os.Getenv("SMTP_SERVER_ADDRESS"),
	)

	err := sender.SendEmail(
		"Test Subject",
		`
			<h1>Привет!</h1>
		`,
		[]string{"maxrdin2004@gmil.com"},
		nil,
		nil,
		nil,
	)
	if err != nil {
		t.Errorf("SendEmail failed: %v", err)
	}
}
