package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/jmoiron/sqlx"
	. "serverless-crud-go/src/models"
	// "github.com/Masterminds/squirrel"
)

func UpdateQuestionnaire(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("UpdateQuestionnaire")

	db, err := sqlx.Connect("postgres", "host=testisi.cuv7sthj4mjd.eu-central-1.rds.amazonaws.com user=root dbname=test1 password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	id, err := strconv.ParseUint(request.PathParameters["id"], 10, 32)
	if err != nil {
		panic(err)
	}
	questionnaire := Questionnaire{ID: uint(id)}
	json.Unmarshal([]byte(request.Body), &questionnaire)

	db.Get(&questionnaire, "UPDATE questionnaires SET name = $2 WHERE id = $1 RETURNING *", id, questionnaire.Name)

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
