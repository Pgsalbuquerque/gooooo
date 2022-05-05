package controller

import (
	"strateegy/user-service/controller/dto"
	repository "strateegy/user-service/repositories/mongo"
	"strateegy/user-service/services/sessions"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var dto dto.UserLoginDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	db := 

	repo := &repository.UserRepository{}
	service := sessions.NewSessionService(repo)

	token, err := service.Login(dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, token)

}
