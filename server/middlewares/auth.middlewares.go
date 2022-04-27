package middlewares

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rd67/gin-react-mySQL-socket/configs"
	"github.com/rd67/gin-react-mySQL-socket/models"
	"github.com/rd67/gin-react-mySQL-socket/pkg"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		unAuthorizedResponse := configs.ICommonResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "Unauthorized",
		}

		token, err := pkg.JwtValidateToken(c)
		if err != nil || !token.Valid {
			c.JSON(unAuthorizedResponse.StatusCode, unAuthorizedResponse)
			c.Abort()
			return
		}

		unAuthorizedResponse.Message = "Invalid token provided, kindly login again"

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["token_type"] != configs.TOKEN_TYPE_AUTH {
			c.JSON(unAuthorizedResponse.StatusCode, unAuthorizedResponse)
			c.Abort()
			return
		}

		user_id, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["id"]), 10, 32)
		if err != nil {
			c.JSON(unAuthorizedResponse.StatusCode, unAuthorizedResponse)
			c.Abort()
			return
		}

		var user models.User
		// var userToken models.UserToken

		//db.Preload("Post", "is_private = ? AND user_id != ?", "true", currentUser.ID).Find(&topicPosts)
		//Joins("user_tokens", utils.DBConn.Where(&userToken{"token": claims["id"]}))
		err = models.DBConn.
			Preload("UserTokens", "token = ?", token.Raw).
			Where(fmt.Sprintf("users.id = %d", user_id)).First(&user).Error
		if err != nil {
			c.JSON(unAuthorizedResponse.StatusCode, unAuthorizedResponse)
			c.Abort()
			return
		}

		c.Set("authUser", user)
		c.Set("authUserId", user.ID)

		c.Next()

	}
}
