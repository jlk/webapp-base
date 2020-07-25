package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jlk/webapp-base/server/controllers"
	"github.com/jlk/webapp-base/server/data"
)

func main() {
	ginRouter := gin.Default()
	ginRouter.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:8080"},
		AllowMethods:  []string{"GET"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	// Database connection happens in data/postgres.go init()

	ginRouter.GET("/devices", controllers.GetDevices)
	ginRouter.Run(fmt.Sprintf(":" + data.Config.GetString("ListenPort")))
}
