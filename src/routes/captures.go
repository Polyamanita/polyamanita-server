package routes

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
	"github.com/polyamanita/polyamanita-server/src/models"
)

func (c *Controller) GetCapturesList(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

// AddCaptures godoc
//	@Summary		Add a new list of captures to the user
//	@Description	Gets a list of mushrooms to add as captures to the user journal
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		routes.AddCaptures.AddCapturesInputStruct	true	"Email address to send code to"
//	@success		200		{object}	routes.AddCaptures.AddCapturesOutputStruct	"Expiry time of code"
//	@Failure		500
//	@Router			/users/:UserID/captures [post]
func (c *Controller) AddCaptures(ctx *gin.Context) {
	type AddCapturesInputStruct struct{}

	type AddCapturesOutputStruct struct{}
}

func (c *Controller) DeleteCaptures(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

// GetCapture godoc
//	@Summary		Get information about a captured mushroom
//	@Description	Gets all relevant information about a mushroom that's been captured for a user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@success		200		{object}	routes.GetCapture.GetCaptureOutputStruct	"mushroom information"
//	@Failure		500
//	@Router			/users/:UserID/captures/:CaptureID [get]
func (c *Controller) GetCapture(ctx *gin.Context) {
	UserID := ctx.Param("UserID")
	CaptureID := ctx.Param("CaptureID")

	res, err := c.DynamoDB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(c.secrets.ddbUserbaseTable),
		Key: map[string]*dynamodb.AttributeValue{
			"UserID":   {S: aws.String(UserID)},
			"MainSort": {S: aws.String(fmt.Sprintf("Capture#%s#%s", UserID, CaptureID))},
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
