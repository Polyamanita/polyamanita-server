package routes

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *Controller) SearchUser(ctx *gin.Context) {
	//input for search query
	type SearchInputStruct struct {
		Username string `json:"username"`
	}
	body := &SearchInputStruct{}
	if err := ctx.BindJSON(body); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	//build expression to query table
	keyEx := expression.Key("username").BeginsWith(body.Username)
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	response, err := c.DynamoDB.Query(&dynamodb.QueryInput{
		TableName:                 aws.String(c.secrets.ddbUserbaseTable),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})

	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	if len(response.Items) == 0 {
		c.l.Debug("User not found: ", body.Username)
		ctx.Status(http.StatusNotFound)
		return
	}

	//return response
	ctx.JSON(http.StatusOK, response)
}

// RegisterUser godoc
//	@Summary		Registers a User
//	@Description	Registers the user with input data to DDB
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body	routes.RegisterUser.RegisterInputStruct	true	"User Data and code from email"
//	@success		201
//	@Failure		500
//	@Router			/users [post]
func (c *Controller) RegisterUser(ctx *gin.Context) {
	// Input for registeration
	type RegisterInputStruct struct {
		Code      string `json:"code"`
		Username  string `json:"username"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Password  string `json:"password"`
		Email     string `json:"email"`
	}

	body := &RegisterInputStruct{}
	if err := ctx.BindJSON(body); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	// Check password requirements are met
	if len(body.Password) < 8 {
		c.l.Error("invalid password of length ", len(body.Password))
		ctx.Status(http.StatusBadRequest)
		return
	}

	// Matcher checks if the password is INVALID, not valid
	ok, err := regexp.MatchString(`"^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[a-zA-Z0-9]*)$"`, body.Password)
	if ok || err != nil {
		c.l.Error("invalid password: ", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	// Check if email and code match
	expr, err := expression.NewBuilder().
		WithFilter(expression.Name("email").Equal(expression.Value(body.Email)).
			And(expression.Name("code").Equal(expression.Value(body.Code)))).
		Build()
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	scanResp, err := c.DynamoDB.Scan(&dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String(c.secrets.ddbVerificationTable),
	})
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusUnauthorized)
		return
	}
	if *scanResp.Count == 0 {
		c.l.Error(fmt.Sprintf(`couldn't find email "%v" with code "%v"`, body.Email, body.Code))
		ctx.Status(http.StatusUnauthorized)
		return
	}

	// check expiration
	expiry, err := time.Parse(time.RFC3339, *scanResp.Items[0]["codeExpiry"].S)
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	if time.Now().After(expiry) {
		c.l.Debug(fmt.Sprintf("code expired when registerring email: %v code: %v", body.Email, body.Code))
		ctx.Status(http.StatusUnauthorized)
		return
	}

	// check if username / email taken
	expr, err = expression.NewBuilder().
		WithFilter(expression.Name("email").Equal(expression.Value(body.Email)).
			Or(expression.Name("username").Equal(expression.Value(body.Username)))).
		Build()
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	scanResp, err = c.DynamoDB.Scan(&dynamodb.ScanInput{
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
	if *scanResp.Count != 0 {
		c.l.Error(fmt.Sprintf(`username "%v" or email "%v" already in use`, body.Email, body.Email))
		ctx.Status(http.StatusBadRequest)
		return
	}

	// put user into table
	id := uuid.NewString()
	_, err = c.DynamoDB.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(c.secrets.ddbUserbaseTable),
		Item: map[string]*dynamodb.AttributeValue{
			"id":        {S: aws.String(id)},
			"username":  {S: aws.String(body.Username)},
			"firstname": {S: aws.String(body.FirstName)},
			"lastname":  {S: aws.String(body.LastName)},
			"password":  {S: aws.String(body.Password)},
			"email":     {S: aws.String(body.Email)},
		},
	})
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c *Controller) GetUser(ctx *gin.Context) {
	//input for getting user
	type GetInputStruct struct {
		Userid string `json:"userid"`
	}

	body := &GetInputStruct{}
	if err := ctx.BindJSON(body); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	//search table using userid
	response, err := c.DynamoDB.GetItem(&dynamodb.GetItemInput{
		//update once user table is added to controller (ex: c.users.userTable)
		TableName: aws.String("Users"),
		Key: map[string]*dynamodb.AttributeValue{
			"userid": {S: aws.String(body.Userid)},
		},
	})
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	//return response
	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) UpdateUser(ctx *gin.Context) {
	//input for updating profile
	type UpdateInputStruct struct {
		Userid    string `json:"userid"`
		Username  string `json:"username"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Password  string `json:"password"`
		Email     string `json:"email"`
	}

	body := &UpdateInputStruct{}
	if err := ctx.BindJSON(body); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	//build expression to update user info only if password matches
	conditionEx := expression.Equal(expression.Name("password"), expression.Value(body.Password))

	update := expression.Set(
		expression.Name("username"),
		expression.Value(body.Username),
	).Set(
		expression.Name("firstname"),
		expression.Value(body.Firstname),
	).Set(
		expression.Name("lastname"),
		expression.Value(body.Lastname),
	).Set(
		expression.Name("email"),
		expression.Value(body.Email))

	expr, err := expression.NewBuilder().WithUpdate(update).WithCondition(conditionEx).Build()

	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	} else {
		_, err = c.DynamoDB.UpdateItem(&dynamodb.UpdateItemInput{
			//update TableName once user table is added to controller (ex: c.users.userTable)
			TableName: aws.String("Users"),
			Key: map[string]*dynamodb.AttributeValue{
				":0": {S: aws.String(body.Userid)},
			},
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			UpdateExpression:          expr.Update(),
			ConditionExpression:       expr.Condition(),
		})
		if err != nil {
			c.l.Error(err)
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}
	//return httpstatusOK
	ctx.JSON(http.StatusOK, "")
}

func (c *Controller) DeleteUser(ctx *gin.Context) {
	//input for deleting user
	type DeleteInputStruct struct {
		Userid string `json:"userid"`
	}

	body := &DeleteInputStruct{}
	if err := ctx.BindJSON(body); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	//search table using userid
	response, err := c.DynamoDB.DeleteItem(&dynamodb.DeleteItemInput{
		//update once user table is added to controller (ex: c.users.userTable)
		TableName: aws.String("Users"),
		Key: map[string]*dynamodb.AttributeValue{
			"userid": {S: aws.String(body.Userid)},
		},
	})
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	//return response
	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) GetUserFollowers(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) GetUserFeed(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }
