package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetUserFollows(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) AddFollow(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) DeleteFollow(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) GetUserFollowers(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) GetUserFeed(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }
