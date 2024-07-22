package services

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func Send(to, subject, htmlbody string) error {

	body, err := laodHtml(htmlbody)

	if err != nil {
		fmt.Println(err)
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", os.Getenv("MY_USER"))
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	dialer := gomail.NewDialer(os.Getenv("SMTP_HOST"), 587, os.Getenv("MY_USER"), os.Getenv("MY_PASS"))

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
