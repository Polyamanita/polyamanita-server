package routes

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/gin-gonic/gin"
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
		//update TableName once user table is added to controller (ex: c.users.userTable)
		TableName:                 aws.String("Users"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})

	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	//return response
	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) RegisterUser(ctx *gin.Context) {
	//input for registeration
	type RegisterInputStruct struct {
		Code      string `json:"code"`
		Username  string `json:"username"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Password  string `json:"password"`
		Email     string `json:"email"`
	}

	body := &RegisterInputStruct{}
	if err := ctx.BindJSON(body); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	//check password requirements are met
	if len(body.Password) < 8 {
		c.l.Error("invalid password of length ", len(body.Password))
		ctx.Status(http.StatusBadRequest)
		return
	}

	//regex ?= broken
	/*ok, err := regexp.MatchString(`(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*\W)`, body.Password)
	if !ok || err != nil {
		c.l.Error("invalid password: ", err)
		ctx.Status(http.StatusBadRequest)
		return
	}*/

	//build expression to verify code is in the verification table
	keyEx := expression.Key("code").BeginsWith(body.Code)
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	_, err = c.DynamoDB.Query(&dynamodb.QueryInput{
		TableName:                 aws.String(c.secrets.ddbVerificationTable),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusUnauthorized)
		return
	}

	//conditionally put user into table if their username and email are not taken
	response, err := c.DynamoDB.PutItem(&dynamodb.PutItemInput{
		//update TableName once user table is added to controller (ex: c.users.userTable)
		TableName:           aws.String("Users"),
		ConditionExpression: aws.String("attribute_not_exists(username) and attribute_not_exists(email)"),
		Item: map[string]*dynamodb.AttributeValue{
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

	//return httpstatusOK
	ctx.JSON(http.StatusOK, response)
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
