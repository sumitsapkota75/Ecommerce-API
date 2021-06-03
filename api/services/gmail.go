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

func (e GmailService) SendGmail(emailParams models.EmailParams) error {
	var msgString string
	to := "To: " + emailParams.To + "\r\n"
	msgString = to
	subject := "Subject: " + emailParams.Subject
	msgString = msgString + subject + "\r\n"
	msgString = msgString + "\n" + emailParams.Body

	msg := gmail.Message{
		Raw: base64.URLEncoding.EncodeToString([]byte(msgString)),
	}
	_, err := e.gmailService.Users.Messages.Send("me", &msg).Do()
	if err != nil {
		return err
	}

	return nil
}
