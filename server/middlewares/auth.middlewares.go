package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rd67/gin-react-mySQL-socket/configs"
	"github.com/rd67/gin-react-mySQL-socket/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := utils.JwtValidateToken(c); err != nil {
			response := configs.ICommonResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized",
			}

			c.JSON(response.StatusCode, response)
			c.Abort()
			return
		}

		c.Next()

	}
}
