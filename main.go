package main

import (
	//"mailer/go-lambda/config"
	"log"
	"mailer/go-lambda/services"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayV2HTTPRequest) events.APIGatewayV2HTTPResponse {
	email := request.QueryStringParameters["email"]
	emailType := request.QueryStringParameters["emailType"]

	err := services.SaveEmail(email, emailType)
	if err != nil {
		log.Printf("Error while saving the email: %v", err)
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       "Error while saving the email",
		}
	}
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       "email sent",
	}
}

func main() {
	//config.CreateTableEmails()
	//services.Send()
	//services.SaveEmail("", "REGISTER")

	lambda.Start(handler)
}
