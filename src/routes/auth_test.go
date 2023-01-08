package routes_test

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/polyamanita/polyamanita-server/src/fakes"
	"github.com/polyamanita/polyamanita-server/src/lib"
	"github.com/polyamanita/polyamanita-server/src/routes"
	"github.com/stretchr/testify/assert"
)

func TestPostAuths(t *testing.T) {
	t.Run("when the call is successful", func(t *testing.T) {
		fakeMail := fakes.MailIFace{}
		fakeDynamo := fakes.DynamoDBAPI{}

		c := routes.NewTestController(nil, &fakeDynamo, &fakeMail, lib.NewLogger(os.Stdout))

		// Setup call
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(
			http.MethodPost,
			"/auth",
			io.NopCloser(strings.NewReader(`
			{
				"email": "some-email@domain.com"
			}
			`)),
		)
		ctx.Request.Header.Set("Content-Type", "application/json")

		// Make call
		c.PostAuths(ctx)

		// Validate that the response is correct
		assert.Equal(t, http.StatusOK, ctx.Writer.Status())

		gotResp, err := ioutil.ReadAll(w.Body)
		assert.NoError(t, err)
		assert.Regexp(t, regexp.MustCompile(`^{"codeExpiry":.*}$`), string(gotResp))

		// Validate that the call made correct function calls
		gotTable := *fakeDynamo.PutItemCall.Receives.PutItemInput.TableName
		assert.Equal(t, "some-table", gotTable)
		gotCode := *fakeDynamo.PutItemCall.Receives.PutItemInput.Item["code"].S
		assert.Regexp(t, regexp.MustCompile(`\b\d{5}\b`), gotCode)
		gotExpiry := fakeDynamo.PutItemCall.Receives.PutItemInput.Item["codeExpiry"].S
		assert.NotNil(t, gotExpiry)

		assert.Equal(t, gotCode, fakeMail.SendEmailAuthCall.Receives.Code)
		assert.Equal(t, "some-email@domain.com", fakeMail.SendEmailAuthCall.Receives.Email)
	})

	t.Run("when the body request is invalid", func(t *testing.T) {
		c := routes.NewTestController(nil, nil, nil, lib.NewLogger(os.Stdout))

		// Setup call
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(
			http.MethodPost,
			"/auth",
			io.NopCloser(strings.NewReader(`
				"email": "some-email@domain.com"
			}
			`)),
		)
		ctx.Request.Header.Set("Content-Type", "application/json")

		// Make call
		c.PostAuths(ctx)

		// Validate that the response is correct
		assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
	})
}
