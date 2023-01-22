package routes_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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

func TestGetCapturesList(t *testing.T) {
	t.Run("when the call is successful", func(t *testing.T) {
		fakeDynamo := fakes.DynamoDBAPI{}

		c := routes.NewTestController(nil, &fakeDynamo, nil, lib.NewLogger(os.Stdout))

		// Setup call
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(
			http.MethodGet,
			"/captures",
			io.NopCloser(strings.NewReader(`
			{
				"userid": "12321jkasdas"
			}
			`)),
		)
		ctx.Request.Header.Set("Content-Type", "application/json")

		// Make call
		c.GetCapturesList(ctx)

		// Validate that the response is correct
		assert.Equal(t, http.StatusOK, ctx.Writer.Status())

		gotResp, err := ioutil.ReadAll(w.Body)
		assert.NoError(t, err)
		assert.NotNil(t, string(gotResp))

		// Validate that the call made correct function calls
		gotTable := *fakeDynamo.QueryCall.Receives.QueryInput.TableName
		assert.Equal(t, "CapturedMushrooms", gotTable)

		gotUserid := fakeDynamo.QueryCall.Receives.QueryInput.ExpressionAttributeValues
		userid := map[string]*dynamodb.AttributeValue{":0": {S: aws.String("12321jkasdas")}}
		assert.Equal(t, gotUserid, userid)
	})

	t.Run("when the body request is invalid", func(t *testing.T) {
		c := routes.NewTestController(nil, nil, nil, lib.NewLogger(os.Stdout))

		// Setup call
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(
			http.MethodGet,
			"/captures",
			io.NopCloser(strings.NewReader(`
			"userid": "12321jkasdas"
			}
			`)),
		)
		ctx.Request.Header.Set("Content-Type", "application/json")

		// Make call
		c.GetCapturesList(ctx)

		// Validate that the response is correct
		assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
	})
}
