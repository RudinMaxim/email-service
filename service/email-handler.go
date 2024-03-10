package service

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// @Summary Отправить письмо
// @Description Отправляет электронное письмо
// @ID send-email
// @Accept  json
// @Produce  json
// @Param request body EmailRequest true "Email request"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /send-email [post]

func SendEmailHandler(c *gin.Context) {
	var req EmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sender := NewEmailSender(
		os.Getenv("EMAIL_SENDER_NAME"),
		os.Getenv("EMAIL_SENDER_ADDRESS"),
		os.Getenv("EMAIL_SENDER_PASSWORD"),
		os.Getenv("SMTP_AUTH_ADDRESS"),
		os.Getenv("SMTP_SERVER_ADDRESS"),
	)

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
