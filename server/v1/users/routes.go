package v1Users

import (
	"github.com/gin-gonic/gin"
	"github.com/rd67/gin-react-mySQL-socket/middlewares"
)

func SetupRoutes(routes *gin.RouterGroup) {

	v1UserRoutes := routes.Group("/users")

	v1UserRoutes.POST("/register", UserRegister)
	v1UserRoutes.POST("/login", UserLogin)

	v1UserAuthRoutes := v1UserRoutes.Use(middlewares.AuthMiddleware())

	v1UserAuthRoutes.GET("/", UsersListing)

}
