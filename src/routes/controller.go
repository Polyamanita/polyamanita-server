package routes

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/polyamanita/polyamanita-server/src/clients/mail"
	"github.com/polyamanita/polyamanita-server/src/lib"
)

type Controller struct {
	S3         s3iface.S3API
	DynamoDB   dynamodbiface.DynamoDBAPI
	secrets    Secrets
	MailClient mail.IFace
	l          *lib.Logger
}

type Secrets struct {
	verificationTable string
	jwtKey            []byte
	environment       string
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
		S3:       s3.New(sess),
		DynamoDB: dynamodb.New(sess),
		secrets: Secrets{
			jwtKey:            []byte(jwtKey),
			environment:       environment,
			verificationTable: "",
		},
		l: lib.NewLogger(os.Stdout),
	}, nil
}
