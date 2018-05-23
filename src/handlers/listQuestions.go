package main

import (
	"encoding/json"
	"fmt"
	// "strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	. "serverless-crud-go/src/models"
)

func ListQuestions(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("ListQuestions")

	db := ConnectDB()
	defer db.Close()

	// Initialize question
	questions := []Question{}
	db.Find(&questions)

	// Success HTTP response
	body, _ := json.Marshal(&questions)
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(ListQuestions)
}
