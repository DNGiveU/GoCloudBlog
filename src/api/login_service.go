package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * 登录
 */
func LoginAccount(c *gin.Context) {
	password := c.PostForm("password")
	userName := c.PostForm("userName")
	userType := c.PostForm("type")

	if password == "ant.design" && userName == "admin" {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"type": userType,
			"currentAuthority": []string {"admin"},
		})
	} else if password == "ant.design" && userName == "user" {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"type": userType,
			"currentAuthority": []string {"user"},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
			"type": userType,
			"currentAuthority": []string {"guest"},
		})
	}
}
