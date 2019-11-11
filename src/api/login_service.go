package api

import (
	"constants"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"model/login"
	"net/http"
)

/**
 * 登录
 */
func LoginAccount(c *gin.Context) {
	bytes, _ := ioutil.ReadAll(c.Request.Body)
	login := login.Login{}
	err := json.Unmarshal(bytes, &login)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
			"type": 0,
			"currentAuthority": []string {"guest"},
		})
		return
	}
	password := login.Password
	userName := login.UserName
	userType := login.Type

	if password == "ant.design" && userName == "admin" {
		setCookie(c)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"type": userType,
			"currentAuthority": []string {"admin"},
		})
	} else if password == "ant.design" && userName == "user" {
		setCookie(c)
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

func setCookie(c *gin.Context) {
	/*
	 * https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Set-Cookie
	 * maxAge 在 cookie 失效之前需要经过的秒数。秒数为 0 或 -1 将会使 cookie 直接过期
	 * path 指定一个 URL 路径，这个路径必须出现在要请求的资源的路径中才可以发送 Cookie 首部
	 * domain 指定 cookie 可以送达的主机名
	 * secure 一个带有安全属性的 cookie 只有在请求使用SSL和HTTPS协议的时候才会被发送到服务器
	 * httpOnly 设置了 HttpOnly 属性的 cookie 不能使用 JavaScript 经由  Document.cookie 属性、XMLHttpRequest 和  Request APIs 进行访问，以防范跨站脚本攻击（XSS）
	 */
	c.SetCookie(constants.COOKIE_TOKEN, "123456", 60 * 60 * 24 * 30, "/", c.Request.Host, false, false)
}
