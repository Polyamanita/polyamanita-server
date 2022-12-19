package routes

import (
	"fmt"
	"os"
)

type Controller struct {
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
		jwtKey: []byte(jwtKey),
		environment: environment,
	}, nil
}
