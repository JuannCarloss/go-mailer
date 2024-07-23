package main

import (
	"context"
	"encoding/json"
	"mailer/go-lambda/services"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestBody struct {
	Email     string `json:"email"`
	EmailType string `json:"emailType"`
}

func handler(context context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var emailBody RequestBody
	err := json.Unmarshal([]byte(event.Body), &emailBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 422,
			Body:       "Invalid JSON payload",
		}, nil
	}

	if emailBody.Email == "" || emailBody.EmailType == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 422,
			Body:       "email or emailType is blank or null",
		}, nil
	}

	err = services.SaveEmail(emailBody.Email, emailBody.EmailType)
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
