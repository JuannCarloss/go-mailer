package services

import (
	"fmt"
	"log"
	"mailer/go-lambda/config"
	domain "mailer/go-lambda/domain/emails"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/joho/godotenv"
)

var (
	sess = config.NewSession()
)

func SaveEmail(emailTo, emailType string) error {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	svc := dynamodb.New(sess)

	if emailType == "REGISTER" {

		email := domain.Emails{
			To:        emailTo,
			Timestamp: time.Now(),
		}
		av, err := dynamodbattribute.MarshalMap(email)
		if err != nil {
			log.Fatalf("Got error marshalling new email item: %s", err)
		}

		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(os.Getenv("table")),
		}

		Send(emailTo, "Thank you for Registering", "assets/ty-registering.html")

		_, err = svc.PutItem(input)
		if err != nil {
			log.Fatalf("Got error calling PutItem: %s", err)
		}
	} else if emailType == "ORDER" {
		email := domain.Emails{
			To:        emailTo,
			Timestamp: time.Now(),
		}

		av, err := dynamodbattribute.MarshalMap(email)
		if err != nil {
			log.Fatalf("Got error marshalling new email item: %s", err)
		}

		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(os.Getenv("table")),
		}

		Send(emailTo, "Thank you for Buying with us", "assets/ty-order.html")

		_, err = svc.PutItem(input)
		if err != nil {
			log.Fatalf("Got error calling PutItem: %s", err)
		}

	} else {
		log.Fatalf("invalid email type")
	}

	return nil
}
