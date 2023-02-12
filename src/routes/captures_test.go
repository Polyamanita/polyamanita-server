package routes_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
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
				ImageLink: "https://polyamanita-images/image",
			},
		},
	}
	GetItemOutput, err := dynamodbattribute.Marshal(wantCapture)
	assert.NoError(t, err)
	dynamoMock := &fakes.DynamoDBAPI{}
	dynamoMock.GetItemCall.Returns.GetItemOutput = &dynamodb.GetItemOutput{
		Item: GetItemOutput.M,
	}

	s3Mock := &fakes.S3API{}
	url, _ := url.Parse("https://polyamanita-images/image")
	s3Mock.GetObjectRequestCall.Returns.Request = &request.Request{
		Operation: &request.Operation{},
		HTTPRequest: &http.Request{
			URL: url,
		},
	}

	cm := routes.NewTestController(s3Mock, dynamoMock, nil, lib.NewLogger(os.Stdout))

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
			"MainSort": {S: aws.String("Metadata")},
		},
		dynamoMock.UpdateItemCall.Receives.UpdateItemInput.Key,
	)

	assert.Equal(t, aws.String("some-userbase-table"), dynamoMock.UpdateItemCall.Receives.UpdateItemInput.TableName)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}

func TestGetCapturesList(t *testing.T) {
	t.Run("when the call is successful", func(t *testing.T) {
		wantCapture := &[]models.Capture{{
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
		}}

		queryOutputItems, err := dynamodbattribute.MarshalMap(wantCapture)
		assert.NoError(t, err)
		dynamoMock := &fakes.DynamoDBAPI{}
		dynamoMock.QueryCall.Returns.QueryOutput = &dynamodb.QueryOutput{
			Count: aws.Int64(1),
			Items: []map[string]*dynamodb.AttributeValue{
				queryOutputItems,
			},
		}

		c := routes.NewTestController(nil, dynamoMock, nil, lib.NewLogger(os.Stdout))

		// Setup call
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(
			http.MethodGet,
			"/users/some-id/captures",
			nil,
		)
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Params = gin.Params{{Key: "UserID", Value: "some-id"}}

		// Make call
		c.GetCapturesList(ctx)

		// Validate that the response is correct
		assert.Equal(t, http.StatusOK, ctx.Writer.Status())

		gotResp, err := ioutil.ReadAll(w.Body)
		assert.NoError(t, err)
		assert.NotNil(t, string(gotResp))

		// Validate that the call made correct function calls
		assert.Equal(t, "some-userbase-table",
			*dynamoMock.QueryCall.Receives.QueryInput.TableName)

	})
}

func TestUploadCaptureImage(t *testing.T) {
	t.Run("when the call is successful", func(t *testing.T) {
		s3Mock := &fakes.S3API{}
		url, _ := url.Parse("https://polyamanita-images/some-image-upload-url")
		s3Mock.PutObjectRequestCall.Returns.Request = &request.Request{
			Operation: &request.Operation{},
			HTTPRequest: &http.Request{
				URL: url,
			},
		}

		cm := routes.NewTestController(s3Mock, nil, nil, lib.NewLogger(os.Stdout))

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(
			http.MethodPost,
			"/users/some-user-id/captures/images?numLinks=2",
			nil,
		)
		ctx.Params = []gin.Param{
			{Key: "UserID", Value: "some-user-id"},
		}

		// Make call
		cm.UploadCaptureImage(ctx)

		// Check body response
		type Link struct {
			UploadLink string `json:"uploadLink"`
			S3Key      string `json:"s3Key"`
		}
		type UploadCaptureImageOutputStruct struct {
			Links []*Link `json:"links"`
		}
		gotBody := &UploadCaptureImageOutputStruct{}
		json.NewDecoder(w.Body).Decode(gotBody)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())

		assert.Len(t, gotBody.Links, 2)
		assert.Equal(t, "https://polyamanita-images/some-image-upload-url", gotBody.Links[0].UploadLink)
		assert.Contains(t, gotBody.Links[0].S3Key, "some-user-id")
		assert.NotEmpty(t, gotBody.Links[1].S3Key)
	})
	t.Run("when the numLinks is unspecified", func(t *testing.T) {
		s3Mock := &fakes.S3API{}
		url, _ := url.Parse("https://polyamanita-images/some-image-upload-url")
		s3Mock.PutObjectRequestCall.Returns.Request = &request.Request{
			Operation: &request.Operation{},
			HTTPRequest: &http.Request{
				URL: url,
			},
		}

		cm := routes.NewTestController(s3Mock, nil, nil, lib.NewLogger(os.Stdout))

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(
			http.MethodPost,
			"/users/some-user-id/captures/images",
			nil,
		)
		ctx.Params = []gin.Param{
			{Key: "UserID", Value: "some-user-id"},
		}

		// Make call
		cm.UploadCaptureImage(ctx)

		// Check body response
		type Link struct {
			UploadLink string `json:"uploadLink"`
			S3Key      string `json:"s3Key"`
		}
		type UploadCaptureImageOutputStruct struct {
			Links []*Link `json:"links"`
		}
		gotBody := &UploadCaptureImageOutputStruct{}
		json.NewDecoder(w.Body).Decode(gotBody)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())

		assert.Len(t, gotBody.Links, 1)
		assert.Equal(t, "https://polyamanita-images/some-image-upload-url", gotBody.Links[0].UploadLink)
		assert.Contains(t, gotBody.Links[0].S3Key, "some-user-id/")
	})
}
