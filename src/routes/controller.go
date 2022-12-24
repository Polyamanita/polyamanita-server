package routes

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
)

type Controller struct {
	SES         sesiface.SESAPI
	S3          s3iface.S3API
	DynamoDB    dynamodbiface.DynamoDBAPI
	l           *log.Logger
	jwtKey      []byte
	environment string
}

func NewController() (*Controller, error) {
	jwtKey, ok := os.LookupEnv("JWT_KEY")
	if !ok {
		return nil, fmt.Errorf("unable to find env var JWT_KEY")
	}

	environment, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		return nil, fmt.Errorf("unable to find env var ENVIRONMENT")
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		return nil, err
	}

	return &Controller{
		SES:         ses.New(sess),
		S3:          s3.New(sess),
		DynamoDB:    dynamodb.New(sess),
		jwtKey:      []byte(jwtKey),
		l:           log.New(os.Stdout, "[DEBUG] ", 0),
		environment: environment,
	}, nil
}
