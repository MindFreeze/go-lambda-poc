package main

import (
	"encoding/json"
	// "fmt"
	// "strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	. "serverless-crud-go/src/models"
)

func ListQuestionnaires(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	db := ConnectDB()
	defer db.Close()

	all := []Questionnaire{}
	db.Find(&all)

	// Success HTTP response
	body, _ := json.Marshal(&all)
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(ListQuestionnaires)
}
