package admin

import (
	"api"
	"github.com/gin-gonic/gin"
)

func RegistryMindRoute(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/getMindData", api.GetMindData)
}
