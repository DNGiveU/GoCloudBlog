package main

import (
	"controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	application := gin.Default()

	// gin.SetMode("release")
	gin.SetMode("debug")

	/*
	 * static resources
	 */
	application.StaticFile("/favicon.png", "web/dist/favicon.png")
	application.StaticFS("/dist", http.Dir("web/dist"))
	application.LoadHTMLGlob("web/dist/*.html")

	/*
	 * 路由注册
	 */
	// 根路由永久重定向至/index.html
	application.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index.html")
	})
	application.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	controller.RegistryRoute(application)

	/*
	 * listen and server on 0.0.0.0:80 (for windows "localhost:80")
	 * default: 8080
	 */
	application.Run(":80")
}