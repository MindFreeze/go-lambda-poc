package main

import (
	"encoding/json"
	"fmt"
	// "strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	. "serverless-crud-go/src/models"
)

func AddQuestionnaire(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("AddQuestionnaire")

	db := ConnectDB()
	defer db.Close()

	questionnaire := Questionnaire{}
	json.Unmarshal([]byte(request.Body), &questionnaire)

	// db.AutoMigrate(&questionnaire)

	err := db.Create(&questionnaire).Error
	if err != nil {
		panic(err)
	}
	err = db.First(&questionnaire).Error
	if err != nil {
		panic(err)
	}

	body, _ := json.Marshal(&questionnaire)
	return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       string(body),
		},
		nil
}

func main() {
	lambda.Start(AddQuestionnaire)
}
