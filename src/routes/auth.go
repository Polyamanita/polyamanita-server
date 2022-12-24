package routes

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gin-gonic/gin"
)

func (c *Controller) Login(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) Logout(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) AuthEmail(ctx *gin.Context) {
	const (
		Sender    = "no-reply@polyamanita.com"
		Recipient = "accrazed@gmail.com"
		Subject   = "Amazon SES Test (AWS SDK for Go)"
		HtmlBody  = "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
			"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
			"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"
		TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."
		CharSet  = "UTF-8"
	)

	result, err := c.SES.SendEmail(&ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(HtmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(Sender),
	})

	if err != nil {
		c.l.Print(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	c.l.Print("Email Sent to address: " + Recipient)
	c.l.Print(result)
}

func (c *Controller) RefreshAuthToken(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }
