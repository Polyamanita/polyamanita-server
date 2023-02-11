package routes

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

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
//	@Param			query		query		string	true	"username query"
//	@success		201		{object}	routes.SearchUser.SearchOutputStruct	"String Array of Usernames"
//	@Failure		500
//	@Router			/users [get]
func (c *Controller) SearchUser(ctx *gin.Context) {
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
		ctx.Status(http.StatusNotFound)
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
		c.l.Debug(fmt.Sprintf("code expired when registering email: %v code: %v", body.Email, body.Code))
		ctx.Status(http.StatusUnauthorized)
		return
	}

	// check if username / email taken
	expr, err = e.NewBuilder().
		WithFilter(e.Name("Email").Equal(e.Value(body.Email)).
			Or(e.Name("Username").Equal(e.Value(body.Username)))).
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
	userID := uuid.NewString()
	user := &models.User{
		UserID:    userID,
		Username:  body.Username,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
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
//	@Param			UserID		path		string	true	"the user ID"
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
			e.Name("FirstName"),
			e.Name("LastName"),
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
//	@Param			UserID		path		string	true	"the user ID"
//	@Param			request	body	routes.UpdateUser.UpdateInputStruct	true	"User data"
//	@success		200
//	@Failure		500
//	@Router			/users/{UserID} [put]
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
	conditionEx := e.Equal(e.Name("password"), e.Value(body.Password))

	update := e.Set(
		e.Name("username"),
		e.Value(body.Username),
	).Set(
		e.Name("firstname"),
		e.Value(body.Firstname),
	).Set(
		e.Name("lastname"),
		e.Value(body.Lastname),
	).Set(
		e.Name("email"),
		e.Value(body.Email))

	expr, err := e.NewBuilder().WithUpdate(update).WithCondition(conditionEx).Build()

	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	} else {
		_, err = c.DynamoDB.UpdateItem(&dynamodb.UpdateItemInput{
			TableName: aws.String(c.secrets.ddbUserbaseTable),
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
	ctx.Status(http.StatusOK)
}

// DeleteUser godoc
//
//	@Summary		Deletes a User
//	@Description	Deletes a User with input data from DDB
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			UserID		path		string	true	"the user ID"
//	@Param			request	body	routes.DeleteUser.DeleteInputStruct	true	"Userid"
//	@success		200
//	@Failure		500
//	@Router			/users/{UserID} [delete]
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
	_, err := c.DynamoDB.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(c.secrets.ddbUserbaseTable),
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
	ctx.Status(http.StatusOK)
}

func (c *Controller) GetUserFollowers(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) GetUserFeed(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }
