package models

import "time"

type Capture struct {
	UserID    string `json:"userID,omitempty" dynamodbav:"UserID"`
	CaptureID string `json:"captureID,omitempty" dynamodbav:"CaptureID"`

	TimesFound int64       `json:"timesFound" dynamodbav:"TimesFound"`
	Notes      string      `json:"notes,omitempty" dynamodbav:"Notes"`
	Instances  []*Instance `json:"instances,omitempty" dynamodbav:"Instances"`
}

type Instance struct {
	Longitude float64   `json:"longitude" dynamodbav:"Longitude"`
	Latitude  float64   `json:"latitude" dynamodbav:"Latitude"`
	Location  string    `json:"location" dynamodbav:"Location"`
	DateFound time.Time `json:"dateFound" dynamodbav:"DateFound"`
	S3Key     string    `json:"s3Key" dynamodbav:"S3Key"`
	ImageLink string `json:"imageLink,omitempty"`
}
