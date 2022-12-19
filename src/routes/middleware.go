package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/polyamanita/polyamanita-server/src/lib"
)

func (c *Controller) AuthenticateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := strings.Replace(ctx.GetHeader("Authorization"), "bearer ", "", 1)
		if len(auth) == 0 {
			auth, _ = ctx.Cookie("token")
		}

		_, valid := lib.ValidateToken(auth, c.jwtKey)
		if !valid {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		ctx.Next()
	}
}
func (c *Controller) AuthorizeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := strings.Replace(ctx.GetHeader("Authorization"), "bearer ", "", 1)
		if len(auth) == 0 {
			auth, _ = ctx.Cookie("token")
		}

		token, valid := lib.ValidateToken(auth, c.jwtKey)
		if !valid {
			ctx.Status(http.StatusUnauthorized)
			return
		}
		accountID := ctx.Param("accountID")
		if token.ID != accountID {
			ctx.Status(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}
