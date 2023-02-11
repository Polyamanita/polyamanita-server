package routes_test

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/polyamanita/polyamanita-server/src/fakes"
	"github.com/polyamanita/polyamanita-server/src/lib"
	"github.com/polyamanita/polyamanita-server/src/routes"
	"github.com/stretchr/testify/assert"
)

func TestSearchUser(t *testing.T) {
	t.Run("when the call is successful", func(t *testing.T) {
		fakeDynamo := fakes.DynamoDBAPI{}
		fakeDynamo.QueryCall.Returns.QueryOutput = &dynamodb.QueryOutput{
			Items: []map[string]*dynamodb.AttributeValue{{}},
		}

		c := routes.NewTestController(nil, &fakeDynamo, nil, lib.NewLogger(os.Stdout))

		// Setup call
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(
			http.MethodGet,
			"/users",
			io.NopCloser(strings.NewReader(`
			{
				"username": "kyle25"
			}
			`)),
		)
		ctx.Request.Header.Set("Content-Type", "application/json")

		// Make call
		c.SearchUser(ctx)

		// Validate that the response is correct
		assert.Equal(t, http.StatusOK, ctx.Writer.Status())

		_, err := ioutil.ReadAll(w.Body)
		assert.NoError(t, err)

		// Validate the correct table
		gotTable := *fakeDynamo.QueryCall.Receives.QueryInput.TableName
		assert.Equal(t, "some-userbase-table", gotTable)

		// Validate the query input
		gotUsername := fakeDynamo.QueryCall.Receives.QueryInput.ExpressionAttributeValues
		username := map[string]*dynamodb.AttributeValue{":0": {S: aws.String("kyle25")}}
		assert.Equal(t, gotUsername, username)

		//validate the query output

	})

}

func TestRegisterUser(t *testing.T) {
	t.Run("when the call is successful", func(t *testing.T) {
		fakeDynamo := fakes.DynamoDBAPI{}
		fakeDynamo.ScanCall.Stub = func(si *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
			if *si.TableName == "some-verification-table" {
				return &dynamodb.ScanOutput{
					Items: []map[string]*dynamodb.AttributeValue{
						{
							"codeExpiry": {S: aws.String("9999-01-30T18:48:49-05:00")},
						},
					},
					Count: aws.Int64(1),
				}, nil
			}
			if *si.TableName == "some-userbase-table" {
				return &dynamodb.ScanOutput{
					Count: aws.Int64(0),
				}, nil
			}

			return nil, nil
		}
		fakeDynamo.ScanCall.Returns.ScanOutput = &dynamodb.ScanOutput{
			Items: []map[string]*dynamodb.AttributeValue{},
		}

		c := routes.NewTestController(nil, &fakeDynamo, nil, lib.NewLogger(os.Stdout))

		// Setup call
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(
			http.MethodPost,
			"/users",
			io.NopCloser(strings.NewReader(`
			{
				"code": "asdafwad12421",
				"username": "kyle25",
				"firstname": "kyle",
				"lastname": "boitz",
				"password": "password12!",
				"email": "kyle@gmail.com"
			}
			`)),
		)
		ctx.Request.Header.Set("Content-Type", "application/json")

		// Make call
		c.RegisterUser(ctx)

		// Validate that the response is correct
		assert.Equal(t, http.StatusCreated, ctx.Writer.Status())

		_, err := ioutil.ReadAll(w.Body)
		assert.NoError(t, err)

		// Validate the correct tables
		gotTableCode := *fakeDynamo.ScanCall.Receives.ScanInput.TableName
		assert.Equal(t, "some-userbase-table", gotTableCode)
		gotTableUser := *fakeDynamo.PutItemCall.Receives.PutItemInput.TableName
		assert.Equal(t, "some-userbase-table", gotTableUser)

		// Validate the query input
		gotCode := fakeDynamo.ScanCall.Receives.ScanInput.ExpressionAttributeValues
		wantCode := map[string]*dynamodb.AttributeValue{":0": {S: aws.String("kyle@gmail.com")}, ":1": {S: aws.String("kyle25")}}
		assert.Equal(t, wantCode, gotCode)

		gotUsername := *fakeDynamo.PutItemCall.Receives.PutItemInput.Item["username"].S
		//assert.Regexp(t, regexp.MustCompile(``), gotUsername)
		assert.Equal(t, "kyle25", gotUsername)

		gotFirstname := *fakeDynamo.PutItemCall.Receives.PutItemInput.Item["firstname"].S
		//assert.Regexp(t, regexp.MustCompile(``), gotFirstname)
		assert.Equal(t, "kyle", gotFirstname)

		gotLastname := *fakeDynamo.PutItemCall.Receives.PutItemInput.Item["lastname"].S
		//assert.Regexp(t, regexp.MustCompile(``), gotLastname)
		assert.Equal(t, "boitz", gotLastname)

		gotPassword := *fakeDynamo.PutItemCall.Receives.PutItemInput.Item["password"].S
		//assert.Regexp(t, regexp.MustCompile(``), gotPassword)
		assert.Equal(t, "password12!", gotPassword)

		gotEmail := *fakeDynamo.PutItemCall.Receives.PutItemInput.Item["email"].S
		//assert.Regexp(t, regexp.MustCompile(``), gotEmail)
		assert.Equal(t, "kyle@gmail.com", gotEmail)

		//validate the query output

	})

	t.Run("when the body request is invalid", func(t *testing.T) {
		c := routes.NewTestController(nil, nil, nil, lib.NewLogger(os.Stdout))

		// Setup call
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(
			http.MethodGet,
			"/users",
			io.NopCloser(strings.NewReader(`
				"username": "kyle25"
			}
			`)),
		)
		ctx.Request.Header.Set("Content-Type", "application/json")

		// Make call
		c.SearchUser(ctx)

		// Validate that the response is correct
		assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
	})
}

func TestGetUser(t *testing.T) {
	t.Run("when the call is successful", func(t *testing.T) {
		fakeDynamo := fakes.DynamoDBAPI{}

		c := routes.NewTestController(nil, &fakeDynamo, nil, lib.NewLogger(os.Stdout))

		// Setup call
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(
			http.MethodGet,
			"/users",
			io.NopCloser(strings.NewReader(`
			{
				"userid": "12321jkasdas"
			}
			`)),
		)
		ctx.Request.Header.Set("Content-Type", "application/json")

		// Make call
		c.GetUser(ctx)

		// Validate that the response is correct
		assert.Equal(t, http.StatusOK, ctx.Writer.Status())

		_, err := ioutil.ReadAll(w.Body)
		assert.NoError(t, err)

		// Validate the correct table
		gotTable := *fakeDynamo.GetItemCall.Receives.GetItemInput.TableName
		assert.Equal(t, "some-userbase-table", gotTable)

		// Validate the query input
		gotUserid := *fakeDynamo.GetItemCall.Receives.GetItemInput.Key["userid"]
		userid := dynamodb.AttributeValue{S: aws.String("12321jkasdas")}
		assert.Equal(t, gotUserid, userid)

		//validate the query output

	})

	t.Run("when the body request is invalid", func(t *testing.T) {
		c := routes.NewTestController(nil, nil, nil, lib.NewLogger(os.Stdout))

		// Setup call
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(
			http.MethodGet,
			"/users",
			io.NopCloser(strings.NewReader(`
				"userid": "12321jkasdas"
			}
			`)),
		)
		ctx.Request.Header.Set("Content-Type", "application/json")

		// Make call
		c.GetUser(ctx)

		// Validate that the response is correct
		assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("when the call is successful", func(t *testing.T) {
		fakeDynamo := fakes.DynamoDBAPI{}

		c := routes.NewTestController(nil, &fakeDynamo, nil, lib.NewLogger(os.Stdout))

		// Setup call
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(
			http.MethodPut,
			"/users",
			io.NopCloser(strings.NewReader(`
			{
				"userid": "12321jkasdas",
				"username": "kyle25",
				"firstname": "kyle",
				"lastname": "boitz",
				"password": "password12!",
				"email": "kyle@gmail.com"
			}
			`)),
		)
		ctx.Request.Header.Set("Content-Type", "application/json")

		// Make call
		c.UpdateUser(ctx)

		// Validate that the response is correct
		assert.Equal(t, http.StatusOK, ctx.Writer.Status())

		_, err := ioutil.ReadAll(w.Body)
		assert.NoError(t, err)

		// Validate the correct tables
		gotTableUser := *fakeDynamo.UpdateItemCall.Receives.UpdateItemInput.TableName
		assert.Equal(t, "some-userbase-table", gotTableUser)

		// Validate the query input
		gotUserid := fakeDynamo.UpdateItemCall.Receives.UpdateItemInput.Key
		userid := map[string]*dynamodb.AttributeValue{":0": {S: aws.String("12321jkasdas")}}
		assert.Equal(t, gotUserid, userid)

		gotUserInfo := fakeDynamo.UpdateItemCall.Receives.UpdateItemInput.ExpressionAttributeValues
		userInfo := map[string]*dynamodb.AttributeValue{
			":0": {S: aws.String("password12!")},
			":1": {S: aws.String("kyle25")},
			":2": {S: aws.String("kyle")},
			":3": {S: aws.String("boitz")},
			":4": {S: aws.String("kyle@gmail.com")},
		}
		assert.Equal(t, gotUserInfo, userInfo)

		//validate the query output

	})

	t.Run("when the body request is invalid", func(t *testing.T) {
		c := routes.NewTestController(nil, nil, nil, lib.NewLogger(os.Stdout))

		// Setup call
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(
			http.MethodPut,
			"/users",
			io.NopCloser(strings.NewReader(`
				"username": "kyle25"
			}
			`)),
		)
		ctx.Request.Header.Set("Content-Type", "application/json")

		// Make call
		c.UpdateUser(ctx)

		// Validate that the response is correct
		assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("when the call is successful", func(t *testing.T) {
		fakeDynamo := fakes.DynamoDBAPI{}

		c := routes.NewTestController(nil, &fakeDynamo, nil, lib.NewLogger(os.Stdout))

		// Setup call
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(
			http.MethodDelete,
			"/users",
			io.NopCloser(strings.NewReader(`
			{
				"userid": "12321jkasdas"
			}
			`)),
		)
		ctx.Request.Header.Set("Content-Type", "application/json")

		// Make call
		c.DeleteUser(ctx)

		// Validate that the response is correct
		assert.Equal(t, http.StatusOK, ctx.Writer.Status())

		_, err := ioutil.ReadAll(w.Body)
		assert.NoError(t, err)

		// Validate the correct table
		gotTable := *fakeDynamo.DeleteItemCall.Receives.DeleteItemInput.TableName
		assert.Equal(t, "some-userbase-table", gotTable)

		// Validate the query input
		gotUserid := *fakeDynamo.DeleteItemCall.Receives.DeleteItemInput.Key["userid"]
		userid := dynamodb.AttributeValue{S: aws.String("12321jkasdas")}
		assert.Equal(t, gotUserid, userid)

		//validate the query output

	})

	t.Run("when the body request is invalid", func(t *testing.T) {
		c := routes.NewTestController(nil, nil, nil, lib.NewLogger(os.Stdout))

		// Setup call
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(
			http.MethodDelete,
			"/users",
			io.NopCloser(strings.NewReader(`
				"userid": "12321jkasdas"
			}
			`)),
		)
		ctx.Request.Header.Set("Content-Type", "application/json")

		// Make call
		c.DeleteUser(ctx)

		// Validate that the response is correct
		assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
	})
}
