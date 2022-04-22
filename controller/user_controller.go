package controller

import (
	"strateegy/user-service/controller/dto"
	repository "strateegy/user-service/repositories/mongo"
	"strateegy/user-service/services/user"

	"github.com/gin-gonic/gin"
)

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

	token := c.GetHeader("Authorization")

	repo := &repository.UserRepository{}
	service := user.NewUserService(repo)

	result, err := service.UpdateUserPassword(token, dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}

func DeleteUser(c *gin.Context) {
	var dto dto.UserDeleteDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	token := c.GetHeader("Authorization")

	repo := &repository.UserRepository{}
	service := user.NewUserService(repo)

	err = service.DeleteUser(token, dto)
	if err != nil {
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(204, gin.H{})

}

func GetUser(c *gin.Context) {
	token := c.GetHeader("Authorization")

	repo := &repository.UserRepository{}
	service := user.NewUserService(repo)

	user, err := service.GetUser(token)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, user)
}
