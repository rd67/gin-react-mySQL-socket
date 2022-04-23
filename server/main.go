package main

import (
	"fmt"
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rd67/gin-react-mySQL-socket/configs"
	"github.com/rd67/gin-react-mySQL-socket/middlewares"

	"github.com/rd67/gin-react-mySQL-socket/socket"
	// socketio "github.com/googollee/go-socket.io"
	"github.com/rd67/gin-react-mySQL-socket/utils"
	v1 "github.com/rd67/gin-react-mySQL-socket/v1"
)

func init() {
	utils.ConnectDb()
}

func main() {
	router := gin.Default()

	hub := socket.NewHub()
	go hub.Run()

	router.Use(middlewares.JSONLogMiddleware())
	router.Use(middlewares.RequestIdHandler())
	router.Use(middlewares.CorsMiddleware(configs.Config.App.AppURL))

	// router.GET("/socket.io/*any", gin.WrapH(server))
	// router.POST("/socket.io/*any", gin.WrapH(server))

	router.GET("/ws/", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
		socket.ServeWs(hub, ctx)
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
