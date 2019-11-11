package controller

import "github.com/gin-gonic/gin"

func RegistryRoute(engine *gin.Engine) {
	registryUserRoute(engine)
	registryLoginRoute(engine)
	registryMindRoute(engine)
}
