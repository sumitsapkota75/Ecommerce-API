package services

import (
	"encoding/base64"
	"travel/infrastructure"
	"travel/models"

	"google.golang.org/api/gmail/v1"
)

type GmailService struct {
	logger       infrastructure.Logger
	gmailService *gmail.Service
}

func NewGmailService(
	logger infrastructure.Logger,
	gmailService *gmail.Service,
) GmailService {
	return GmailService{
		logger:       logger,
		gmailService: gmailService,
	}
}

func (e *GmailService) SendGmail(emailParams models.EmailParams) error {
	var msgString string
	to := emailParams.To
	msgString = to + "\r\n"
	subject := emailParams.Subject + "\r\n"
	msgString = msgString + subject
	body := emailParams.Body
	msgString = msgString + "\r\n" + body

	msg := gmail.Message{
		Raw: base64.StdEncoding.EncodeToString([]byte(msgString)),
	}

	_, err := e.gmailService.Users.Messages.Send("me", &msg).Do()
	if err != nil {
		return err
	}

	return nil
}
