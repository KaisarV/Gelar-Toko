package gomail

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"time"

	model "GelarToko/models"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "GelarToko<nomen.test123@gmail.com>"
const CONFIG_AUTH_EMAIL = "nomen.test123@gmail.com"
const CONFIG_AUTH_PASSWORD = "tes12345"

var jwtKey = []byte("Jksdgbfkd334dsj")

type BodylinkEmail struct {
	Name  string
	Email string
	URL   string
}

func SendRegisterMail(email string, name string, id int) {
	token := GenerateVerifyToken(id)
	templateData := BodylinkEmail{
		Name:  name,
		Email: email,
		URL:   "http://localhost:8080/verify/" + token,
	}

	fmt.Println(token)

	result, _ := ParseTemplate("gomail/register.html", templateData)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", email, email)
	mailer.SetAddressHeader("Cc", email, "Pemberitahuan Verifikasi Akun")
	mailer.SetHeader("Subject", "Pemberitahuan Verifikasi Akun")
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

func GenerateVerifyToken(id int) (token string) {
	tokenExpiryTime := time.Now().Add(60 * time.Minute)

	claims := &model.Claims{
		ID:       id,
		Name:     "",
		UserType: 0,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}
	token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token2.SignedString(jwtKey)
	if err != nil {
		return
	}

	return signedToken
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

func SendBlockMail(email string, name string) {
	templateData := BodylinkEmail{
		Name: name,
	}

	result, _ := ParseTemplate("gomail/block.html", templateData)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", email, email)
	mailer.SetAddressHeader("Cc", email, "Pemberitahuan Block")
	mailer.SetHeader("Subject", "Pemberitahuan Block")
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

func SendUnblockMail(email string, name string) {
	templateData := BodylinkEmail{
		Name: name,
	}

	result, _ := ParseTemplate("gomail/unblock.html", templateData)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", email, email)
	mailer.SetAddressHeader("Cc", email, "Pemberitahuan Unblock")
	mailer.SetHeader("Subject", "Pemberitahuan Unblock")
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
