package models

import "time"

type Capture struct {
	UserID    string `json:"userID" dynamodbav:"UserID"`
	CaptureID string `json:"captureID" dynamodbav:"CaptureID"`

	TimesFound int64       `json:"timesFound" dynamodbav:"TimesFound"`
	Notes      string      `json:"notes" dynamodbav:"Notes"`
	Instances  []*Instance `json:"instances" dynamodbav:"Instances"`
}

type Instance struct {
	Longitude float64   `json:"longitude" dynamodbav:"Longitude"`
	Latitude  float64   `json:"latitude" dynamodbav:"Latitude"`
	Location  string    `json:"location" dynamodbav:"Location"`
	DateFound time.Time `json:"dateFound" dynamodbav:"DateFound"`
	S3Key     string    `json:"s3Key" dynamodbav:"S3Key"`
}
