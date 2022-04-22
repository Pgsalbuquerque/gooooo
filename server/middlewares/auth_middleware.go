package middlewares

import (
	"strateegy/user-service/services/sessions"
	"strateegy/user-service/utils"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utils.RemoveBearer(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatus(403)
		}
		if !sessions.NewJWTService().ValidateToken(token) {
			c.AbortWithStatus(403)
		}
	}
}
