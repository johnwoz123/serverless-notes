package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"os"
)

var ddb *dynamodb.DynamoDB

type Note struct {
	UserId     string    `json:"id"`
	NoteId     uuid.UUID `json:"UUID"`
	Content    string    `json:"content"`
	Attachment string      `json:"attachment"`
	CreatedAt  string    `json:"created_at"`
}

func init() {
	region := os.Getenv("AWS_REGION")

	if session, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	}); err != nil {
		fmt.Println(fmt.Sprintf("Falied to connect to AWS: %s", err.Error()))
	} else {
		ddb = dynamodb.New(session)
	}
}

func addNote(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

}
