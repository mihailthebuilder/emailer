package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func handleRequestToSendEmail(c *gin.Context) {

	var body EmailRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("error parsing request body:", err)
		c.String(http.StatusBadRequest, "error parsing request body")
	}

	emailContent := fmt.Sprintf("Received email from: %s", body.EmailAddress)
	from := mail.NewEmail("Sender", getEnv("EMAIL_SENDER"))
	subject := body.Title
	to := mail.NewEmail("Recipient", getEnv("EMAIL_RECIPIENT"))
	message := mail.NewSingleEmailPlainText(from, subject, to, emailContent)
	client := sendgrid.NewSendClient(getEnv("EMAIL_API_KEY"))

	response, err := client.Send(message)
	if err != nil {
		log.Panic("error sending email: ", err)
	}

	if response.StatusCode != 202 {
		log.Panic("non-200 status code when sending email: ", response.StatusCode, "; status text: ", response.Body)
	}

	c.Status(http.StatusOK)
}

type EmailRequestBody struct {
	EmailAddress string
	Title        string
}
