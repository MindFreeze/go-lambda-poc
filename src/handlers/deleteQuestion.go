package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"

	. "serverless-crud-go/src/models"
)

func DeleteQuestion(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("DeleteQuestion")

	db := ConnectDB()
	defer db.Close()

	id, err := strconv.ParseUint(request.PathParameters["id"], 10, 32)
	if err != nil {
		panic(err)
	}
	question := Question{ID: uint(id)}
	err = db.Delete(&question).Error

	if err != nil {
		return events.APIGatewayProxyResponse{ // Error HTTP response
			Body:       err.Error(),
			StatusCode: 500,
		}, nil
	} else {
		return events.APIGatewayProxyResponse{ // Success HTTP response
			StatusCode: 204,
		}, nil
	}
}

func main() {
	lambda.Start(DeleteQuestion)
}
