package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AvraamMavridis/randomcolor"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	e "github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/polyamanita/polyamanita-server/src/models"
)

// SearchUser godoc
//
//	@Summary		Searchs for a User
//	@Description	Searchs for Users with input data from DDB
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			query	query		string										true	"username query"
//	@success		200		{object}	routes.SearchUser.SearchUsersOutputStruct	"String Array of Usernames"
//	@Failure		500
//	@Router			/users [get]
func (c *Controller) SearchUser(ctx *gin.Context) {
	type ErrorOutputStruct struct {
		Response string `json:"response"`
	}

	//input for search query
	usernameQuery := ctx.Query("query")

	//build expression to query table
	expr, err := e.NewBuilder().
		WithFilter(e.And(
			e.Name("MainSort").BeginsWith("Metadata"),
			e.Name("Username").Contains(usernameQuery))).
		WithProjection(e.NamesList(
			e.Name("UserID"),
			e.Name("Username"),
			e.Name("TotalCaptures"),
			e.Name("Color1"),
			e.Name("Color2"))).
		Build()
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	scanResp, err := c.DynamoDB.Scan(&dynamodb.ScanInput{
		TableName:                 aws.String(c.secrets.ddbUserbaseTable),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	if scanResp == nil {
		c.l.Debug("User not found: ", usernameQuery)
		ctx.JSON(http.StatusNotFound, &ErrorOutputStruct{Response: "No users found"})
		return
	}

	type SearchUsersOutputStruct struct {
		Users []*models.User `json:"users"`
	}
	results := SearchUsersOutputStruct{}
	if err := dynamodbattribute.UnmarshalListOfMaps(scanResp.Items, &results.Users); err != nil {
		c.l.Error("unable to unmarshal user results: ", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, results)
}

// RegisterUser godoc
//
//	@Summary		Registers a User
//	@Description	Registers the user with input data to DDB
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body	routes.RegisterUser.RegisterInputStruct	true	"User Data and code from email"
//	@success		200
//	@Failure		500
//	@Router			/users [post]
func (c *Controller) RegisterUser(ctx *gin.Context) {
	// Input for registeration
	type RegisterInputStruct struct {
		Code     string `json:"code"`
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	type ErrorOutputStruct struct {
		Response string `json:"response"`
	}

	body := &RegisterInputStruct{}
	if err := ctx.BindJSON(body); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	// Check if email and code match
	expr, err := e.NewBuilder().
		WithFilter(e.Name("email").Equal(e.Value(body.Email)).
			And(e.Name("code").Equal(e.Value(body.Code)))).
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
		ctx.Status(http.StatusInternalServerError)
		return
	}
	if *scanResp.Count == 0 {
		c.l.Error(fmt.Sprintf(`couldn't find email "%v" with code "%v"`, body.Email, body.Code))
		ctx.JSON(http.StatusUnauthorized, &ErrorOutputStruct{Response: "couldn't match the email with the code entered"})
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
		c.l.Debug(fmt.Sprintf("code expired when registering email: %v code: %v", body.Email, body.Code))
		ctx.JSON(http.StatusUnauthorized, &ErrorOutputStruct{Response: "code has expired, try requesting a new one"})
		return
	}

	// put user into table
	userID := uuid.NewString()

	color1 := randomcolor.GetRandomColorInHex()
	time.Sleep(10 * time.Millisecond)
	color2 := randomcolor.GetRandomColorInHex()

	user := &models.User{
		UserID:   userID,
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
		MainSort: "Metadata",
		Color1:   color1,
		Color2:   color2,
	}
	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		c.l.Error("couldn't marshal user when registering: ", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	_, err = c.DynamoDB.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(c.secrets.ddbUserbaseTable),
		Item:      item,
	})
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}

// GetUser godoc
//
//	@Summary		Get a User
//	@Description	Gets one user with input data from DDB
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			UserID	path		string								true	"the user ID"
//	@success		201		{object}	routes.GetUser.GetUserOutputStruct	"string username"
//	@Failure		500
//	@Router			/users/{UserID} [get]
func (c *Controller) GetUser(ctx *gin.Context) {
	userID := ctx.Param("UserID")

	expr, err := e.NewBuilder().
		WithProjection(e.NamesList(
			e.Name("UserID"),
			e.Name("Username"),
			e.Name("Email"),
			e.Name("Password"),
			e.Name("Follows"),
			e.Name("TotalCaptures"),
			e.Name("Color1"),
			e.Name("Color2"),
		)).
		Build()
	if err != nil {
		c.l.Error("couldn't build expression: ", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	getItemResp, err := c.DynamoDB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(c.secrets.ddbUserbaseTable),
		Key: map[string]*dynamodb.AttributeValue{
			"UserID":   {S: aws.String(userID)},
			"MainSort": {S: aws.String("Metadata")},
		},
		ProjectionExpression:     expr.Projection(),
		ExpressionAttributeNames: expr.Names(),
	})
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	if len(getItemResp.Item) == 0 {
		ctx.Status(http.StatusNotFound)
		return
	}

	type GetUserOutputStruct struct {
		User *models.User `json:"user"`
	}
	result := GetUserOutputStruct{}
	if err := dynamodbattribute.UnmarshalMap(getItemResp.Item, &result.User); err != nil {
		c.l.Error("unable to unmarshal user results: ", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// UpdateUser godoc
//
//	@Summary		Updates a User
//	@Description	Updates a User with input data to DDB
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			UserID	path	string		true	"the user ID"
//	@Param			request	body	models.User	true	"User data"
//	@success		200
//	@Failure		500
//	@Router			/users/{UserID} [put]
func (c *Controller) UpdateUser(ctx *gin.Context) {
	type UpdateInputStruct struct {
		Username string `json:"username" dynamodbav:"Username"`
		Email    string `json:"email" dynamodbav:"Email"`
		Color1   string `json:"color1" dynamodbav:"Color1"`
		Color2   string `json:"color2" dynamodbav:"Color2"`
	}
	type ErrorOutputStruct struct {
		Response string `json:"response"`
	}

	userID := ctx.Param("UserID")

	body := &UpdateInputStruct{}
	if err := ctx.BindJSON(body); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusBadRequest)
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
		c.l.Error(fmt.Sprintf(`username "%v" or email "%v" already in use`, body.Email, body.Email))
		ctx.JSON(http.StatusBadRequest, &ErrorOutputStruct{Response: "username or email already in use"})
		return
	}

	// Update user
	update := e.UpdateBuilder{}
	if body.Username != "" {
		update = update.Set(e.Name("Username"), e.Value(body.Username))
	}
	if body.Email != "" {
		update = update.Set(e.Name("Email"), e.Value(body.Email))
	}
	if body.Color1 != "" {
		update = update.Set(e.Name("Color1"), e.Value(body.Color1))
	}
	if body.Color2 != "" {
		update = update.Set(e.Name("Color2"), e.Value(body.Color2))
	}
	expr, err = e.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	_, err = c.DynamoDB.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String(c.secrets.ddbUserbaseTable),
		Key: map[string]*dynamodb.AttributeValue{
			"UserID":   {S: aws.String(userID)},
			"MainSort": {S: aws.String("Metadata")},
		},
		UpdateExpression:          expr.Update(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	//return httpstatusOK
	ctx.Status(http.StatusOK)
}

// DeleteUser godoc
//
//	@Summary		Deletes a User
//	@Description	Deletes a User with input data from DDB
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			UserID	path	string	true	"the user ID"
//	@success		200
//	@Failure		500
//	@Router			/users/{UserID} [delete]
func (c *Controller) DeleteUser(ctx *gin.Context) {
	type ErrorOutputStruct struct {
		Response string `json:"response"`
	}

	userID := ctx.Param("UserID")

	// Delete UserID#Metadata
	_, err := c.DynamoDB.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(c.secrets.ddbUserbaseTable),
		Key: map[string]*dynamodb.AttributeValue{
			"UserID":   {S: aws.String(userID)},
			"MainSort": {S: aws.String("Metadata")},
		},
	})
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	// Delete UserID#Captures
	expr, err := e.NewBuilder().
		WithKeyCondition(e.KeyAnd(
			e.Key("UserID").Equal(e.Value(userID)),
			e.Key("MainSort").BeginsWith("Capture#"),
		)).
		WithProjection(e.NamesList(
			e.Name("UserID"),
			e.Name("MainSort"),
		)).
		Build()
	if err != nil {
		c.l.Error("unable to build expression: ", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	queryResp, err := c.DynamoDB.Query(&dynamodb.QueryInput{
		TableName:                 aws.String(c.secrets.ddbUserbaseTable),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})
	if err != nil {
		c.l.Error("unable to query user captures: ", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	for _, item := range queryResp.Items {
		_, err := c.DynamoDB.DeleteItem(&dynamodb.DeleteItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"UserID":   item["UserID"],
				"MainSort": item["MainSort"],
			},
		})
		if err != nil {
			c.l.Error("unable to delete item: ", err)
		}
	}
	// Delete S3 pictures

	ctx.Status(http.StatusOK)
}
