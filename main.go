package main

import (

	// "encoding/json"
	// "fmt"
	// "log"
	"context"
	"mailer/go-lambda/services"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestBody struct {
	Email     string `json:"email"`
	EmailType string `json:"emailType"`
}

func handler(context context.Context, email RequestBody) (events.APIGatewayProxyResponse, error) {

	if email.Email == "" || email.EmailType == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 422,
			Body:       "email or emailType is blank or null",
		}, nil
	}

	err := services.SaveEmail(email.Email, email.EmailType)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 422,
			Body:       "got error while saving the email",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Body:       "email sent",
	}, nil
}

func main() {
	lambda.Start(handler)
}
