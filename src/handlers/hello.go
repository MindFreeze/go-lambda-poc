package main

import (
	// "fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Hello(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       "Hello " + request.QueryStringParameters["name"],
		},
		nil
}

func main() {
	lambda.Start(Hello)
}
