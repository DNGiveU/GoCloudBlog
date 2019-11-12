package controller

import (
	"constants"
	"controller/admin"
	"github.com/gin-gonic/gin"
	"middleware"
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
	adminApiGroup.Use(middleware.Logger())
	adminApiGroup.Use(middleware.BasicAuth())
	{
		admin.RegistryUserRoute(adminApiGroup)
		admin.RegistryMindRoute(adminApiGroup)
	}
}
