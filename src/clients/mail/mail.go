package mail

import (
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type IFace interface {
	SendEmailAuth(email, code string) error
}

type MailClient struct {
	sc *sendgrid.Client
	l  *log.Logger
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
		c.l.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return nil

}

var _ IFace = (*MailClient)(nil)
