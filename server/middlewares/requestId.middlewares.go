package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Request-Id", uuid.NewString())
	}
}
