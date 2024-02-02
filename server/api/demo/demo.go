package demo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HellowDemo(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}
