package gomail

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "GelarToko<nomen.test123@gmail.com>"
const CONFIG_AUTH_EMAIL = "nomen.test123@gmail.com"
const CONFIG_AUTH_PASSWORD = "tes12345"

type BodylinkEmail struct {
	Name  string
	Email string
	URL   string
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println(err)
		return "", err
	}

	return buf.String(), nil
}

func SendLoginMail(email string, name string) {
	templateData := BodylinkEmail{
		Name: name,
	}

	result, _ := ParseTemplate("gomail/login.html", templateData)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", email, email)
	mailer.SetAddressHeader("Cc", email, "Pemberitahuan Login")
	mailer.SetHeader("Subject", "Pemberitahuan Login")
	mailer.SetBody("text/html", result)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

}
