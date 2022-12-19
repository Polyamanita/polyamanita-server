package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) SearchUser(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) RegisterUser(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) GetUser(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) UpdateUser(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) DeleteUser(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) GetUserFollowers(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) GetUserFeed(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }
