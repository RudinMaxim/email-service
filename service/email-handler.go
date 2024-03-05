package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendEmailHandler(c *gin.Context) {
	var req EmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new GmailSender instance
	sender := NewEmailSender(
		"maxrudin",
		"maxrudin2004@gmail.com",
		"gdrv xhcs oapk dusm",
	)

	//sender := NewGmailSender(
	//	os.Getenv("EMAIL_SENDER_NAME"),
	//	os.Getenv("EMAIL_SENDER_ADDRESS"),
	//	os.Getenv("EMAIL_SENDER_PASSWORD"),
	//)

	// Send the email
	err := sender.SendEmail(
		req.Subject,
		req.Content,
		req.To,
		req.Cc,
		req.Bcc,
		req.AttachFiles,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}
