package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Login(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) Logout(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) AuthEmail(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) RefreshAuthToken(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }
