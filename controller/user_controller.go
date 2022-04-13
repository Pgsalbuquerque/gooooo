package controller

import (
	"strateegy/user-service/controller/dto"
	repository "strateegy/user-service/repositories/mongo"
	"strateegy/user-service/services/user"

	"github.com/gin-gonic/gin"
)

type Hello struct {
	Message string
}

func CreateUser(c *gin.Context) {
	var dto dto.UserRequestDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	repo := &repository.UserRepository{}
	service := user.NewUserService(repo)

	result, err := service.CreateUser(dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}

func UpdateUserPassword(c *gin.Context) {
	var dto dto.UserUpdatePasswordDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	repo := &repository.UserRepository{}
	service := user.NewUserService(repo)

	result, err := service.UpdateUserPassword(dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}
