package v1Users

import "github.com/gin-gonic/gin"

func SetupRoutes(routes *gin.RouterGroup) {

	v1UserRoutes := routes.Group("/users")
	{
		v1UserRoutes.POST("/register", UserRegister)
		v1UserRoutes.POST("/login", UserLogin)
	}

}
