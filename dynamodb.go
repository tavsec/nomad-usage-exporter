package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	sess *session.Session
	svc  *dynamodb.DynamoDB
)

func InitDynamoDb() {
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc = dynamodb.New(sess)
}

func StoreResourceUsage(ru ResourceUsage) error {
	item, err := dynamodbattribute.MarshalMap(ru)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("NomadResources"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return err
	}
	return nil
}
