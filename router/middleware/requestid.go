package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Check for incoming header, ues it if exists
		requestId := context.Request.Header.Get("X-Request-Id")

		// Create request id with UUID4
		if requestId == "" {
			u4, _ := uuid.NewV4()
			requestId = u4.String()
		}

		// Expose it for use in the application
		context.Set("X-Request-Id", requestId)

		// Set X-Request-Id header
		context.Writer.Header().Set("X-Request-Id", requestId)
		context.Next()
	}
}
