package routes

import (
	"net/http"
	"unicode"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/gin-gonic/gin"
)

func (c *Controller) SearchUser(ctx *gin.Context) {
	//input for search query
	type SearchInputStruct struct {
		username string `json:"username"`
	}
	body := &SearchInputStruct{}
	if err := ctx.BindJSON(body); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	//build expression to query table
	keyEx := expression.Key("username").BeginsWith(body.username)
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
	ctx.JSON(http.StatusOK, response.Items)
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
	if len(body.Password) >= 8 {
		var number = false
		var upper = false
		var special = false
		for idx, val := range body.Password {
			if unicode.IsNumber(val) {
				number = true
			}
			if unicode.IsUpper(val) {
				upper = true
			}
			if unicode.IsPunct(val) || unicode.IsSymbol(c) {
				special = true
			}
		}

		if number == false || upper == false || special == false {
			//password either doesnt have a number, upper or special
			c.l.Error(err)
			ctx.Status(http.StatusBadRequest)
		}

	} else {
		//password less than 8 characters
		c.l.Error(err)
		ctx.Status(http.StatusBadRequest)
	}

	//build expression to verify code is in the verification table
	keyEx := expression.Key("code").BeginsWith(body.Code)
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else {
		response, err = c.DynamoDB.Query(&dynamodb.QueryInput{
			TableName:                 aws.String(c.secrets.verificationTable),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			KeyConditionExpression:    expr.KeyCondition(),
		})
		if err != nil {
			c.l.Error(err)
			ctx.Status(http.StatusUnauthorized)
			return
		}
	}

	//conditionally put user into table if their username and email are not taken
	_, err = c.DynamoDB.PutItem(&dynamodb.PutItemInput{
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
	ctx.JSON(http.StatusOK, "")
}

func (c *Controller) GetUser(ctx *gin.Context) {
	//input for getting user
	type GetInputStruct struct {
		userid string `json:"userid"`
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
			"userid": {S: aws.String(body.userid)},
		},
	})
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	//return response
	ctx.JSON(http.StatusOK, response.Item)
}

func (c *Controller) UpdateUser(ctx *gin.Context) {
	//input for updating profile
	type UpdateInputStruct struct {
		userid    string `json:"userid"`
		username  string `json:"username"`
		firstname string `json:"firstname"`
		lastname  string `json:"lastname"`
		password  string `json:"password"`
		email     string `json:"email"`
	}

	body := &UpdateInputStruct{}
	if err := ctx.BindJSON(body); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	//build expression to update user info only if password matches
	conditionEx := expression.Equal(expression.Name("password"), expression.Value(body.password))
	updateUsername := expression.Set(expression.Name("username"), expression.Value(body.username))
	updateFirstname := expression.Set(expression.Name("firstname"), expression.Value(body.firstname))
	updateLastname := expression.Set(expression.Name("lastname"), expression.Value(body.lastname))
	updateEmail := expression.Set(expression.Name("email"), expression.Value(body.email))

	expr, err := expression.NewBuilder().WithUpdate(updateUsername, updateFirstname, updateLastname, updateEmail).WithCondition(conditionEx).Build()

	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	} else {
		_, err = c.DynamoDB.UpdateItem(&dynamodb.PutItemInput{
			//update TableName once user table is added to controller (ex: c.users.userTable)
			TableName: aws.String("Users"),
			Key: map[string]*dynamodb.AttributeValue{
				"userid": {S: aws.String(body.userid)},
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
	type GetInputStruct struct {
		userid string `json:"userid"`
	}

	body := &GetInputStruct{}
	if err := ctx.BindJSON(body); err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	//search table using userid
	response, err := c.DynamoDB.DeleteItem(&dynamodb.GetItemInput{
		//update once user table is added to controller (ex: c.users.userTable)
		TableName: aws.String("Users"),
		Key: map[string]*dynamodb.AttributeValue{
			"userid": {S: aws.String(body.userid)},
		},
	})
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	//return response
	ctx.JSON(http.StatusOK, "")
}

func (c *Controller) GetUserFollowers(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) GetUserFeed(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }
