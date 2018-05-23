package main

import (
	"encoding/json"
	"fmt"
	// "strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/jmoiron/sqlx"
	. "serverless-crud-go/src/models"
)

func AddQuestion(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("AddQuestion")

	db, err := sqlx.Connect("postgres", "host=testisi.cuv7sthj4mjd.eu-central-1.rds.amazonaws.com user=root dbname=test1 password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Initialize question
	question := Question{}
	json.Unmarshal([]byte(request.Body), &question)

	// db.AutoMigrate(&question)

	// db.Create(&question)
	// db.First(&question)
	err = db.Get(&question, "INSERT INTO question (libraryid, questiontypeid, text) VALUES ($1, $2, $3) RETURNING *", question.Libraryid, question.Questiontypeid, question.Text)
	if err != nil {
		panic(err)
	}

	body, _ := json.Marshal(&question)
	return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       string(body),
		},
		nil
}

func main() {
	lambda.Start(AddQuestion)
}
