package services

import (
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

	godotenv.Load()

	svc := dynamodb.New(sess)

	email := domain.Emails{
		To:        emailTo,
		Type:      emailType,
		Timestamp: time.Now(),
	}
	av, err := dynamodbattribute.MarshalMap(email)
	if err != nil {
		log.Fatalf("Got error marshalling new email item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(os.Getenv("TABLE")),
	}

	if emailType == "REGISTER" {
		Send(emailTo, "Thank you for Registering", "assets/ty-registering.html")

	} else if emailType == "ORDER" {
		Send(emailTo, "Thank you for Buying with us", "assets/ty-order.html")

	} else {
		log.Println("invalid email type")
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}
	return nil
}
