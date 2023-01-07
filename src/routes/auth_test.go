package routes_test

import (
	"os"
	"testing"

	"github.com/polyamanita/polyamanita-server/src/fakes"
	"github.com/polyamanita/polyamanita-server/src/lib"
	"github.com/polyamanita/polyamanita-server/src/routes"
)

func TestPostAuths(t *testing.T) {
	fakeMail := fakes.MailIFace{}
	fakeDynamo := fakes.DynamoDBAPI{}

	c := routes.NewTestController(nil, &fakeDynamo, &fakeMail, lib.NewLogger(os.Stdout))
}
