package emailjs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type EmailJS struct {
	UserID    string
	ServiceID string
}

func NewEmailJS(userID string, serviceID string) EmailJS {
	emailJS := EmailJS{
		UserID:    userID,
		ServiceID: serviceID,
	}

	return emailJS
}

func (ejs *EmailJS) SendMail(templateId string, templateParams interface{}) {
	apiURL := "https://api.emailjs.com/api/v1.0/email/send"

	body := SendMailRequestBody{
		ServiceID:      ejs.ServiceID,
		UserID:         ejs.UserID,
		TemplateID:     templateId,
		TemplateParams: templateParams,
	}

	jsonBody, err := json.Marshal(body)
	_, err = http.Post(apiURL, "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		fmt.Println(err)
	}
}
