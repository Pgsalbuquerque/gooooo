package routes

import (
	"strateegy/user-service/controller"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.GET("/")
			user.POST("/", controller.CreateUser)
			user.PUT("/")
			user.DELETE("/", controller.DeleteUser)
			user.PATCH("/password", controller.UpdateUserPassword)

		}
	}

	return router
}
