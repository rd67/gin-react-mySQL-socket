package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rd67/gin-react-mySQL-socket/configs"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtConfig = configs.Config.Jwt

func JwtGenerateToken(id uint, token_type string) (string, error) {
	claims := jwt.MapClaims{}

	claims["id"] = id
	claims["token_type"] = token_type
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(JwtConfig.HourLifespan)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(JwtConfig.SecretKey))

}

func JwtExtractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func JwtValidateToken(c *gin.Context) (*jwt.Token, error) {
	jwtToken := JwtExtractToken(c)

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JwtConfig.SecretKey), nil

	})

	if err != nil {
		return nil, err
	}

	return token, nil

	//	https://seefnasrul.medium.com/create-your-first-go-rest-api-with-jwt-authentication-in-gin-framework-dbe5bda72817

	//	https://github.com/akmamun/go-jwt
}
