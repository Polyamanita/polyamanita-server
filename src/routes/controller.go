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
	ddbVerificationTable string
	ddbUserbaseTable     string
	s3StoreBucket        string
	jwtKey               []byte
	environment          string
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

	ddbUserbaseTable, ok := os.LookupEnv("AWS_DDB_TABLE_USERBASE")
	if !ok {
		return nil, fmt.Errorf("unable to find env var AWS_DDB_TABLE_USERBASE")
	}

	ddbVerificationTable, ok := os.LookupEnv("AWS_DDB_TABLE_VERIFICATIONCODES")
	if !ok {
		return nil, fmt.Errorf("unable to find env var AWS_DDB_TABLE_VERIFICATIONCODES")
	}

	s3StoreTable, ok := os.LookupEnv("AWS_S3_BUCKET_POLYAMANITA_STORE")
	if !ok {
		return nil, fmt.Errorf("unable to find env var AWS_S3_BUCKET_POLYAMANITA_STORE")
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
			jwtKey:               []byte(jwtKey),
			environment:          environment,
			s3StoreBucket:        s3StoreTable,
			ddbVerificationTable: ddbVerificationTable,
			ddbUserbaseTable:     ddbUserbaseTable,
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
			ddbVerificationTable: "some-verification-table",
			ddbUserbaseTable:     "some-userbase-table",
			s3StoreBucket:        "some-s3store-bucket",
			jwtKey:               []byte("some-key"),
			environment:          "some-env",
		},
		l: l,
	}
}
