package services

import (
	"log"
	"net/smtp"
	"travel/infrastructure"
	"travel/models"
	"travel/utils"
)

//SMTPService -> struct
type SMTPService struct {
	logger infrastructure.Logger
	env    infrastructure.Env
}

//NewSMTPService -> constructor
func NewSMTPService(logger infrastructure.Logger, env infrastructure.Env) SMTPService {
	return SMTPService{
		logger: logger,
		env:    env,
	}
}

func (s *SMTPService) SendMail(params models.EmailParams) (bool, error) {
	// Create authentication
	smtpHost := s.env.SMTPHost
	smtpPort := s.env.SMTPPort
	auth := smtp.PlainAuth("", params.From, s.env.SMTPPassword, s.env.SMTPHost)
	emailSubject, err := utils.ParseTemplate(params.SubjectTemplate, params.SubjectData)
	if err != nil {
		return false, err
	}
	emailBody, cerr := utils.ParseTemplate(params.BodyTemplate, params.BodyData)
	if cerr != nil {
		return false, err
	}
	var msgString string
	subject := "Subject: " + emailSubject
	msgString = msgString + subject + "\r\n"
	msgString = msgString + "\r\n" + emailBody
	msg := []byte(msgString)
	to := []string{params.To}
	// Send actual message
	sendMailerr := smtp.SendMail(smtpHost+":"+smtpPort, auth, params.From, to, msg)
	if sendMailerr != nil {
		log.Fatal(err)
		return false, err
	}
	s.logger.Zap.Info("Mail sent successfully")
	return true, nil
}
