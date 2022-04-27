package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rd67/gin-react-mySQL-socket/configs"
	"github.com/rd67/gin-react-mySQL-socket/middlewares"
	"github.com/rd67/gin-react-mySQL-socket/models"

	"github.com/rd67/gin-react-mySQL-socket/pkg/websocket"
	v1 "github.com/rd67/gin-react-mySQL-socket/v1"
)

func init() {
	models.ConnectDb()
}

func main() {
	router := gin.Default()

	pool := websocket.NewPool()
	go pool.Start()

	// router.Use(gin.Recovery())
	router.Use(middlewares.JSONLogMiddleware()) //Custom Logs
	router.Use(middlewares.RequestIdHandler())
	router.Use(middlewares.CorsMiddleware(configs.Config.App.AppURL))

	router.GET("/ws/", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
		websocket.ServeWs(pool, ctx)
	})

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
