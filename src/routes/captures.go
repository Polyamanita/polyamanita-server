package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetCapturesList(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) AddCaptures(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) DeleteCaptures(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) GetCapture(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }

func (c *Controller) DownloadCaptureImage(ctx *gin.Context) { ctx.Status(http.StatusNotImplemented) }
