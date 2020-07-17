package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jlk/webapp-base/server/controllers"
	"github.com/jlk/webapp-base/server/data"
)

func main() {
	ginRouter := gin.Default()
	// log.Println("Now server is running on port " + data.Config.GetString("ListenPort"))
	// err := http.ListenAndServe(":"+data.Config.GetString("ListenPort"), nil)
	// connect db
	// setup logs

	ginRouter.GET("/devices", controllers.GetDevices)
	ginRouter.Run(fmt.Sprintf(":" + data.Config.GetString("ListenPort")))
}
