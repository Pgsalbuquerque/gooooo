package routes

import (
	"strateegy/user-service/controller"
	"strateegy/user-service/server/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.GET("/", middlewares.Auth(), controller.GetUser)
			user.POST("/", controller.CreateUser)
			user.PUT("/")
			user.DELETE("/", middlewares.Auth(), controller.DeleteUser)
			user.PATCH("/password", middlewares.Auth(), controller.UpdateUserPassword)

		}

		session := main.Group("session")
		{
			session.POST("/login", controller.Login)
		}
	}

	return router
}
