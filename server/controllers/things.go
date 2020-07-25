package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jlk/webapp-base/server/data"
)

// Gets all devices
func GetThings(c *gin.Context) {
	// TODO: SECURITY: Should only return devices that the caller can see.
	things := []data.Thing{}
	err := data.PG.Select(&things, "SELECT * FROM thing")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": things})
	}
}
