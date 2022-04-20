package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rd67/gin-react-mySQL-socket/configs"
	"github.com/rd67/gin-react-mySQL-socket/middlewares"
	"github.com/rd67/gin-react-mySQL-socket/utils"
	v1 "github.com/rd67/gin-react-mySQL-socket/v1"
)

func init() {
	utils.ConnectDb()
}

func main() {
	router := gin.Default()

	router.Use(middlewares.JSONLogMiddleware())
	router.Use(middlewares.RequestIdHandler())
	router.Use(cors.Default())

	v1.SetupRoutes(router)

	//	Not Found Route
	router.NoRoute(func(c *gin.Context) {
		response := configs.ICommonResponse{
			StatusCode: http.StatusNotFound,
			Message:    "Route not found",
		}

		c.JSON(response.StatusCode, response)
		return
	})

	router.Run(fmt.Sprintf(":%d", configs.Config.App.Port))
}
