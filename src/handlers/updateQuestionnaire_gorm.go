package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	. "serverless-crud-go/src/models"
)

func UpdateQuestionnaire(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("UpdateQuestionnaire")

	db := ConnectDB()
	defer db.Close()

	id, err := strconv.ParseUint(request.PathParameters["id"], 10, 32)
	if err != nil {
		panic(err)
	}
	questionnaire := Questionnaire{ID: uint(id)}
	json.Unmarshal([]byte(request.Body), &questionnaire)

	db.Save(&questionnaire)
	db.First(&questionnaire)

	body, _ := json.Marshal(&questionnaire)
	return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       string(body),
		},
		nil
}

func main() {
	lambda.Start(UpdateQuestionnaire)
}
