package lib

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
