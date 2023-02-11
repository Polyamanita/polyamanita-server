package routes_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
	"github.com/polyamanita/polyamanita-server/src/fakes"
	"github.com/polyamanita/polyamanita-server/src/lib"
	"github.com/polyamanita/polyamanita-server/src/models"
	"github.com/polyamanita/polyamanita-server/src/routes"
	"github.com/stretchr/testify/assert"
)

func TestGetCapture(t *testing.T) {
	wantCapture := &models.Capture{
		CaptureID:  "some-capture-id",
		UserID:     "some-user-id",
		TimesFound: 10,
		Notes:      "some notes",
		Instances: []*models.Instance{
			{
				Longitude: 101,
				Latitude:  102,
				Location:  "Forest",
				DateFound: time.Date(2022, 2, 22, 22, 22, 22, 22, time.UTC),
			},
		},
	}
	GetItemOutput, err := dynamodbattribute.Marshal(wantCapture)
	assert.NoError(t, err)
	dynamoMock := &fakes.DynamoDBAPI{}
	dynamoMock.GetItemCall.Returns.GetItemOutput = &dynamodb.GetItemOutput{
		Item: GetItemOutput.M,
	}

	cm := routes.NewTestController(nil, dynamoMock, nil, lib.NewLogger(os.Stdout))

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest(
		http.MethodGet,
		"/users/some-user-id/captures/some-capture-id",
		nil,
	)
	ctx.Params = []gin.Param{
		{Key: "UserID", Value: "some-user-id"},
		{Key: "CaptureID", Value: "some-capture-id"},
	}

	// Make call
	cm.GetCapture(ctx)

	// Check GetItem call
	assert.Equal(t,
		map[string]*dynamodb.AttributeValue{
			"UserID":   {S: aws.String("some-user-id")},
			"MainSort": {S: aws.String(fmt.Sprintf("Capture#%s#%s", "some-user-id", "some-capture-id"))},
		},
		dynamoMock.GetItemCall.Receives.GetItemInput.Key,
	)

	assert.Equal(t, aws.String("some-userbase-table"), dynamoMock.GetItemCall.Receives.GetItemInput.TableName)

	// Check body response
	gotBody := &struct {
		Capture *models.Capture `json:"capture"`
	}{}
	json.NewDecoder(w.Body).Decode(gotBody)

	assert.Equal(t, wantCapture, gotBody.Capture)
	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}

func TestAddCaptures(t *testing.T) {
	dynamoMock := &fakes.DynamoDBAPI{}
	cm := routes.NewTestController(nil, dynamoMock, nil, lib.NewLogger(os.Stdout))

	body := &struct {
		Captures []*struct {
			Notes     string             `json:"notes"`
			CaptureID string             `json:"captureID"`
			Instances []*models.Instance `json:"instances"`
		} `json:"captures"`
	}{
		Captures: []*struct {
			Notes     string             `json:"notes"`
			CaptureID string             `json:"captureID"`
			Instances []*models.Instance `json:"instances"`
		}{
			{
				Notes:     "some notes",
				CaptureID: "some-capture-id",
				Instances: []*models.Instance{
					{
						Longitude: 123,
						Latitude:  456,
						Location:  "somewhere",
						S3Key:     "some-key",
					},
				},
			},
		},
	}
	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(body)

	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest(
		http.MethodGet,
		"/users/some-user-id/captures",
		b,
	)
	ctx.Params = []gin.Param{
		{Key: "UserID", Value: "some-user-id"},
	}

	// Make call
	cm.AddCaptures(ctx)

	// Check GetItem call
	assert.Equal(t,
		map[string]*dynamodb.AttributeValue{
			"UserID":   {S: aws.String("some-user-id")},
			"MainSort": {S: aws.String(fmt.Sprintf("Capture#%s#%s", "some-user-id", "some-capture-id"))},
		},
		dynamoMock.UpdateItemCall.Receives.UpdateItemInput.Key,
	)

	assert.Equal(t, aws.String("some-userbase-table"), dynamoMock.UpdateItemCall.Receives.UpdateItemInput.TableName)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}
