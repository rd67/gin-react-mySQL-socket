package v1

import (
	"github.com/gin-gonic/gin"

	v1Users "github.com/rd67/gin-react-mySQL-socket/v1/users"
)

func SetupRoutes(router *gin.Engine) {

	v1Routes := router.Group("/v1")

	v1Users.SetupRoutes(v1Routes)

}
