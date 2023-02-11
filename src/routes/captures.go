package routes

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	e "github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/gin-gonic/gin"
	"github.com/polyamanita/polyamanita-server/src/models"
)

// GetCapturesList godoc
//
//	@Summary		Gets a list of captures from a User
//	@Description	Gets a list of captures from a User with input data from DDB
//	@Tags			Captures
//	@Accept			json
//	@Produce		json
//	@Param			UserID		path		string	true	"the user ID"
//	@success		201		{object}	routes.GetUser.GetOutputStruct				"string username"
//	@Failure		500
//	@Router			/users/{UserID}/captures [get]
func (c *Controller) GetCapturesList(ctx *gin.Context) {
	userID := ctx.Param("UserID")

	//build expression to query table
	partKey := e.Key("UserID").Equal(e.Value(userID))
	sortKey := e.Key("MainSort").BeginsWith("Capture#")
	projection := e.NamesList(e.Name("CaptureID"), e.Name("TimesFound"))
	expr, err := e.NewBuilder().
		WithKeyCondition(e.KeyAnd(partKey, sortKey)).
		WithProjection(projection).
		Build()
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
	}

	queryResp, err := c.DynamoDB.Query(&dynamodb.QueryInput{
		TableName:                 aws.String(c.secrets.ddbUserbaseTable),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	type GetCapturesListOutputStruct struct {
		Captures []*models.Capture `json:"captures"`
	}
	resp := &GetCapturesListOutputStruct{}
	if err := dynamodbattribute.UnmarshalListOfMaps(queryResp.Items, &resp.Captures); err != nil {
		c.l.Error("couldn't unmarshal query resp: ", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// AddCaptures godoc
//	@Summary		Add a new list of captures to the user
//	@Description	Gets a list of mushrooms to add as captures to the user journal
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			UserID		path		string	true	"the user ID"
//	@Param			request	body	routes.AddCaptures.AddCapturesInputStruct	true	"info to add and update the capture with. Will NOT overwrite notes if notes already exist. Instances will append"
//	@Failure		400
//	@Failure		401
//	@Router			/users/{UserID}/captures [post]
func (c *Controller) AddCaptures(ctx *gin.Context) {
	type AddCapturesInputStruct struct {
		Captures []*struct {
			Notes     string             `json:"notes"`
			CaptureID string             `json:"captureID"`
			Instances []*models.Instance `json:"instances"`
		} `json:"captures"`
	}

	body := &AddCapturesInputStruct{}
	if err := ctx.BindJSON(body); err != nil {
		c.l.Error("couldn't bind json: ", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	userID := ctx.Param("UserID")
	for _, capture := range body.Captures {
		expr, err := e.NewBuilder().
			WithUpdate(e.Set(
				e.Name("Instances"),
				e.ListAppend(
					e.IfNotExists(e.Name("Instances"), e.Value(&dynamodb.AttributeValue{L: []*dynamodb.AttributeValue{}})),
					e.Value(capture.Instances))).
				Add(e.Name("TimesFound"), e.Value(len(capture.Instances))).
				Set(e.Name("CaptureID"), e.IfNotExists(e.Name("CaptureID"), e.Value(capture.CaptureID))).
				Set(e.Name("Notes"), e.Value(capture.Notes))).
			Build()
		if err != nil {
			c.l.Error("unable to build expression: ", err)
		}

		if _, err := c.DynamoDB.UpdateItem(&dynamodb.UpdateItemInput{
			TableName: aws.String(c.secrets.ddbUserbaseTable),
			Key: map[string]*dynamodb.AttributeValue{
				"UserID":   {S: aws.String(userID)},
				"MainSort": {S: aws.String(fmt.Sprintf("Capture#%s#%s", userID, capture.CaptureID))},
			},
			UpdateExpression:          expr.Update(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
		}); err != nil {
			c.l.Error("unable to updateitem: ", err)
		}

		// Update stats
		expr, err = e.NewBuilder().
			WithUpdate(e.Add(e.Name("TotalCaptures"), e.Value(len(capture.Instances)))).
			Build()
		if err != nil {
			c.l.Error("unable to build expression: ", err)
		}
		if _, err := c.DynamoDB.UpdateItem(&dynamodb.UpdateItemInput{
			TableName: aws.String(c.secrets.ddbUserbaseTable),
			Key: map[string]*dynamodb.AttributeValue{
				"UserID":   {S: aws.String(userID)},
				"MainSort": {S: aws.String("Metadata")},
			},
			UpdateExpression:          expr.Update(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
		}); err != nil {
			c.l.Error("unable to updateitem: ", err)
		}
	}

	ctx.Status(http.StatusOK)
}

func (c *Controller) DeleteCaptures(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

// GetCapture godoc
//	@Summary		Get information about a captured mushroom
//	@Description	Gets all relevant information about a mushroom that's been captured for a user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			UserID		path		string	true	"the user ID"
//	@Param			CaptureID		path		string	true	"the capture ID"
//	@Success		200	{object}	routes.GetCapture.GetCaptureOutputStruct	"mushroom information"
//	@Failure		500
//	@Router			/users/{UserID}/captures/{CaptureID} [get]
func (c *Controller) GetCapture(ctx *gin.Context) {
	userID := ctx.Param("UserID")
	captureID := ctx.Param("CaptureID")

	res, err := c.DynamoDB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(c.secrets.ddbUserbaseTable),
		Key: map[string]*dynamodb.AttributeValue{
			"UserID":   {S: aws.String(userID)},
			"MainSort": {S: aws.String(fmt.Sprintf("Capture#%s#%s", userID, captureID))},
		},
	})
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	if len(res.Item) == 0 {
		ctx.Status(http.StatusNotFound)
	}

	type GetCaptureOutputStruct struct {
		Capture *models.Capture `json:"capture"`
	}
	response := GetCaptureOutputStruct{}
	if err := dynamodbattribute.UnmarshalMap(res.Item, &response.Capture); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) DownloadCaptureImage(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }
