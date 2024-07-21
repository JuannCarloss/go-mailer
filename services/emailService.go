package services

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func Send(to, subject, htmlbody string) error {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	body, err := laodHtml(htmlbody)

	if err != nil {
		fmt.Println(err)
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", os.Getenv("user"))
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	dialer := gomail.NewDialer(os.Getenv("host"), 587, os.Getenv("user"), os.Getenv("password"))

	if err := dialer.DialAndSend(msg); err != nil {
		return err
	} else {
		return nil
	}
}

func laodHtml(filepath string) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
