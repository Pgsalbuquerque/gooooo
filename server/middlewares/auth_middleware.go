package middlewares

import (
	"context"
	"strateegy/user-service/grpc"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")

		conn := grpc.GetConn()
		client := grpc.NewSendValidateClient(conn)

		req := &grpc.Token{
			Token: token,
		}

		res, err := client.RequestValidate(context.Background(), req)
		if err != nil {
			println(err.Error())
			c.AbortWithStatus(403)
		}

		isValid := res.GetValidate()

		if !isValid {
			c.AbortWithStatus(403)
		}

	}
}
