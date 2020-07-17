package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jlk/webapp-base/server/controllers"
	"github.com/jlk/webapp-base/server/data"
)

func main() {
	ginRouter := gin.Default()
	// Database connection happens in data/postgres.go init()

	ginRouter.GET("/devices", controllers.GetDevices)
	ginRouter.Run(fmt.Sprintf(":" + data.Config.GetString("ListenPort")))
}
