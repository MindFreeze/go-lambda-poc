package main

import (
	"encoding/json"
	// "strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	. "serverless-crud-go/src/models"
)

func AddLibrary(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db := ConnectDB()
	defer db.Close()

	library := Library{}
	json.Unmarshal([]byte(request.Body), &library)

	// db.AutoMigrate(&library)

	db.Create(&library)
	db.First(&library)

	body, _ := json.Marshal(&library)
	return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       string(body),
		},
		nil
}

func main() {
	lambda.Start(AddLibrary)
}
