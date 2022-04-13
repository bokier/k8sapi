package lib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckErr 基础定义错误
func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// CheckErrMsg 基础定义错误，自定义输入错误信息
func CheckErrMsg(err error, msg string) {
	if err != nil {
		fmt.Printf("[error] find a err in %v", msg)
	}
}

// Success gin中返回的正确信息
func Success(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
