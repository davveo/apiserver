package middleware

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	// Parse the json web token
	return func(context *gin.Context) {
		if _, err := token.ParseRequest(context); err != nil {
			handler.SendResponse(context, errno.ErrTokenInvalid, nil)
			context.Abort()
			return
		}
		context.Next()
	}
}
