package email

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestSendEmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	
	errLoad := godotenv.Load()
	if errLoad != nil {
		log.Fatal("Error loading .env file")
	}

	sender := NewGmailSender(
		os.Getenv("EMAIL_SENDER_NAME"),
		os.Getenv("EMAIL_SENDER_ADDRESS"),
		os.Getenv("EMAIL_SENDER_PASSWORD"),
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
