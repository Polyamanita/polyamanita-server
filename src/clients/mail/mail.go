package mail

import (
	"fmt"
	"os"

	"github.com/polyamanita/polyamanita-server/src/lib"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type IFace interface {
	SendEmailAuth(email, code string) error
}

type MailClient struct {
	sc *sendgrid.Client
	l  *lib.Logger
}

func New(l *lib.Logger) (*MailClient, error) {
	key, ok := os.LookupEnv("SENDGRID_KEY")
	if !ok {
		return nil, fmt.Errorf("unable to find env var SENDGRID_KEY")
	}

	sc := sendgrid.NewSendClient(key)

	return &MailClient{
		sc: sc,
		l:  l,
	}, nil
}

func (c *MailClient) SendEmailAuth(email, code string) error {
	from := mail.NewEmail("Amelia", "no-reply@polyamanita.com")
	to := mail.NewEmail(email, email)
	msg := fmt.Sprintf(
		`<h1>Hey!</h1>
		<p>Please verify your email by using the code %s</p>
		<br/>
		<p>Amelia</p>`,
		code)

	response, err := c.sc.Send(
		mail.NewSingleEmail(from, "Authorize your Email", to, "", msg),
	)
	if err != nil {
		c.l.Error(err)
	} else {
		c.l.Debug("Successfully sent email")
		c.l.Debug(response.StatusCode)
		c.l.Debug(response.Body)
		c.l.Debug(response.Headers)
	}
	return nil
}

var _ IFace = (*MailClient)(nil)
