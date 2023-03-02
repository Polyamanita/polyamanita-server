package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	e "github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/polyamanita/polyamanita-server/src/models"
)

// GetCapturesList godoc
//
//	@Summary		Gets a list of captures from a User
//	@Description	Gets a list of captures from a User with input data from DDB
//	@Tags			Captures
//	@Accept			json
//	@Produce		json
//	@Param			UserID	path		string												true	"the user ID"
//	@success		200		{object}	routes.GetCapturesList.GetCapturesListOutputStruct	"string username"
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

// GetAllCaptures godoc
//
//	@Summary		Gets captures from all Users
//	@Description	Gets all captures from DDB
//	@Tags			Captures
//	@Accept			json
//	@Produce		json
//	@success		201	{object}	routes.GetAllCaptures.GetAllCapturesOutputStruct	"List of all mushrooms"
//	@Failure		500
//	@Router			/users/captures [get]
func (c *Controller) GetAllCaptures(ctx *gin.Context) {

	//build expression to query table
	filter := e.Name("CaptureID").AttributeExists()
	expr, err := e.NewBuilder().
		WithFilter(filter).
		Build()
	if err != nil {
		c.l.Error(err)
		ctx.Status(http.StatusInternalServerError)
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

	type GetAllCapturesOutputStruct struct {
		Captures []*models.Capture `json:"capture"`
	}

	resp := &GetAllCapturesOutputStruct{}
	if err := dynamodbattribute.UnmarshalListOfMaps(scanResp.Items, &resp.Captures); err != nil {
		c.l.Error("couldn't unmarshal query resp: ", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	for _, capture := range resp.Captures {
		for _, instance := range capture.Instances {
			req, _ := c.S3.GetObjectRequest(&s3.GetObjectInput{
				Bucket: aws.String(c.secrets.s3StoreBucket),
				Key:    aws.String(instance.S3Key),
			})

			url, err := req.Presign(1 * time.Hour)
			if err != nil {
				c.l.Error("couldn't presign GetObjectRequest for key "+instance.S3Key+": ", err)
			}

			instance.ImageLink = url
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

// UploadCaptureImage godoc
//
//	@Summary		Uploads an image for the user to S3, FOR CAPTURE UPLOAD
//	@Description	USED FOR CAPTURE UPLOAD. MAKE REQUEST HERE TO GET S3 KEY, THEN INCLUDE S3KEY IN ADDCAPTURES REQUEST
//	@Tags			Captures
//	@Accept			json
//	@Produce		json
//	@Param			numLinks	query	number	false	"number of image upload links / S3 keys to generate. Default is 1"
//	@Param			UserID		path	string	true	"the user ID"
//	@Failure		400
//	@Failure		401
//	@Router			/users/{UserID}/images [post]
func (c *Controller) UploadCaptureImage(ctx *gin.Context) {
	numLinksQuery, ok := ctx.GetQuery("numLinks")
	if !ok {
		numLinksQuery = "1"
	}
	numLinks, err := strconv.Atoi(numLinksQuery)
	if err != nil {
		c.l.Error("invalid query for numLinks: ", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	userID := ctx.Param("UserID")

	type Link struct {
		UploadLink string `json:"uploadLink"`
		S3Key      string `json:"s3Key"`
	}
	type UploadCaptureImageOutputStruct struct {
		Links []*Link `json:"links"`
	}
	links := &UploadCaptureImageOutputStruct{
		Links: make([]*Link, numLinks),
	}
	for i := 0; i < numLinks; i++ {
		key := fmt.Sprintf("%s/%s", userID, uuid.NewString())

		req, _ := c.S3.PutObjectRequest(&s3.PutObjectInput{
			Bucket: aws.String(c.secrets.s3StoreBucket),
			Key:    aws.String(key),
		})

		url, err := req.Presign(24 * time.Hour)
		if err != nil {
			c.l.Error("couldn't presign PutObjectRequest for key "+key+": ", err)
		}

		links.Links[i] = &Link{
			S3Key:      key,
			UploadLink: url,
		}
	}

	ctx.JSON(http.StatusOK, links)
}

// AddCaptures godoc
//
//	@Summary		Add a new list of captures to the user
//	@Description	Gets a list of mushrooms to add as captures to the user journal
//	@Tags			Captures
//	@Accept			json
//	@Produce		json
//	@Param			UserID	path	string										true	"the user ID"
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

// GetCapture godoc
//
//	@Summary		Get information about a captured mushroom
//	@Description	Gets all relevant information about a mushroom that's been captured for a user, including image links
//	@Tags			Captures
//	@Accept			json
//	@Produce		json
//	@Param			UserID		path		string										true	"the user ID"
//	@Param			CaptureID	path		string										true	"the capture ID"
//	@Success		200			{object}	routes.GetCapture.GetCaptureOutputStruct	"mushroom information"
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
	capture := GetCaptureOutputStruct{}
	if err := dynamodbattribute.UnmarshalMap(res.Item, &capture.Capture); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	for _, instance := range capture.Capture.Instances {
		req, _ := c.S3.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String(c.secrets.s3StoreBucket),
			Key:    aws.String(instance.S3Key),
		})

		url, err := req.Presign(1 * time.Hour)
		if err != nil {
			c.l.Error("couldn't presign GetObjectRequest for key "+instance.S3Key+": ", err)
		}

		instance.ImageLink = url
	}

	ctx.JSON(http.StatusOK, capture)
}
