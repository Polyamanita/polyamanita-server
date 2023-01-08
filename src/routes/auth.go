package routes

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func (c *Controller) Login(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) Logout(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

// PostAuths godoc
//	@Summary		Send a Verification Code
//	@Description	Sends an email to the address passed in with a verification code to verify their email
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		routes.PostAuths.AuthEmailInputStruct	true	"Email address to send code to"
//	@success		200		{object}	routes.PostAuths.AuthEmailOutputStruct	"Expiry time of code"
//	@Failure		500
//	@Router			/auths [post]
func (c *Controller) PostAuths(ctx *gin.Context) {
	type AuthEmailInputStruct struct {
		Email string `json:"email"`
	}
	body := &AuthEmailInputStruct{}
	if err := ctx.BindJSON(body); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	rand.Seed(time.Now().Unix())
	code := fmt.Sprintf("%05d", rand.Intn(100000))
	codeExpiry := time.Now().Add(24 * time.Hour).String()

	if _, err := c.DynamoDB.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(c.secrets.verificationTable),
		Item: map[string]*dynamodb.AttributeValue{
			"code":       {S: aws.String(code)},
			"codeExpiry": {S: aws.String(codeExpiry)},
		},
	}); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	if err := c.Mail.SendEmailAuth(body.Email, code); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	type AuthEmailOutputStruct struct {
		CodeExpiry string `json:"codeExpiry"`
	}
	ctx.JSON(http.StatusOK, &AuthEmailOutputStruct{
		CodeExpiry: codeExpiry,
	})
}

func (c *Controller) RefreshAuthToken(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }
