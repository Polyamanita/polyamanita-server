package routes

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type Controller struct {
	S3          s3iface.S3API
	DynamoDB    dynamodbiface.DynamoDBAPI
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

	return &Controller{
		jwtKey:      []byte(jwtKey),
		environment: environment,
	}, nil
}
