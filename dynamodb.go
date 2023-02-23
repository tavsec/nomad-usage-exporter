package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"os"
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

	filt := expression.Name("ID").Equal(expression.Value(ru.ID))
	proj := expression.NamesList(expression.Name("JobId"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		return err
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(os.Getenv("DYNAMODB_TABLE_NAME")),
	}

	result, err := svc.Scan(params)
	if err != nil {
		return err
	}

	if *result.Count > 0 {
		return nil
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(os.Getenv("DYNAMODB_TABLE_NAME")),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return err
	}
	return nil
}
