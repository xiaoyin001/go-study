package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowHome(c *gin.Context) {
	mJsonObj := gin.H{"message": "Welcome to the Homepage!"}
	mJsonObj["state"] = "success"

	c.JSON(http.StatusOK, mJsonObj)
}
