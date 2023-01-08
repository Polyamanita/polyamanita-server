package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) TestHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello World!\nENVIRONMENT: %s", c.secrets.environment),
	})
}
