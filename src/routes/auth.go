package routes

import (
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	e "github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/polyamanita/polyamanita-server/src/lib"
)

// Login godoc
//
//	@Summary	Login to an account
//	@Tags		Auth
//	@Param		request	body		routes.Login.LoginInputStruct	true	"login credentials"
//	@success	200		{object}	routes.Login.LoginOutputStruct	"string username"
//	@Failure	500
//	@Failure	400
//	@Failure	401
//	@Router		/session [post]
func (c *Controller) Login(ctx *gin.Context) {
	// Input for registeration
	type LoginInputStruct struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type ErrorOutputStruct struct {
		Response string `json:"response"`
	}

	body := &LoginInputStruct{}
	if err := ctx.BindJSON(body); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	expr, err := expression.NewBuilder().
		WithFilter(expression.Name("Email").Equal(expression.Value(body.Email)).
			And(expression.Name("Password").Equal(expression.Value(body.Password)))).
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
		ctx.Status(http.StatusInternalServerError)
		return
	}
	if *scanResp.Count == 0 {
		c.l.Debug(fmt.Sprintf(`invalid credentials for %v`, body.Email))
		ctx.JSON(http.StatusUnauthorized, &ErrorOutputStruct{Response: "invalid credentials"})
		return
	}

	id := *scanResp.Items[0]["UserID"].S
	token, err := lib.NewSignedToken(id, c.secrets.jwtKey, 24*time.Hour)
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.SetCookie("token", token, int(time.Now().Add(24*time.Hour).Unix()), "", "", true, false)

	type LoginOutputStruct struct {
		UserID      string `json:"userID"`
		AccessToken string `json:"accessToken"`
	}
	ctx.JSON(http.StatusOK, LoginOutputStruct{
		UserID:      id,
		AccessToken: token,
	})
}

// Logout godoc
//
//	@Summary		Logs user out of account
//	@Description	Just deletes the cookie the user was using to login
//	@Tags			Auth
//	@success		200
//	@Failure		500
//	@Router			/session [delete]
func (c *Controller) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "", "", false, false)

	ctx.Status(http.StatusOK)
}

// PostAuths godoc
//
//	@Summary		Send a register Verification Code
//	@Description	Sends an email to the address passed in with a verification code to verify their email
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		routes.PostAuths.AuthEmailInputStruct	true	"Email address to send code to"
//	@success		200		{object}	routes.PostAuths.AuthEmailOutputStruct	"Expiry time of code"
//	@Failure		500
//	@Router			/auth [post]
func (c *Controller) PostAuths(ctx *gin.Context) {
	type AuthEmailInputStruct struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	type ErrorOutputStruct struct {
		Response string `json:"response"`
	}

	body := &AuthEmailInputStruct{}
	if err := ctx.BindJSON(body); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	// Check password requirements are met
	if len(body.Password) < 8 {
		c.l.Error("invalid password length ", len(body.Password))
		ctx.JSON(http.StatusBadRequest, &ErrorOutputStruct{Response: fmt.Sprintf("invalid password length %v", body.Password)})
		return
	}

	// Matcher checks if the password is INVALID, not valid
	ok, err := regexp.MatchString(`"^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[a-zA-Z0-9]*)$"`, body.Password)
	if ok || err != nil {
		c.l.Error("invalid password: ", err)
		ctx.JSON(http.StatusBadRequest, &ErrorOutputStruct{Response: "password does not meet requirements"})
		return
	}

	// check if username / email taken
	expr, err := e.NewBuilder().
		WithFilter(e.Name("Email").Equal(e.Value(body.Email)).
			Or(e.Name("Username").Equal(e.Value(body.Username)))).
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
		ctx.Status(http.StatusInternalServerError)
		return
	}
	if *scanResp.Count != 0 {
		c.l.Error(fmt.Sprintf(`username "%v" or email "%v" already in use`, body.Username, body.Email))
		ctx.JSON(http.StatusBadRequest, &ErrorOutputStruct{Response: "username or email already in use"})
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

// PostAuthsGen godoc
//
//	@Summary		Send a general Verification Code
//	@Description	Sends an email to the address passed in with a verification code
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		routes.PostAuthsGen.AuthGenInputStruct	true	"Email address to send code to"
//	@success		201		{object}	routes.PostAuthsGen.AuthGenOutputStruct	"Expiry time of code"
//	@Failure		500
//	@Router			/authGen [post]
func (c *Controller) PostAuthsGen(ctx *gin.Context) {
	type AuthGenInputStruct struct {
		Email string `json:"email"`
	}
	body := &AuthGenInputStruct{}
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

	// Send email
	if err := c.Mail.SendEmailAuth(body.Email, code); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	type AuthGenOutputStruct struct {
		CodeExpiry string `json:"codeExpiry"`
	}
	ctx.JSON(http.StatusOK, &AuthGenOutputStruct{
		CodeExpiry: codeExpiry,
	})
}

func (c *Controller) RefreshAuthToken(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }
