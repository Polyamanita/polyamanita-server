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
	S3       s3iface.S3API
	DynamoDB dynamodbiface.DynamoDBAPI
	Mail     mail.MailIFace

	secrets Secrets
	l       *lib.Logger
}

type Secrets struct {
	verificationTable string
	jwtKey            []byte
	environment       string
}

func NewController(l *lib.Logger) (*Controller, error) {
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

	mailClient, err := mail.New(l)
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
		Mail: mailClient,
		l:    l,
	}, nil
}

func NewTestController(S3 s3iface.S3API, DynamoDB dynamodbiface.DynamoDBAPI, Mail mail.MailIFace, l *lib.Logger) *Controller {
	return &Controller{
		S3:       S3,
		DynamoDB: DynamoDB,
		Mail:     Mail,
		secrets: Secrets{
			verificationTable: "some-table",
			jwtKey:            []byte("some-key"),
			environment:       "some-env",
		},
		l: l,
	}
}
