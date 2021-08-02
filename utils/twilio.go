package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"travel/infrastructure"
)

type Twilio struct {
	logger infrastructure.Logger
	env    infrastructure.Env
}

func NewTwilio(logger infrastructure.Logger, env infrastructure.Env) Twilio {
	return Twilio{
		logger: logger,
		env:    env,
	}
}

func (t Twilio) SendSms(to string, message string) {

	// Set account keys & information
	accountSid := t.env.TWILIO_SID
	authToken := t.env.TWILIO_authToken
	from := t.env.TWILIO_FROM

	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages"

	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", from)
	msgData.Set("Body", message)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["accountSid"])
		}
		log.Print("------------Sent sms successfully------------")
	} else {
		fmt.Println(resp.Status)
		t.logger.Zap.Error("---------> Error sending sms <---------")
	}
}
