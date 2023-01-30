package routes

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/polyamanita/polyamanita-server/src/lib"
)

func (c *Controller) Login(ctx *gin.Context) {
	// Input for registeration
	type LoginInputStruct struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	body := &LoginInputStruct{}
	if err := ctx.BindJSON(body); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	expr, err := expression.NewBuilder().
		WithFilter(expression.Name("email").Equal(expression.Value(body.Email)).
			And(expression.Name("password").Equal(expression.Value(body.Password)))).
		Build()
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	scanResp, err := c.DynamoDB.Scan(&dynamodb.ScanInput{
		TableName:                 aws.String(c.secrets.ddbUserbaseTable),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	})
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusUnauthorized)
		return
	}
	if *scanResp.Count == 0 {
		c.l.Debug(fmt.Sprintf(`invalid credentials for %v`, body.Email))
		ctx.Status(http.StatusBadRequest)
		return
	}

	id := *scanResp.Items[0]["id"].S
	token, err := lib.NewSignedToken(id, c.secrets.jwtKey, 24*time.Hour)
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.SetCookie("token", token, int(time.Now().Add(24*time.Hour).Unix()), "", "", true, false)

	type LoginOutputStruct struct {
		Id          string `json:"id"`
		AccessToken string `json:"accessToken"`
	}
	ctx.JSON(http.StatusOK, LoginOutputStruct{
		Id:          id,
		AccessToken: token,
	})
}

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
	codeExpiry := time.Now().Add(24 * time.Hour).Format(time.RFC3339)

	//TODO: check if email exists already, if so replace entry

	if _, err := c.DynamoDB.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(c.secrets.ddbVerificationTable),
		Item: map[string]*dynamodb.AttributeValue{
			"id":         {S: aws.String(uuid.New().String())},
			"email":      {S: aws.String(body.Email)},
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
