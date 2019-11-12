package controller

import (
	"constants"
	"controller/admin"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

func RegistryRoute(engine *gin.Engine) {

	/*
	 * 可匿名访问
	 */
	admin.RegistryLoginRoute(engine)

	/*
	 * 需要认证的服务路由
	 */
	adminApiGroup := engine.Group(constants.ADMIN_API)
	adminApiGroup.Use(Logger())
	adminApiGroup.Use(BasicAuth())
	{
		admin.RegistryUserRoute(adminApiGroup)
		admin.RegistryMindRoute(adminApiGroup)
	}
}

/*
 * 认证
 * TODO implement
 */
func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// invoke
		c.Next()
	}
}

/*
 * 日志
 */
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before
		t := time.Now()

		// invoke
		c.Next()

		// after
		latency := time.Since(t)
		status := c.Writer.Status()
		log.Println("times: " + latency.String() + ", status: " + strconv.Itoa(status))
	}
}
